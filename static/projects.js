(() => {
	let navigationObserver;
	let timelineObserver;
	let descriptionObserver;
	let pendingProject;
	let openerNumber = 0;
	let closingOpener;
	let closing = false;
	let scrollFrame;

	const dialog = () => document.getElementById('project-dialog');
	const projectID = (url = location.href) => {
		const parsed = new URL(url, location.href);
		return parsed.origin === location.origin
			? parsed.pathname.match(/^\/projects\/id\/([^/]+)\/?$/)?.[1]
			: undefined;
	};

	const status = (message = '', failedURL) => {
		const modal = dialog()?.open;
		for (const id of ['project-request-status', 'project-dialog-status']) {
			const element = document.getElementById(id);
			if (!element) continue;
			element.hidden = !message || (id === 'project-dialog-status') !== modal;
			element.textContent = message;
			if (failedURL) {
				const link = document.createElement('a');
				link.href = failedURL;
				link.textContent = 'Open project page';
				link.setAttribute('hx-boost', 'false');
				element.append(' ', link);
			}
		}
	};

	const cancelPendingProject = () => {
		const request = pendingProject;
		pendingProject = undefined;
		request?.xhr.abort();
		status();
	};

	const enhanceProjectLinks = () => {
		for (const link of document.querySelectorAll('a[href]')) {
			if (!projectID(link.href) || link.hasAttribute('data-project-link')) continue;
			link.setAttribute('data-project-link', '');
			link.removeAttribute('target');
			link.setAttribute('hx-target', '#project-dialog-content');
			link.setAttribute('hx-select', '#project-detail');
			link.setAttribute('hx-swap', 'innerHTML show:none');
			link.setAttribute('hx-push-url', 'true');
			link.setAttribute('hx-sync', '#project-dialog:replace');
			if (!link.id) link.id = `project-opener-${++openerNumber}`;
			window.htmx?.process(link);
		}
	};

	const setDescriptionExpanded = (preview, expanded) => {
		const toggle = preview.querySelector('.project-description-toggle');
		const name = preview.closest('.project-card').querySelector('h3').textContent;
		const label = expanded ? 'Show less' : 'Read more';
		preview.toggleAttribute('data-description-expanded', expanded);
		toggle.setAttribute('aria-expanded', String(expanded));
		toggle.setAttribute('aria-label', `${label} about ${name}`);
		toggle.textContent = label;
		preview.querySelector('.project-description').scrollTop = 0;
	};

	const setupDescriptions = () => {
		descriptionObserver?.disconnect();
		const previews = [...document.querySelectorAll('.project-description-preview')];
		if (!previews.length) return;
		const widths = new WeakMap();
		const update = (preview) => {
			const description = preview.querySelector('.project-description');
			const toggle = preview.querySelector('.project-description-toggle');
			const expanded = toggle.getAttribute('aria-expanded') === 'true';
			preview.setAttribute('data-description-ready', '');
			// Measure at the collapsed height, including when an expanded card
			// is restored from history or becomes wider after a viewport change.
			preview.removeAttribute('data-description-expanded');
			const truncated = description.scrollHeight > description.clientHeight + 1;
			preview.toggleAttribute('data-description-truncated', truncated);
			toggle.hidden = !truncated;
			setDescriptionExpanded(preview, expanded && truncated);
		};
		descriptionObserver = new ResizeObserver((entries) => {
			for (const entry of entries) {
				if (Math.abs((widths.get(entry.target) ?? 0) - entry.contentRect.width) < 1) continue;
				widths.set(entry.target, entry.contentRect.width);
				update(entry.target);
			}
		});
		for (const preview of previews) {
			update(preview);
			const toggle = preview.querySelector('.project-description-toggle');
			toggle.onclick = () => {
				const expanded = toggle.getAttribute('aria-expanded') !== 'true';
				setDescriptionExpanded(preview, expanded);
				if (!expanded && toggle.getBoundingClientRect().top < 0) {
					preview.closest('.project-card').scrollIntoView({ block: 'start', behavior: 'instant' });
				}
			};
			descriptionObserver.observe(preview);
		}
	};

	const revealYear = (link) => {
		const scroller = link?.closest('.project-years-scroll');
		if (!scroller) return;
		const label = link.querySelector('span').getBoundingClientRect();
		const viewport = scroller.getBoundingClientRect();
		if (label.left < viewport.left + 8 || label.right > viewport.right - 8) {
			// Only scroll the track, keeping the project being read in place.
			scroller.scrollBy({
				left: (label.left + label.right - viewport.left - viewport.right) / 2,
				behavior: 'instant',
			});
		}
	};

	const selectYear = (hash, reveal = false) => {
		for (const link of document.querySelectorAll('.project-years a')) {
			if (link.hash === hash) {
				const changed = !link.hasAttribute('aria-current');
				link.setAttribute('aria-current', 'location');
				if (changed || reveal) revealYear(link);
			} else link.removeAttribute('aria-current');
		}
	};

	const updateYear = (reveal = false) => {
		const headings = [...document.querySelectorAll('.project-year')];
		if (!headings.length) return;
		const selected = document.querySelector('.project-years a[aria-current]');
		// Returning to the timeline should retain the year we were reading.
		if (selected && document.querySelector('.project-years-navigation').getBoundingClientRect().bottom > 0) {
			if (reveal) revealYear(selected);
			return;
		}
		let active = headings[0];
		for (const heading of headings) {
			if (heading.getBoundingClientRect().top > innerHeight * 0.3) break;
			active = heading;
		}
		selectYear(`#${active.id}`, reveal);
	};

	const setupTimeline = () => {
		const track = document.querySelector('.project-years-navigation');
		const scroller = track?.querySelector('.project-years-scroll');
		if (!scroller) return;
		for (const link of scroller.querySelectorAll('.project-years a')) {
			link.onclick = (event) => {
				if (event.defaultPrevented || event.button !== 0 || event.ctrlKey || event.metaKey || event.shiftKey || event.altKey) return;
				selectYear(link.hash, true);
			};
		}
		const arrows = [...track.querySelectorAll('[data-year-scroll]')];
		const updateControls = () => {
			// Compare against the full track width so reserving room for arrows
			// does not keep them visible after the viewport grows wide enough.
			const overflow = scroller.scrollWidth > track.clientWidth + 1;
			track.toggleAttribute('data-overflow', overflow);
			const end = scroller.scrollWidth - scroller.clientWidth;
			const canBack = overflow && scroller.scrollLeft > 1;
			const canForward = overflow && scroller.scrollLeft < end - 1;
			track.toggleAttribute('data-can-scroll-back', canBack);
			track.toggleAttribute('data-can-scroll-forward', canForward);
			for (const arrow of arrows) {
				arrow.hidden = !overflow;
				arrow.disabled = Number(arrow.dataset.yearScroll) < 0 ? !canBack : !canForward;
			}
		};
		const scrollTrack = (direction) => {
			// Focusing a control can bring the timeline back onto the screen.
			// Settle that active-year change before starting a manual scroll.
			updateYear();
			scroller.scrollBy({
				left: direction * scroller.clientWidth * 0.75,
				behavior: matchMedia('(prefers-reduced-motion: reduce)').matches ? 'instant' : 'smooth',
			});
		};
		for (const arrow of arrows) {
			arrow.onclick = () => scrollTrack(Number(arrow.dataset.yearScroll));
		}
		scroller.onscroll = updateControls;
		scroller.onfocusin = (event) => revealYear(event.target.closest('.project-years a'));
		scroller.onkeydown = (event) => {
			if (event.target !== scroller) return;
			if (event.key === 'ArrowLeft' || event.key === 'ArrowRight') {
				event.preventDefault();
				scrollTrack(event.key === 'ArrowLeft' ? -1 : 1);
			} else if (event.key === 'Home' || event.key === 'End') {
				event.preventDefault();
				scroller.scrollTo({ left: event.key === 'Home' ? 0 : scroller.scrollWidth, behavior: 'instant' });
			}
		};
		timelineObserver = new ResizeObserver(() => {
			updateControls();
			revealYear(track.querySelector('[aria-current]'));
		});
		timelineObserver.observe(track);
		timelineObserver.observe(scroller);
		updateControls();
	};

	const setupNavigation = () => {
		navigationObserver?.disconnect();
		timelineObserver?.disconnect();
		const navigation = document.getElementById('project-navigation');
		const button = document.getElementById('projects-back-to-top');
		if (!navigation || !button) return;
		button.onclick = () => {
			navigation.focus({ preventScroll: true });
			window.scrollTo({ top: 0, left: 0, behavior: 'instant' });
		};
		navigationObserver = new IntersectionObserver(([entry]) => {
			button.hidden = entry.isIntersecting || entry.boundingClientRect.bottom > 0;
		});
		navigationObserver.observe(navigation);
		setupTimeline();
		updateYear(true);
	};

	const syncDialog = () => {
		const modal = dialog();
		if (!modal) return;
		const detail = modal.querySelector('#project-detail');
		const shouldOpen = detail && detail.dataset.projectId === projectID();
		if (shouldOpen) {
			const close = modal.querySelector('[data-project-close]');
			close.href = modal.dataset.returnUrl || `/projects/#${detail.dataset.projectId}`;
			window.htmx?.process(close);
			// History restores HTML, but the dialog must re-enter the browser's
			// top layer to restore focus trapping and the inert background.
			if (!modal.matches(':modal')) {
				modal.removeAttribute('open');
				modal.showModal();
				detail.querySelector('h2').focus({ preventScroll: true });
			}
		} else if (modal.open) {
			modal.close();
		}
	};

	const setup = () => {
		closing = false;
		const initial = document.getElementById('project-initial-detail');
		if (initial) {
			const modal = dialog();
			const detail = initial.querySelector('#project-detail');
			modal.querySelector('#project-dialog-content').replaceChildren(detail);
			modal.dataset.returnUrl = `/projects/#${detail.dataset.projectId}`;
			modal.dataset.returnViaHistory = 'false';
			modal.dataset.depth = '0';
			initial.remove();
		}
		enhanceProjectLinks();
		setupDescriptions();
		setupNavigation();
		syncDialog();
		if (closingOpener) {
			const opener = closingOpener;
			closingOpener = undefined;
			// Removing the old modal also schedules native focus restoration.
			// Restore the opener after that browser cleanup has completed.
			requestAnimationFrame(() => document.getElementById(opener)?.focus({ preventScroll: true }));
		}
	};

	const closeDialog = () => {
		const modal = dialog();
		if (!modal || closing) return;
		cancelPendingProject();
		if (modal.dataset.returnViaHistory === 'true' && Number(modal.dataset.depth) > 0) {
			closing = true;
			closingOpener = modal.dataset.opener;
			history.go(-Number(modal.dataset.depth));
		} else {
			// A shared URL has no in-site origin to go back to. This boosted
			// link replaces its history entry with the catalogue at the card.
			modal.querySelector('[data-project-close]').click();
		}
	};

	document.addEventListener('click', (event) => {
		const close = event.target.closest('[data-project-close]');
		if (!close || event.ctrlKey || event.metaKey || event.shiftKey || event.altKey) return;
		cancelPendingProject();
		if (dialog().dataset.returnViaHistory === 'true') {
			event.preventDefault();
			event.stopPropagation();
			closeDialog();
		}
	}, true);

	document.addEventListener('keydown', (event) => {
		if (event.key === ' ' && event.target.matches('[data-project-close]')) {
			event.preventDefault();
			closeDialog();
		}
	});

	// Links in clipped descriptions remain reachable when tabbing through cards.
	document.addEventListener('focusin', (event) => {
		const preview = event.target.closest('.project-description-preview');
		if (preview?.hasAttribute('data-description-truncated') && event.target.closest('.project-description')) {
			setDescriptionExpanded(preview, true);
		}
	});

	document.addEventListener('cancel', (event) => {
		if (event.target.id !== 'project-dialog') return;
		event.preventDefault();
		closeDialog();
	}, true);

	let backdropPress = false;
	document.addEventListener('pointerdown', (event) => {
		backdropPress = event.target === dialog();
	});
	document.addEventListener('click', (event) => {
		if (backdropPress && event.target === dialog()) closeDialog();
		backdropPress = false;
	});

	document.addEventListener('htmx:beforeRequest', (event) => {
		if (event.detail.target?.id !== 'project-dialog-content') return;
		const modal = dialog();
		const fromDialog = modal.open;
		pendingProject = {
			xhr: event.detail.xhr,
			url: event.detail.elt.href,
			context: fromDialog ? { ...modal.dataset } : {
				returnUrl: location.pathname + location.search + location.hash,
				returnViaHistory: 'true',
				opener: event.detail.elt.id,
				depth: '0',
			},
		};
		status('Loading project…');
	});

	document.addEventListener('htmx:beforeSwap', (event) => {
		if (event.detail.target?.id !== 'project-dialog-content' || !event.detail.shouldSwap) return;
		const response = new DOMParser().parseFromString(event.detail.serverResponse, 'text/html');
		if (response.querySelector('#project-detail')?.dataset.projectId !== projectID(pendingProject?.url)) {
			event.detail.shouldSwap = false;
			event.detail.isError = true;
		}
	});

	document.addEventListener('htmx:afterSwap', (event) => {
		if (event.detail.target?.id === 'project-dialog-content' && pendingProject) {
			const modal = dialog();
			Object.assign(modal.dataset, pendingProject.context);
			modal.dataset.depth = String(Number(modal.dataset.depth) + 1);
			pendingProject = undefined;
			status();
			enhanceProjectLinks();
			syncDialog();
			modal.querySelector('#project-detail-title')?.focus({ preventScroll: true });
		} else if (event.detail.target === document.body) {
			setup();
		}
	});

	document.addEventListener('htmx:afterRequest', (event) => {
		if (pendingProject?.xhr !== event.detail.xhr) return;
		status('Couldn’t load this project. Try again or', pendingProject.url);
		pendingProject = undefined;
	});

	document.addEventListener('htmx:historyRestore', () => {
		cancelPendingProject();
		setup();
	});
	window.addEventListener('popstate', () => {
		cancelPendingProject();
		const modal = dialog();
		if (modal?.open && !projectID()) closingOpener = modal.dataset.opener;
	}, true);
	document.addEventListener('DOMContentLoaded', setup);
	window.addEventListener('scroll', () => {
		if (scrollFrame) return;
		scrollFrame = requestAnimationFrame(() => {
			updateYear();
			scrollFrame = undefined;
		});
	}, { passive: true });
})();

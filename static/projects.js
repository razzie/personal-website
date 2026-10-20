(() => {
	let observer;

	const setupProjectNavigation = () => {
		if (observer) {
			observer.disconnect();
			observer = null;
		}

		const navigation = document.getElementById('project-navigation');
		const button = document.getElementById('projects-back-to-top');
		if (!navigation || !button)
			return;

		button.hidden = true;
		button.onclick = () => {
			navigation.focus({ preventScroll: true });
			window.scrollTo({ top: 0, left: 0, behavior: 'instant' });
		};

		observer = new IntersectionObserver(([entry]) => {
			button.hidden = entry.isIntersecting || entry.boundingClientRect.bottom > 0;
		});
		observer.observe(navigation);
	};

	document.addEventListener('DOMContentLoaded', setupProjectNavigation);
	document.addEventListener('htmx:afterSwap', setupProjectNavigation);
})();

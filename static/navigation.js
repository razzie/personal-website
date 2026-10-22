(() => {
	const closeMenu = () => {
		document.querySelector('.site-menu')?.removeAttribute('open');
	};

	document.addEventListener('click', (event) => {
		if (!event.target.closest('.site-menu') || event.target.closest('.site-menu a')) {
			closeMenu();
		}
	});

	document.addEventListener('keydown', (event) => {
		const menu = document.querySelector('.site-menu[open]');
		if (event.key !== 'Escape' || !menu) return;
		menu.removeAttribute('open');
		menu.querySelector('summary').focus();
		event.preventDefault();
	});

	// History and boosted navigation should always start with a closed menu.
	for (const event of ['DOMContentLoaded', 'htmx:afterSwap', 'htmx:beforeHistorySave', 'htmx:historyRestore']) {
		document.addEventListener(event, closeMenu);
	}
	matchMedia('(max-width: 768px)').addEventListener('change', closeMenu);
})();

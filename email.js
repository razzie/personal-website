['DOMContentLoaded', 'htmx:afterSwap'].forEach(e => {
	document.addEventListener(e, () => {
		const emailElement = document.getElementById('email');
		if (emailElement === null)
			return;

		const user = 'gabor';
		const domain = 'gorzsony.com';
		const email = user + '@' + domain;

		const mailtoLink = document.createElement('a');
		mailtoLink.href = 'mailto:' + email;
		mailtoLink.textContent = email;

		emailElement.appendChild(mailtoLink);
	})
})
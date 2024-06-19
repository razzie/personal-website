document.addEventListener('DOMContentLoaded', (event) => {
	const user = 'gabor';
	const domain = 'gorzsony.com';
	const email = user + '@' + domain;
	const emailElement = document.getElementById('email');

	const mailtoLink = document.createElement('a');
	mailtoLink.href = 'mailto:' + email;
	mailtoLink.textContent = email;

	emailElement.appendChild(mailtoLink);
});

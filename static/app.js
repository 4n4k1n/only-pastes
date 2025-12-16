document.addEventListener('DOMContentLoaded', function() {

	const button = this.getElementById('button');

	button.addEventListener('click', function() {

		const content = document.getElementById('content').value;
		const language = document.getElementById('language').value;
		const expires_in = document.getElementById('expires').value;

		if (!content) {
			document.getElementById('message').innerHTML = 'Content can\'t be empty'
			return
		}

		const request = {
			content: content,
			language: language,
			expires_in: expires_in
		};

		fetch('/api/paste', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(request)
		})
		.then(response => response.json())
		.then(data => {
			if (data.slug) {
				window.location.href = '/' + data.slug;
			} else {
				document.getElementById('message').innerHTML = 'error'
			}
		})
		.catch(error => {
			document.getElementById('message').innerHTML = 'Failed to create paste'
		});
	});
});

document.addEventListener('DOMContentLoaded', function() {
	const slug = window.location.pathname.substring(1)

	fetch('/api/paste/' + slug)
		.then(response => {
			if (!response.ok) {
				throw new Error('Paste not found');
				
			}
			return response.json();
		})
		.then(data => {
			document.getElementById('err-msg').style.display = 'none';
            document.getElementById('paste-container').style.display = 'block';
            
			document.getElementById('language').textContent = data.language;
            document.getElementById('views').textContent = data.views;
            document.getElementById('created').textContent = data.createdAt;
            
			document.getElementById('content').textContent = data.content;

			// apply syntax later
		})
		.catch(error => {
			document.getElementById('paste-container').style.display = 'none';
            document.getElementById('err-msg').style.display = 'block';
            document.getElementById('err-msg').textContent = 'Paste not found or expired';
		})

		document.getElementById('copy_button').addEventListener('click', function() {
			const content = document.getElementById('content').textContent;
			navigator.clipboard.writeText(content);
			alert('Copied to clipboard!');
		});
});
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

			const codeElement = document.getElementById('content');
			codeElement.textContent = data.content;

			// Apply syntax highlighting
			if (typeof hljs !== 'undefined') {
				codeElement.className = 'language-' + (data.language || 'plaintext');
				hljs.highlightElement(codeElement);
			}
		})
		.catch(error => {
			document.getElementById('paste-container').style.display = 'none';
            document.getElementById('err-msg').style.display = 'block';
            document.getElementById('err-msg').textContent = 'Paste not found or expired';
		})

		const copyButton = document.getElementById('copy_button');
		copyButton.addEventListener('click', function() {
			const content = document.getElementById('content').textContent;
			navigator.clipboard.writeText(content).then(() => {
				const originalText = copyButton.textContent;
				copyButton.textContent = 'Copied!';
				copyButton.style.background = 'linear-gradient(135deg, #059669 0%, #047857 100%)';

				setTimeout(() => {
					copyButton.textContent = originalText;
					copyButton.style.background = '';
				}, 2000);
			}).catch(err => {
				alert('Failed to copy to clipboard');
			});
		});
});
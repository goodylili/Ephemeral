window.addEventListener('DOMContentLoaded', function() {
    const form = document.querySelector('form[action="/create-chat"]');
    const copyButton = document.getElementById('copyButton');

    form.addEventListener('submit', function(event) {
        event.preventDefault();
        copyButton.style.display = 'block';  // Show the button
    });

    copyButton.addEventListener('click', function() {
        const textToCopy = this.innerText;
        const textArea = document.createElement('textarea');
        textArea.value = textToCopy;
        document.body.appendChild(textArea);
        textArea.select();
        document.execCommand('copy');
        document.body.removeChild(textArea);
        alert('Text copied to clipboard!');
    });
});
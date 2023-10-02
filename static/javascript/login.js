document.addEventListener("DOMContentLoaded", function () {
    document.getElementById('myForm').addEventListener('submit', function (event) {
        event.preventDefault();

        // Your logic for generating the text after submission.
        let generatedText = "new text goes here";

        // Display the generated text in the popup.
        document.getElementById('copyText').value = generatedText;
        document.getElementById('copyPopup').style.display = 'block';
    });
});

function copyToClipboard() {
    let copyText = document.getElementById('copyText');
    copyText.select();
    document.execCommand("copy");

    closePopup();  // Close the popup after copying.
}


function closePopup() {
    document.getElementById('copyPopup').style.display = 'none';
}

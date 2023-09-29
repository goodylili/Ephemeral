const input = document.getElementById("input");
const output = document.getElementById("output");
const sendButton = document.getElementById("sendButton"); // Assuming you have a send button with id="sendButton"
const socket = new WebSocket("ws://localhost:8080/todo");

socket.onopen = function () {
    output.innerHTML += "Status: Connected<br>";
};

socket.onmessage = function (e) {
    output.innerHTML += "Server: " + e.data + "<br>";
};

input.addEventListener("keypress", function (e) {
    if (e.key === "Enter") {
        send();
    }
});

if (sendButton) {
    sendButton.addEventListener("click", send);
}

function send() {
    if (input.value.trim() !== "") { // Check to prevent sending empty messages
        socket.send(input.value);
        output.innerHTML += "You: " + input.value + "<br>"; // Display the sent message in the chat
        input.value = "";
    }
}

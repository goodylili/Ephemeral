document.addEventListener('DOMContentLoaded', function () {
    // Declare a websocket variable.
    let ws;

    // Determine whether to use "ws" (unsecured) or "wss" (secured) based on the current page's protocol.
    if (window.location.protocol === "https:") {
        ws = new WebSocket("wss://" + window.location.host + "/ws");
    } else {
        ws = new WebSocket("ws://" + window.location.host + "/ws");
    }

    // Set a handler function for when the websocket connection is opened.
    ws.onopen = function () {
        // Update the status display to show the connection is successful.
        document.getElementById("status").innerText = "Connected!";
    };

    // Set a handler function for when a message is received from the websocket connection.
    ws.onmessage = function (evt) {
        // Create a new message div for the received message.
        let div = document.createElement("div");
        div.className = "message received";

        // Add the name of the sender to the message (assuming "Other" for simplicity).
        let nameDiv = document.createElement("div");
        nameDiv.className = "name";
        nameDiv.innerText = "Other";
        div.appendChild(nameDiv);

        // Add the received message text.
        let textDiv = document.createElement("div");
        textDiv.className = "text";
        textDiv.innerText = evt.data;
        div.appendChild(textDiv);

        // Append the new message to the message container.
        document.getElementById("messageContainer").appendChild(div);
    };

    // Add an event listener to the "Send" button to handle message sending.
    document.getElementById("sendButton").onclick = function () {
        // Get the input element and the message text.
        let input = document.getElementById("inputMessage");

        // Send the message through the websocket.
        ws.send(input.value);

        // Create a new message div for the sent message.
        let div = docume3nt.createElement("div");
        div.className = "message sent";

        // Add the sender's name (You) to the message.
        let nameDiv = document.createElement("div");
        nameDiv.className = "name";
        nameDiv.innerText = "You";
        div.appendChild(nameDiv);

        // Add the sent message text.
        let textDiv = document.createElement("div");
        textDiv.className = "text";
        textDiv.innerText = input.value;
        div.appendChild(textDiv);

        // Append the new message to the message container.
        document.getElementById("messageContainer").appendChild(div);

        // Clear the input field after sending.
        input.value = "";
    };
});


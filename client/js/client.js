console.log("Hello Helow World");

var connection = new WebSocket('ws://' + document.location.host + '/game/foo');
connection.onopen = function () {
	console.log("Open");
};

connection.onerror = function (error) {
	console.log('WebSocket Error ' + error);
};

connection.onmessage = function(e) {
	console.log("You got message");
	document.getElementById("messages").innerHtml = e.data;
};

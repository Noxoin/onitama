console.log(window.location.pathname);

var connection = new WebSocket('ws://' + document.location.host + '/socket' + window.location.pathname);
connection.onopen = function () {
	console.log("Open");
};

connection.onerror = function (error) {
	console.log('WebSocket Error ' + error);
};

connection.onmessage = function(e) {
	console.log("You got message");
};

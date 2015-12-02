var busy = 0;
var errcode = 0;

$(document).ready(function() {
	var username = $("#user").val();
	var password = $("#pwd").val();

	console.log("document READY");

	$.get("/cgi-bin/login.cgi", {name:username, password:password}, function(data){
		console.log("Here only need check busy or not. data = " + data);
		$.each(data, function(key, value){
			//console.log("key = " + key + " ,value = " + value);
			if (key == "BUSY") {
				busy = value;
			}
		});

		if (busy == 1) {
			$("#description").text("busy");
		} else {
			$("#description").text("please input password");
		};
	});

	$("#form_login").submit(function(e){
		console.log("Submit the form!");
		e.preventDefault();

		username = $("#user").val();
		password = $("#pwd").val();
		$.get("/cgi-bin/login.cgi", {name:username, password:password}, function(data){
			console.log("ajax get call back function: return " + data);
			$.each(data, function(key, value){
				console.log("key = " + key + " ,value = " + value);
				if (key == "BUSY") {
					busy = value;
				};

				if (key == "ERRCODE") {
					errcode = value;
				};
			});

			if (busy == 1) {
				$("#description").text("busy");
			} else if (errcode == 1) {
				$("#description").text("wrong password");
			} else {
				//$("#description").text("CORRECT, please go to next page");
				setCookie("username", "admin");
				window.location.assign("./cgi_main.html");
			}
		});
	});

	$("#login_submit").click(function(){

		console.log("button submit click");
		// the password can't be NULL
		//console.log("password is " + $("#pwd").val());
		if ($("#pwd").val() == "") {
			alert("The password can't be NULL!");
		};

	 });
});



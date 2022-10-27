var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}

setTimeout(() =>{
	const errDiv = document.getElementById("login_error");

	errDiv.style.display= 'none';
}, 2500);
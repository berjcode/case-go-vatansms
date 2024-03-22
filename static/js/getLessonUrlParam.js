var urlParams = new URLSearchParams(window.location.search);
var userId = urlParams.get('id');
document.getElementById("userID").value = userId;

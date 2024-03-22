var urlParams = new URLSearchParams(window.location.search);
var userId = urlParams.get('userID');
document.getElementById("userID").value = userId;
console.log(userId);


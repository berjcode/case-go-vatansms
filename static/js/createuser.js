document.getElementById("userForm").addEventListener("submit", function(event) {
    event.preventDefault(); 

    var formData = new FormData(this);

    var jsonData = {};
    formData.forEach(function(value, key) {
        jsonData[key] = value;
    });

    fetch("/users", {
method: "POST",
headers: {
    "Content-Type": "application/json"
},
body: JSON.stringify(jsonData)
})
.then(response => {
if (!response.ok) {
    throw new Error("Kullanıcı oluşturma hatası: Sunucudan geçersiz yanıt alındı.");
}
return response.json();
})
.then(data => {
document.getElementById("message").innerText = "Kullanıcı başarıyla oluşturuldu: " + jsonData.userName;
})
.catch(error => {
console.error("Kullanıcı oluşturma hatası:", error);
document.getElementById("message").innerText = error.message;
});
});
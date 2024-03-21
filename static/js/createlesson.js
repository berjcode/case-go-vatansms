document.getElementById("lessonForm").addEventListener("submit", function(event) {
    event.preventDefault(); 

    var formData = new FormData(this);
    console.log(formData);
    var jsonData = {};
    formData.forEach(function(value, key) {
        jsonData[key] = value;
    });
    console.log(JSON.stringify(jsonData));
    fetch("/lesson", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(jsonData)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Ders oluşturma hatası: Sunucudan geçersiz yanıt alındı.");
        }
        return response.json();
    })
    .then(data => {
        console.error(jsonData.lessonName);
    })
    .catch(error => {
        console.error("Ders oluşturma hatası:", error);
    });
});

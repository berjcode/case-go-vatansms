document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault();
    const formData = new FormData(this);
    fetch("/login", {
        method: "POST",
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        const token = data.token;
        localStorage.setItem('token', token);
        
        fetch("/plan", {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + token
            }
        })
        .then(response => {
            window.location.href = "/plan";
        })
        .catch(error => {
            console.error('Error:', error);
        });
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

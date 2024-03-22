  window.onload = function() {
    fetch("/userid?username=" + "aa")
      .then(response => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then(data => {
        // Kullanıcı kimliğini al
        var userID = data.userIDs[0];
        var lessonURL = "/lesson?userID=" + userID;

        document.getElementById("lessonLink").href = lessonURL;
      })
      .catch(error => {
        console.error("There was a problem with the fetch operation:", error);
      });
  };
 

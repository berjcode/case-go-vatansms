

var userID = 3;
fetch("/lessons?userID=" + userID)
  .then(response => {
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.json();
  })
  .then(data => {
    var lessonTableBody = document.getElementById("lessonBody");

    data.forEach(lesson => {
      var row = document.createElement("tr");

      var cellID = document.createElement("td");
      cellID.textContent = lesson.ID;
      row.appendChild(cellID);

      var cellLessonName = document.createElement("td");
      cellLessonName.textContent = lesson.LessonName;
      row.appendChild(cellLessonName);

      var cellLessonDescription = document.createElement("td");
      cellLessonDescription.textContent = lesson.LessonDescription;
      row.appendChild(cellLessonDescription);

      var cellCreatedOn = document.createElement("td");
      cellCreatedOn.textContent = lesson.CreatedOn;
      row.appendChild(cellCreatedOn);

  var cellPlan = document.createElement("td");
  var PlanLink = document.createElement("a");
  PlanLink.href = "/plan?id=" + lesson.ID; 
  PlanLink.textContent = "Plan Oluştur";
  cellPlan.appendChild(PlanLink);
  row.appendChild(cellPlan);

  var cellEdit = document.createElement("td");
  var editLink = document.createElement("a");
  editLink.href = "/edit?id=" + lesson.ID; 
  editLink.textContent = "Düzenle";
  cellEdit.appendChild(editLink);
  row.appendChild(cellEdit);


  var cellDelete = document.createElement("td");
  var deleteLink = document.createElement("a");
  deleteLink.href = "/delete?id=" + lesson.ID; 
  deleteLink.textContent = "Sil";
  cellDelete.appendChild(deleteLink);
  row.appendChild(cellDelete);
      lessonTableBody.appendChild(row);
    });
  })
  .catch(error => {
    console.error("There was a problem with the fetch operation:", error);
  });

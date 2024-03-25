# Vatan Sms  Öğrenci Bilgi Sistemi Case  


Merhabalar,  Bana sunduğunuz  fırsat için teşekkür ederim. İstenen özellikler sağlamaya çalıştım. 
Eksikliklerim olabilir. 

Not: frontend tarafını'da tasarlamaya çalıştım bu yüzden çok fazla zaman kaybettim. Genel yazılım bilgimi tam anlamıyla yansıtamadım.

## Mimari 
-Standart Go Mimarisi 

Go tarafında çalışmadığım için çok fazla go ile mimari geliştirmedim. Bu  yüzden standart bir yapı ile case'i tamamladım.

database.go => handler.go => main.go 



# Yapı ve Versiyon
rest api .go 1.21.3 

# Git 
Master ve develop olmak üzere 2  farklı branch vardır.
# İndirme 
```
  git clone  
```

## Database Kurulumu
dbconfig.json dosyasına gerekli bilgileri girmeniz gerekir. 
```
 {
    "Database": {
        "Username": "root",
        "Password": "123456",
        "Host": "localhost",
        "Port": "3306",
        "Name": "newdb"
    }
}

```

```
Ardından  mysql tarafında veritabanını oluşturmalısınız . 
 CREATE DATABASE newdb;

```

* Projenin ana dizininine gelin ve şu komutları verin 
* Bu komut projeyi ayağa kaldıracaktır.
```
 go run main.go 
```

```
Uygulama İlk ayağa kalkarken seed data eklenmiştir.
Not: İlk Önce Login Olunuz. 
Login EndPoint :  http://localhost:8080/login
{
  "userNameAndEmail": "aaa",
  "password": "aaa"
}

Bir sorunla karşılaşırsanız kullanıcı oluşturma JSON:
Create User Endpoint: http://localhost:8080/users
{
    "UserName": "aaaa",
    "NameSurname": "aaaa",
    "Email": "aaa@gmail.com",
    "PasswordHash": "aaa",
    "CreatedBy" : "Admin"
}

Create PlanStatus Endpoint :  http://localhost:8080/planstatus
{
    "Name": "tammalandı",
    "CreatedBy": "Admin"
}

Create Lesson Endpoint : http://localhost:8080/lessons
{
    "LessonName": "felsefe",
    "LessonDescription": "felsefe iyidir",
    "UserID": 6,
    "CreatedBy" : "Admin"
}

Create PlanEndPoint : http://localhost:8080/plan  , HTTP POST, JSON

{
    "LessonID" : 23,
    "StartTime": "2024-03-23T14:30:00+03:00",
    "EndTime": "2024-03-23T16:30:00+03:00",
    "PlanStatusID" : 7,
    "CreatedBy": "Admin"
}
Response : 201 Created, True



Request : http://localhost:8080/plan/6  , HTTP Get, JSON, 6 = UserID
[
    {
        "ID": 20,
        "LessonName": "Matematik",
        "StartDay": "2024-03-23",
        "StartDate": "Saturday",
        "StartTime": "14:30:00",
        "EndDay": "2024-03-23",
        "EndDate": "Saturday",
        "EndTime": "16:30:00",
        "PlanStatusName": "tamamlandı",
        "CreatedOn": "2024-03-25T03:20:22+03:00"
    }
]
Response : 200 Ok, []Plans

```
Veriye erişimi gorm ile yaptığım için ilk önce database'i kurmalısınız.
Varlıkları oluştururken dikkat edilmesi gereken sıra!
```
1 - İlk önce bir User oluşturmalısınız 
2 - İkinci olarak PlanStatus oluşturmalısınız
3 - ve ardından Lesson Oluşturmalısınız
4 - Artık Plan'ı oluşturabilirsiniz.

Not: Lesson ders adını temsil eder. Daha esnek tutulabilir. Örnek : Bügün Matematik Çalışacağım gibi.
Not:  Aynı dersi birden fazla kez kullanabilinsin diye veritabanı daha esnek   tasarlanmıştır.
```
## Endpointler
```
    // User
	e.GET("/users/:id", handlers.GetUserData, mymiddleware.AuthMiddleware)
	e.PUT("/updateusers", handlers.UpdateUser, mymiddleware.AuthMiddleware)
	e.POST("/users", handlers.CreateUser)

	// Lesson
	e.PUT("/lesson", staticHandler.LessonDetailPageHtml, mymiddleware.AuthMiddleware)
	e.PUT("/updatelesson", handlers.UpdateLesson, mymiddleware.AuthMiddleware)
	e.POST("/lesson", handlers.CreateUserLesson, mymiddleware.AuthMiddleware)
	e.GET("/lesson/:id", handlers.GetLessonById, mymiddleware.AuthMiddleware)
	e.GET("/lessons/user/:userid", handlers.GetAllLessonsByUser, mymiddleware.AuthMiddleware)

	//Plan Status
	e.POST("/planstatus", handlers.CreatePlanStatus, mymiddleware.AuthMiddleware)
	e.PUT("/planstatus", handlers.UpdatePlanStatus, mymiddleware.AuthMiddleware)
	e.GET("/planstatus", handlers.GetAllPlanStatus, mymiddleware.AuthMiddleware)
	e.GET("/planstatus/:id", handlers.GetPlanStatusById, mymiddleware.AuthMiddleware)

	//Plan
	e.POST("/plan", handlers.CreatePlan, mymiddleware.AuthMiddleware)
	e.PUT("/plan", handlers.UpdatePlan, mymiddleware.AuthMiddleware)
	e.GET("/plan/:id", handlers.GetPlanById, mymiddleware.AuthMiddleware)
	e.GET("/plan/:userid", handlers.GetPlanDetails, mymiddleware.AuthMiddleware)
```



## Packages

* Paketlerin tam bilgileri go.mod dosyasında bulunmaktadır.

* github.com/dgrijalva/jwt-go
* github.com/go-sql-driver/mysql
* github.com/golang-jwt/jwt
* github.com/jinzhu/gorm
* github.com/labstack/echo/v4
* golang.org/x/crypto
* gorm.io/gorm


                                                                                                                      
   ###    By Abdullah Balikci - berjcode


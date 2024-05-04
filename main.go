package main

import (
	"cats_social/postgres"
	"cats_social/routers"

	"github.com/gin-gonic/gin"
)


func main(){
 	router := gin.Default()
 	routers.SetupUserRoutes(router)
 	router.Run(":8080")

	defer postgres.DB.Close()
}

// type User struct {
// 	ID int 
// 	Name string
// 	Email string
// 	Password string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// func createUser(){
// 	var user = User{}

// 	sqlStatement := `
// 	INSERT INTO users (name, email, password, createdat, updatedat)
// 	VALUES ($1, $2, $3, $4, $5)
// 	RETURNING *
// 	`

// 	err := db.QueryRow(sqlStatement, "Ismail", "ismailbinmail@gmail.com", "123456", time.Now(),time.Now()).
// 	Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("New user ID: %d\n", user.ID)
// }

// type Cat struct {
// 	ID int 
// 	UserId int
// 	HasMatched bool
// 	Name string
// 	Race string
// 	Sex string
// 	AgeInMonth int
// 	Description string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// func createCat(){
// 	var cat = Cat{}

// 	sqlStatement := `
// 	INSERT INTO cats (user_id, hasmatched, name, race, sex, ageinmonth, description, createdat, updatedat)
// 	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
// 	RETURNING *
// 	`

// 	err := db.QueryRow(sqlStatement, 3, false, "Mely", "Persian", "female", 3, "Mely is a cute cat", time.Now(),time.Now()).
// 	Scan(&cat.ID, &cat.UserId, &cat.HasMatched, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.CreatedAt, &cat.UpdatedAt)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("New cat ID: %d\n", cat.ID)
// }

// type CatImage struct {
// 	ID int 
// 	CatId int
// 	ImageUrl string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// func createCatImage(){
// 	var catImage = CatImage{}

// 	sqlStatement := `
// 	INSERT INTO catimage (cat_id, image_url, createdat, updatedat)
// 	VALUES ($1, $2, $3, $4)
// 	RETURNING *
// 	`

// 	err := db.QueryRow(sqlStatement, 1, "https://o-cdn-cas.sirclocdn.com/parenting/images/Cara-Merawat-Kucing-Persia.width-800.format-webp.webp", time.Now(),time.Now()).
// 	Scan(&catImage.ID, &catImage.CatId, &catImage.ImageUrl, &catImage.CreatedAt, &catImage.UpdatedAt)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("New cat image ID: %d\n", catImage.ID)
// }

// type MatchCats struct {
// 	ID int
// 	MatchCatID int
// 	UserCatID int
// 	Message string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// func createMatchCats(){
// 	var matchCats = MatchCats{}

// 	sqlStatement := `
// 	INSERT INTO matchcats (matchcatid, usercatid, message, createdat, updatedat)
// 	VALUES ($1, $2, $3, $4, $5)
// 	RETURNING *
// 	`

// 	err := db.QueryRow(sqlStatement, 1, 2, "I love you Mely as a cat persia", time.Now(),time.Now()).
// 	Scan(&matchCats.ID, &matchCats.MatchCatID, &matchCats.UserCatID, &matchCats.Message, &matchCats.CreatedAt, &matchCats.UpdatedAt)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("New match cat ID: %d\n", matchCats.ID)
// }
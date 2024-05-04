package controllers

import (
	"cats_social/postgres"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func generateJWT(id int, name string, email string) (string, error ){
	var jwtSecret = []byte("aku_suka_kucing_miaw")
	var jwtExpiry = time.Now().Add(time.Hour * 8).Unix()

	claims := jwt.MapClaims{
		"id": id,
		"name": name,
		"email": email,
		"exp": jwtExpiry,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}


type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func RegisterUser(ctx *gin.Context){
	var newUser User

    if err := ctx.ShouldBindJSON(&newUser); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        ctx.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    newUser.Password = string(hashedPassword)

    sqlStatement := `
    INSERT INTO users (name, email, password, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, name, email, created_at, updated_at
    `

    err = postgres.DB.QueryRow(sqlStatement, newUser.Name, newUser.Email, newUser.Password, time.Now(), time.Now()).
    Scan(&newUser.ID, &newUser.Name, &newUser.Email, &newUser.CreatedAt, &newUser.UpdatedAt)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := generateJWT(newUser.ID, newUser.Name, newUser.Email) 
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data": map[string]interface{}{
			"name":  newUser.Name,
			"email": newUser.Email,
			"accessToken": token,
		},
	})
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func LoginUser(ctx *gin.Context){
    // login using email and password
    var loginReq LoginRequest
    var user User

    if err := ctx.ShouldBindJSON(&loginReq); err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    sqlStatement := `
    SELECT id, name, email, password, created_at, updated_at
    FROM users
    WHERE email = $1
    `

    err := postgres.DB.QueryRow(sqlStatement, loginReq.Email).
        Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

    if err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
    if err != nil {
        ctx.AbortWithError(http.StatusBadRequest, err)
        return
    }

	token, err := generateJWT(user.ID, user.Name, user.Email)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

    ctx.JSON(http.StatusOK, gin.H{
        "message": "User logged in successfully",
        "data": map[string]interface{}{
            "name":  user.Name,
            "email": user.Email,
			"accessToken": token,
        },
    })
    
}
package api

import (
	db "at01/db/sqlc"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// POST
// Sign up
type createAccountRequest struct {
	Email    string  `json:"email" binding:"required"`
	Password string  `json:"password" binding:"required"`
	TypeUser db.Type `json:"type_user" binding:"required,oneof=user admin"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error() + " body",
		})
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	arg := db.CreateAccountParams{
		Email:    req.Email,
		Password: string(hash),
		TypeUser: db.Type(req.TypeUser),
	}

	result, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": " faild " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// Signin
func (server *Server) signIn(ctx *gin.Context) {

	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error() + " body",
		})
		return
	}

	args := db.Login{
		Email: req.Email,
	}

	// Find user
	result, err := server.store.First(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": " faild " + err.Error(),
		})
		return
	}
	// hash the password
	error := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password))

	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "faild password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})
		return
	}

	ctx.Request.Header.Add("Authorization", "Bearer "+tokenString)

	ctx.Request.Header.Add("Content-Type", "application/json")
	ctx.Set("id", result.ID)
	ctx.Set("email", result.Email)
	ctx.Next()
	ctx.JSON(http.StatusOK, tokenString)
}

package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jwt-go"
)

//var mySigningKey = os.Getenv("MY_JWT_TOKEN")
var mySigningKey = []byte("mysupersecrectphrase")

//homePage function
func homePage(ctx gin.Context) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(ctx.Writer, err.Error())
	}
	fmt.Fprintf(ctx.Writer, validToken)

}

func reqHandler() {
	server := gin.Default()
	server.GET("/", homePage(*gin.Context))
	server.Run("localhost:8081")
}

//GenerateJWT function
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorised"] = true
	claims["user"] = "Ramesh Mantri"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func main() {
	fmt.Println("My simple client")
	reqHandler()
	// fmt.Println(mySigningKey)

	// tokenString, err := GenerateJWT()
	// if err != nil {
	// 	fmt.Errorf("Error generating token string")
	// }
	// fmt.Println(tokenString)
}

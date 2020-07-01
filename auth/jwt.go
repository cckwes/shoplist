package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"

	service "github.com/cckwes/shoplist/services"
)

func JwtMiddleware(context *fiber.Ctx) {
	auth := context.Get(fiber.HeaderAuthorization)
	authContent := strings.Split(auth, " ")

	if len(authContent) != 2 {
		log.Println("No JWT token found")
		context.Status(403)
		return
	}

	token := authContent[1]

	jwtToken, err := jwt.Parse(token, getKey())
	if err != nil {
		log.Println("JWT parsing error: ", err)
		context.Status(403)
		return
	}

	email := jwtToken.Claims.(jwt.MapClaims)["email"].(string)
	if email == "" {
		log.Println("No email address in JWT claims")
		context.Status(403)
		return
	}

	user, err := service.GetOrCreateUser(email)
	if err != nil {
		log.Println("Fail to create or get user, error", err)
		context.Status(403)
		return
	}

	context.Locals("user", user)
	context.Next()
}

func getKey() func(*jwt.Token) (interface{}, error) {
	if os.Getenv("APP_ENV") == "production" {
		return getPublicKey
	} else {
		return getKeyDev
	}
}

func getKeyDev(token *jwt.Token) (interface{}, error) {
	secretString := os.Getenv("JWT_DEV_SECRET")
	return []byte(secretString), nil
}

func getPublicKey(token *jwt.Token) (interface{}, error) {
	audience := os.Getenv("JWT_AUDIENCE")
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(audience, true)
	if !checkAud {
		return nil, errors.New("Invalid audience")
	}

	issuer := os.Getenv("JWT_ISSUER")
	checkIssuer := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, true)
	if !checkIssuer {
		return nil, errors.New("Invalid issuer")
	}

	const jwksUrl = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

	resp, err := http.Get(jwksUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(bodyData, &data)
	if err != nil {
		return nil, err
	}

	kid := token.Header["kid"].(string)

	cert, ok := data[kid].(string)
	if !ok {
		errorMessage := fmt.Sprintf("No cert found with kid %v", kid)
		return nil, errors.New(errorMessage)
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))

	return result, nil
}

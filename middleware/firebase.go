package middleware

import (
	"cc-course-service/helper"
	"context"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

var authClient *auth.Client

func init() {
	firebaseCred := "creds.json"

	ctx := context.Background()
	opt := option.WithCredentialsFile(firebaseCred)
	conf := &firebase.Config{
		ProjectID: "capstone-project-46d2b",
	}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		panic(err)
	}

	authClient, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		panic(err)
	}
}

func FirebaseAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		if header == "" {
			return helper.Response(c, 401, "Firebase 401 1", nil)
		}
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return helper.Response(c, 401, "Firebase 401 2", nil)
		}
		token := parts[1]

		claims, err := authClient.VerifyIDToken(context.Background(), token)
		if err != nil {
			return helper.Response(c, 401, "Firebase 401 3", nil)
		}

		c.Locals("claims", claims.Claims)

		// print id and name
		log.Println("[Id]" + claims.Claims["user_id"].(string))
		log.Println("[Name]" + claims.Claims["name"].(string))

		return c.Next()
	}
}

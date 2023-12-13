package database

import (
	"context"
	"fmt"
	"log"
	"mail-client/configs"
	"mail-client/internal"
	"mail-client/internal/cors"
	"mail-client/internal/models"
	"mail-client/internal/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/resendlabs/resend-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var validate = validator.New()

type DB struct {
	client *mongo.Client
}

func ConnectDB() (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
	internal.Handle(err)

	internal.Handle(err)

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")
	return &DB{client: client}, err
}

func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("MailClient").Collection(collectionName)
}

func (db *DB) ctxDeferHelper(collectionName string) (*mongo.Collection, context.Context, context.CancelFunc) {
	collection := colHelper(db, collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	return collection, ctx, cancel
}

func (db *DB) resErrHelper(collectionName string, input any) (*mongo.InsertOneResult, context.CancelFunc, error) {
	collection, ctx, cancel := db.ctxDeferHelper(collectionName)

	res, err := collection.InsertOne(ctx, input)

	internal.Handle(err)

	return res, cancel, err
}

func welcomeMail(email string) {
	apiKey := configs.EnvApiKey()

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "BottleHub <team@bottlehub.github.io>",
		To:      []string{email},
		Subject: "Ahoy, Welcome to the Revolution",
		Html:    "<p>Congrats mate! You've been added to the list of crew mates that are in line for the future of competitive gambling.</p><br/> <p>In the coming weeks you'll receive exclusive updates, which will include a private alpha, as well as a dicord community. Just go on out and claim tokens before we close the <a href='https://bottlehub.io/faucet'> faucet</a>.<p><br /> <p>We will send you a confirmation to claim your spot as soon as we're done with the initial test version. So look out and get ready to come aboard. <br/><br/> <strong>The BottleHub Team.</strong></p>",
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		log.Panic(err)
	}
}

func (db *DB) AddEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.EnableCors(&c.Writer)
		var email models.Email

		if err := c.BindJSON(&email); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&email); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": validationErr.Error()}})
			return
		}

		address := models.Email{
			Address: email.Address,
		}
		res, cancel, err := db.resErrHelper("addresses", address)
		//welcomeMail(email.Address)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"error": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{
			"id": res,
		}})
	}
}

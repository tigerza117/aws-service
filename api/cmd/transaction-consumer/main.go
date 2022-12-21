package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"api/model"
	"api/query"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	queue := flag.String("q", "", "The name of the queue")
	flag.Parse()

	if *queue == "" {
		*queue = os.Getenv("SQS_NAME")
	}

	gormDB, err := gorm.Open(mysql.Open(os.Getenv("DB_DSN")))
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&model.Customer{}, &model.Account{}, &model.Tx{})
	query.SetDefault(gormDB)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("SQS_REGION")),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sqs.NewFromConfig(cfg)

	// Get URL of queue
	gQInput := &sqs.GetQueueUrlInput{
		QueueName: queue,
	}

	result, err := GetQueueURL(context.TODO(), client, gQInput)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := result.QueueUrl

	go func() {
		for {
			time.Sleep(time.Second)
			gMInput := &sqs.ReceiveMessageInput{
				MessageAttributeNames: []string{
					string(types.QueueAttributeNameAll),
				},
				QueueUrl:            queueURL,
				MaxNumberOfMessages: 1,
				VisibilityTimeout:   int32(100),
			}

			msgResult, err := GetMessages(context.TODO(), client, gMInput)
			if err != nil {
				fmt.Println("Got an error receiving messages:")
				fmt.Println(err)
				return
			}

			if msgResult.Messages != nil {
				fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
				fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
				idStr := msgResult.Messages[0].MessageAttributes["TxID"].StringValue
				id, err := strconv.Atoi(*idStr)
				if err != nil {
					panic(err)
				}
				t, err := query.Tx.Where(query.Tx.ID.Eq(uint(id))).First()
				if err != nil {
					panic(err)
				}

				if err := query.Q.Transaction(func(tx *query.Query) error {
					if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(t.DstAccountID)).UpdateSimple(tx.Account.Balance.Add(t.Amount)); err != nil {
						return err
					}
					if t.AccountID != nil {
						if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(*t.AccountID), tx.Account.Balance.Gte(t.Amount)).UpdateSimple(tx.Account.Balance.Sub(t.Amount)); err != nil {
							return err
						}
					}
					t.Status = model.TransactionSuccess
					if _, err := tx.WithContext(context.Background()).Tx.Updates(t); err != nil {
						return err
					}
					return nil
				}); err != nil {
					panic(err)
				}
			} else {
				fmt.Println("No messages found")
			}
		}
	}()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Am healthy fuck you")
	})
	app.Listen(os.Getenv("LISTEN"))
}

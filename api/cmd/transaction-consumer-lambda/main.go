package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"os"
	"strconv"

	"api/model"
	"api/query"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
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
	//gormDB.AutoMigrate(&model.Customer{}, &model.Account{}, &model.Tx{})
	query.SetDefault(gormDB)

	//if msgResult.Messages != nil {
	for i := 0; i < len(sqsEvent.Records); i++ {
		fmt.Println("Message ID:     " + sqsEvent.Records[i].MessageId)
		fmt.Println("Message Handle: " + sqsEvent.Records[i].ReceiptHandle)
		idStr := sqsEvent.Records[i].MessageAttributes["TxID"].StringValue
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
		} else {
		}
		fmt.Printf("done %d\n", idStr)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}

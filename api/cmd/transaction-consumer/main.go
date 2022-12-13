package main

import (
	"api/model"
	"api/query"
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func main() {
	queue := flag.String("q", "", "The name of the queue")
	flag.Parse()

	if *queue == "" {
		*queue = "transaction"
	}

	gormDB, err := gorm.Open(mysql.Open("root:pass@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	gormDB.AutoMigrate(&model.Customer{}, &model.Account{}, &model.Tx{})
	query.SetDefault(gormDB)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
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
			t.Status = model.TransactionSuccess

			if err := query.Q.Transaction(func(tx *query.Query) error {
				if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(t.DesAccountID)).UpdateSimple(tx.Account.Balance.Add(t.Amount)); err != nil {
					return err
				}
				if _, err := tx.WithContext(context.Background()).Account.Where(tx.Account.ID.Eq(t.AccountID), tx.Account.Balance.Gte(t.Amount)).UpdateSimple(tx.Account.Balance.Sub(t.Amount)); err != nil {
					return err
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
}

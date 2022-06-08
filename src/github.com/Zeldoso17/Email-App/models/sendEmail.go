package models

type SendEmail struct {
	From    string `bson:"from"`
	Subject string `bson:"subject"`
	Body    string `bson:"body"`
}
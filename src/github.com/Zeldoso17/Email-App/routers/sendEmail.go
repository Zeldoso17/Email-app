package routers

import (
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/Zeldoso17/Email-App/models"
	gomail "gopkg.in/gomail.v2"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {

	var info models.SendEmail
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	newMessage := gomail.NewMessage()

	newMessage.SetHeader("From", info.From)
	newMessage.SetHeader("To", "zeldosomeza17@gmail.com")
	newMessage.SetHeader("Subject", info.Subject)
	newMessage.SetBody("text/plain", info.Body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "zeldosomeza17@gmail.com", "LukeSkywalker")
	fmt.Println("Sending email...")
	if err = dialer.DialAndSend(newMessage); err != nil {
		fmt.Println(err)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("El email ha sido enviado"))
}
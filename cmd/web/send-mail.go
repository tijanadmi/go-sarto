package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/tijanadmi/go-sarto/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {
	
	server := mail.NewSMTPClient()
	server.Host = app.MailServer.Host
	a, e := strconv.Atoi(app.MailServer.PortNumber) 
    if e == nil { 
        fmt.Printf("%T \n %v", a, a) 
    }
	server.Port = a
	server.Username = app.MailServer.Username
	server.Password = app.MailServer.Password 
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	// Enable STARTTLS for a secure connection
	server.Encryption = mail.EncryptionSTARTTLS 

	client, err := server.Connect()
	if err != nil {
		errorLog.Println((err))
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)

	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
		

	} else {
		data, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
		}

		mailTemplate := string(data)
		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)
		
	}

	/*email.SetBody(mail.TextHTML, string(m.Content))*/

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent!")
	}
}

package mail

import (
	"fmt"
	"net/smtp"
)

type EmailRequest struct {
	From        string            `json:"from"`
	To          []string          `json:"to"`
	Subject     string            `json:"subject"`
	Bcc         []string          `json:"bcc"`
	Cc          []string          `json:"cc"`
	ReplyTo     string            `json:"reply_to"`
	Html        string            `json:"html"`
	Text        string            `json:"text"`
	Tags        []Tag             `json:"tags"`
	Attachments []Attachment      `json:"attachments"`
	Headers     map[string]string `json:"headers"`
}

type SendEmailResponse struct {
	Id string `json:"id"`
}

type Email struct {
	Id        string   `json:"id"`
	Object    string   `json:"object"`
	To        []string `json:"to"`
	From      string   `json:"from"`
	CreatedAt string   `json:"created_at"`
	Subject   string   `json:"subject"`
	Html      string   `json:"html"`
	Text      string   `json:"text"`
	Bcc       []string `json:"bcc"`
	Cc        []string `json:"cc"`
	ReplyTo   []string `json:"reply_to"`
	LastEvent string   `json:"last_event"`
}

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Attachment struct {
	Content  []string `json:"content,omitempty"`
	Filename string   `json:"filename"`
	Path     string   `json:"path,omitempty"`
}

func Send() {
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"

	// Set up the authentication
	auth := smtp.PlainAuth("", "ewanretorokugbe@gmail.com", "", smtpServer)

	from := "BottleHub <team@bottlehub.io>"
	to := "ewanretorokugbe@gmail.com"
	subject := "Ahoy, Welcome to the Revolution\n"
	body := "<p>Congrats mate! You've been added to the list of crew mates that are in line for the future of competitive gambling.</p><br/> <p>In the coming weeks you'll receive exclusive updates, which will include a private alpha, as well as a dicord community. Just go on out and claim tokens before we close the <a href='https://bottlehub.io/faucet'> faucet</a>.<p><br /> <p>We will send you a confirmation to claim your spot as soon as we're done with the initial test version. So look out and get ready to come aboard. <br/><br/> <strong>The BottleHub Team.</strong></p>"

	message := []byte(subject + body)

	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}

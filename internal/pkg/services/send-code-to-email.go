package services

import (
	"RolePlayModule/internal/utils/config"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"time"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(subject string, content string, to []string, cc []string, bcc []string) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}

func SendCodeToEmailService(cfg config.Config, code string, email string) error {
	sender := NewGmailSender("Your verification code", cfg.AppEmail, cfg.AppEmailPassword)
	subject := "52FOOD"
	content := fmt.Sprintf("<body>  <style>   body {    margin: 0;    padding: 0;    box-sizing: border-box;    font-family: sans-serif;    background: #f6f7f8;    pointer-events: none;    user-select: none;   }   .container {    max-width: 600px;    margin: 0 auto;    padding: 16px;   }   .wrapper {    background: #fff;    border-radius: 24px;    padding: 24px;    display: flex;    flex-direction: column;    row-gap: 8px;   }   .heading {    font-size: 24px;    font-weight: 500;   }   .footnote {    font-size: 16px;    color: rgba(0, 0, 0, 0.5);   }   .code {    margin-top: 16px;    display: flex;    align-items: center;    justify-content: center;    background: #f6f7f8;    padding: 12px;    font-size: 32px;    font-weight: 500;    letter-spacing: 16px;    text-align: center;    border-radius: 16px;   }   .border {    margin: 12px 0;    width: 100%;    height: 1px;    background: rgba(0, 0, 0, 0.08);   }   .logo {    font-size: 24px;    font-weight: 500;   margin-bottom: 16px;  }</style><div class=\"container\"><center><div class=\"logo\">52FOOD</div></center><div class=\"wrapper\"><div class=\"heading\">Ваш код подтверждения</div><div class=\"footnote\">Введите код ниже в приложение</div><center><div class=\"code\">%s</div></center><div class=\"border\"></div><div class=\"footnote\">Если вы не отправляли этот код - проигнорируйте это сообщение</div></div></div></body>", nil, code)
	to := []string{email}
	err := sender.SendEmail(subject, content, to, nil, nil)
	return err
}

func GenerateRandomCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%04d", rand.Intn(10000))
}

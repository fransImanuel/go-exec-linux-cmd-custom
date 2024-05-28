package mail

import (
	"fmt"
	"go-exec-linux-cmd-custom/dto"

	"gopkg.in/gomail.v2"
)

type SMTP struct {
	Config *dto.SMTPConfig
}

func InitEmail(config *dto.SMTPConfig) *SMTP {
	return &SMTP{
		Config: config,
	}
}

func (e *SMTP) Send(to []string, cc []string, bcc []string, subject string, bodyType string, body string, attachment []string) error {
	fmt.Println("Email - Send - starting...")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.Config.Email)
	mailer.SetHeader("To", to...)

	if cc != nil && len(cc) > 0 {
		mailer.SetHeader("Cc", cc...)
	}

	if bcc != nil && len(bcc) > 0 {
		mailer.SetHeader("Bcc", bcc...)
	}

	mailer.SetHeader("Subject", subject)
	mailer.SetBody(bodyType, body)

	if attachment != nil && len(attachment) > 0 {
		for _, path := range attachment {
			mailer.Attach(path)
		}
	}

	dialer := gomail.NewDialer(
		e.Config.Host,
		e.Config.Port,
		e.Config.Email,
		e.Config.Password,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		fmt.Println("Email - Send - error: ", err)
		return err
	}

	fmt.Println("***************************EMAIL SENT***************************")
	fmt.Println("Email - Send - finished")
	fmt.Println("***************************EMAIL SENT***************************")
	return nil
}

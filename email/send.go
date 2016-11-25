package email

import (
	"github.com/mailgun/mailgun-go"
	"github.com/zuzuleinen/dave/config"
	"log"
)

//Sends an HTML body e-mail
func Send(from, subject, body string) {
	mailGunConfig := config.Config()

	mg := mailgun.NewMailgun(
		mailGunConfig.Domain,
		mailGunConfig.PrivateApiKey,
		mailGunConfig.PublicApiKey,
	)

	m := mg.NewMessage(from, subject, "", config.YourEmail())
	m.SetHtml(body)

	_, _, err := mg.Send(m)

	if err != nil {
		log.Fatal(err)
	}
}

package email

import (
	"database/sql"
	"github.com/mailgun/mailgun-go"
	"github.com/zuzuleinen/jobber/config"
	"log"
)

//Sends an HTML body e-mail
func Send(from, subject, body string, db *sql.DB) {
	mailGunConfig := config.Config(db)

	mg := mailgun.NewMailgun(
		mailGunConfig.Domain,
		mailGunConfig.PrivateApiKey,
		mailGunConfig.PublicApiKey,
	)

	m := mg.NewMessage(from, subject, "", from)
	m.SetHtml(body)

	_, _, err := mg.Send(m)

	if err != nil {
		log.Fatal(err)
	}
}

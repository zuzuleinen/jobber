package config

import (
	"database/sql"
	"github.com/zuzuleinen/jobber/database"
)

type MailGunConfig struct {
	Domain        string
	PrivateApiKey string
	PublicApiKey  string
}

func Config(db *sql.DB) MailGunConfig {
	m := database.Mailgun(db)

	config := MailGunConfig{Domain: m.Domain, PrivateApiKey: m.PrivKey, PublicApiKey: m.PubKey}
	return config
}

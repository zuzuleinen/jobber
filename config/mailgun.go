package config

type MailGunConfig struct {
	Domain        string
	PrivateApiKey string
	PublicApiKey  string
}

func Config() MailGunConfig {
	config := MailGunConfig{
		"yourdomain",
		"your priv key",
		"your pub key",
	}
	return config
}

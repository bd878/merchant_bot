package merchant_bot

type Config struct {
	PGConn           string     `json:"pg_conn"`
	WebhookPath      string     `json:"webhook_path"`
	DBPath           string     `json:"db_path"`
	WebhookURL       string     `json:"webhook_url"`
	Addr             string     `json:"addr"`
	RootDir          string     `json:"root_dir"`
}
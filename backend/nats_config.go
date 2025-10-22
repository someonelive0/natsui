package backend

type NatsConfig struct {
	Url      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
	Timeout  int    `json:"timeout"`
}

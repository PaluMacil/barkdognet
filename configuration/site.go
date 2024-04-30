package configuration

type Site struct {
	Name          string `koanf:"name"`
	ListenAddr    string `koanf:"listen_addr"`
	ListenPort    string `koanf:"listen_port"`
	BaseURL       string `koanf:"base_url"`
	LiveTemplates bool   `koanf:"live_templates"`
}

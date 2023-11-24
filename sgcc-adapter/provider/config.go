package provider

type config struct {
	CloudServerURL string    `yaml:"cloud-server-url"`
	UeServerURL    string    `yaml:"ue-server-url"`
	Register       *register `yaml:"register"`
}

type register struct {
	Maintainer string   `yaml:"maintainer"`
	Url        string   `yaml:"url"`
	Stations   []string `yaml:"stations"`
}

var Config = &config{}

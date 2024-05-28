package config

type config struct {
	CloudServerURL        string    `yaml:"cloud-server-url"`
	Register              *register `yaml:"register"`
	PeerConnectionOptions string    `json:"peerConnectionOptions" yaml:"peerConnectionOptions"`
}

type register struct {
	Maintainer string   `yaml:"maintainer"`
	Url        string   `yaml:"url"`
	Stations   []string `yaml:"stations"`
}

var Config = &config{}

package model

type MetaData struct {
	Labels map[string]string `yaml:"labels"`
}

type PakConfig struct {
	Paks []Pak `yaml:"paks"`
}

type Pak struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PlayerConfig struct {
	MatchViewportRes bool   `json:"matchViewportRes"`
	HideUI           bool   `json:"hideUI"`
	IdleDisconnect   bool   `json:"idleDisconnect"`
	IdleTimeout      uint32 `json:"idleTimeout"`
}

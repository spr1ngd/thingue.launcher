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

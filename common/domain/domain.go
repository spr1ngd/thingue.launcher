package domain

type VersionInfo struct {
	Version   string
	GitCommit string
	BuildDate string
}

type DeviceInfo struct {
	Version    string   `json:"version"`
	Workdir    string   `json:"workdir"`
	Hostname   string   `json:"hostname"`
	Memory     string   `json:"memory"`
	Cpus       []string `json:"cpus"`
	Gpus       []string `json:"gpus"`
	Ips        []string `json:"ips"`
	OsArch     string   `json:"osArch"`
	OsType     string   `json:"osType"`
	SystemUser string   `json:"systemUser"`
}

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

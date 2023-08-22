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

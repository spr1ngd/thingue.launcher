package model

type DeviceInfo struct {
	Version    string      `json:"version"`
	Workdir    string      `json:"workdir"`
	Hostname   string      `json:"hostname"`
	Memory     string      `json:"memory"`
	Cpus       StringSlice `json:"cpus"`
	Gpus       StringSlice `json:"gpus"`
	Ips        StringSlice `json:"ips"`
	OsArch     string      `json:"osArch"`
	OsType     string      `json:"osType"`
	SystemUser string      `json:"systemUser"`
}

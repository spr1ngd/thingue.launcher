package model

type DeviceInfo struct {
	ID         string
	Version    string
	Workdir    string
	Hostname   string
	Memory     string
	Cpus       []string
	Gpus       []string
	Ips        []string
	OsArch     string
	OsType     string
	SystemUser string
}

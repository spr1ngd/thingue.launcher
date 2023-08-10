package message

type ControlMsg struct {
	InstanceName string `mapstructure:"instanceName"`
	Type         string `mapstructure:"type"`
}

type UpdateMsg struct {
}

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

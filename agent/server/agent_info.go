package server

type AgentInfo struct {
	ID         string
	Version    string
	Workdir    string
	Hostname   string
	Memory     string
	Cpus       []string
	Ips        []string
	OsArch     string
	OsType     string
	SystemUser string
}

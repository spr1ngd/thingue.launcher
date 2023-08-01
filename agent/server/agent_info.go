package server

type AgentInfo struct {
	ID        string
	Version   string
	Path      string
	Hostname  string
	TotalMem  string
	Cpus      uint
	Ips       []string
	OsArch    string
	OsVersion string
	OsUser    string
}

package domain

type IceServer struct {
	URLs       []string `json:"urls" yaml:"urls"`
	Username   string   `json:"username" yaml:"username"`
	Credential string   `json:"credential" yaml:"credential"`
}

type PeerConnectionOptions struct {
	IceServers []IceServer `json:"iceServers" yaml:"iceServers"`
}

type VersionInfo struct {
	Version   string
	GitCommit string
	BuildDate string
}

type SyncConfig struct {
}

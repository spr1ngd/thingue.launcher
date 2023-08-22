package domain

type Instance struct {
	CID               uint     `json:"cid"` //客户端ID
	SID               string   `json:"sid"` //服务端ID
	Name              string   `json:"name"`
	ExecPath          string   `json:"execPath"`
	LaunchArguments   []string `json:"launchArguments"`
	Metadata          string   `json:"metadata"`
	PaksConfig        string   `json:"paksConfig"`
	FaultRecover      bool     `json:"faultRecover"`
	Pid               int      `json:"pid"`
	StateCode         int8     `json:"stateCode"`
	StreamerConnected bool     `json:"streamerConnected"`
}

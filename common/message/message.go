package message

type ControlMsg struct {
	InstanceName string `mapstructure:"instanceName"`
	Type         string `mapstructure:"type"`
}

type NodeProcessControlMsg struct {
	Type    string
	ID      uint   `json:"id"`
	Command string `json:"command"`
}

type UpdateMsg struct {
}

type ProcessStateUpdate struct {
	SID       string `json:"sid"`
	StateCode int8   `json:"stateCode"`
}

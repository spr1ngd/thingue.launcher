package request

type PublishJson struct {
	Topic   string         `json:"topic"`
	Payload map[string]any `json:"payload"`
	Retain  bool           `json:"retain"`
	Qos     byte           `json:"qos"`
}

type PublishText struct {
	Topic  string `json:"topic"`
	Text   string `json:"text"`
	Retain bool   `json:"retain"`
	Qos    byte   `json:"qos"`
}

type SelectorCond struct {
	StreamerId        string `json:"sid"`
	Name              string `json:"name"`
	PlayerCount       *int   `json:"playerCount"`
	LabelSelector     string `json:"labelSelector"`
	StreamerConnected bool   `json:"streamerConnected"`
}

type ProcessControl struct {
	StreamerId string `json:"streamerId"`
	Command    string `json:"command"`
}

type PakControl struct {
	StreamerId string `json:"streamerId"`
	Type       string `json:"type"`
	PakName    string `json:"pakName"`
}

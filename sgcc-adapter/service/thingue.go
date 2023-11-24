package service

import (
	"encoding/json"
	"errors"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
	"thingue-launcher/sgcc-adapter/provider"
)

type thingue struct {
	*BaseRequest
}

var ThingUE = thingue{}

func (s *thingue) Init() {
	s.BaseRequest = NewBaseRequest(provider.Config.UeServerURL)
}

func (s *thingue) GetInstances() ([]*model.ServerInstance, error) {
	if result, err := s.HttpGet("/api/instance/instanceList"); err == nil {
		responseBody := response.Response[response.PageResult[*model.ServerInstance]]{}
		err = json.Unmarshal(result, &responseBody)
		return responseBody.Data.List, nil
	} else {
		return nil, err
	}
}

func (s *thingue) GetMaxNode() int {
	if instances, err := s.GetInstances(); err == nil {
		return len(instances)
	} else {
		return 0
	}
}

func (s *thingue) SendPrecessCommand(sid string, command string) {
	control := request.ProcessControl{
		SID:     sid,
		Command: command,
	}
	data, _ := json.Marshal(control)
	_, _ = s.HttpPost("/api/instance/processControl", data)
}

func (s *thingue) SendPakControl(sid, controlType, pak string) {
	control := request.PakControl{
		SID:  sid,
		Type: controlType,
		Pak:  pak,
	}
	data, _ := json.Marshal(control)
	_, _ = s.HttpPost("/api/instance/pakControl", data)
}

func (s *thingue) GetInstanceBySID(sid string) (*model.ServerInstance, error) {
	cond := request.SelectorCond{
		SID:               sid,
		StreamerConnected: true,
	}
	data, _ := json.Marshal(cond)
	if result, err := s.HttpPost("/api/instance/instanceSelect", data); err == nil {
		responseBody := response.Response[[]*model.ServerInstance]{}
		err = json.Unmarshal(result, &responseBody)
		if responseBody.Code == 200 && len(responseBody.Data) > 0 {
			return responseBody.Data[0], nil
		} else {
			return nil, errors.New("not found")
		}
	} else {
		return nil, err
	}
}

package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/bluele/gcache"
	"math"
	"thingue-launcher/common/message"
	"thingue-launcher/common/message/types"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/server/core/provider"
	"thingue-launcher/server/global"
	"time"
)

type nodeService struct {
	//BufferCache map[string]*bytes.Buffer
	WsIdMap     map[string]int
	BufferCache gcache.Cache
}

var NodeService = nodeService{
	//BufferCache: make(map[string]*bytes.Buffer),
	WsIdMap:     make(map[string]int),
	BufferCache: gcache.New(math.MaxInt64).LRU().Build(),
}

func (s *nodeService) NodeRegister(registerInfo *request.NodeRegisterInfo) error {
	var node model.Node
	global.SERVER_DB.Find(&node, registerInfo.NodeID)
	node.SetDeviceInfo(*registerInfo.DeviceInfo)
	var serverInstances = make([]*model.ServerInstance, 0)
	for _, instance := range registerInfo.Instances {
		serverInstances = append(serverInstances, instance.ToServerInstance())
	}
	node.Instances = serverInstances
	global.SERVER_DB.Save(&node)
	provider.AdminConnProvider.BroadcastUpdate()
	return nil
}

func (s *nodeService) NodeList() []model.Node {
	var nodeList []model.Node
	global.SERVER_DB.Preload("Instances").Find(&nodeList)
	return nodeList
}

func (s *nodeService) NodeOnline(node *model.Node) {
	global.SERVER_DB.Create(&node)
}

func (s *nodeService) NodeOffline(node *model.Node) {
	global.SERVER_DB.Delete(&node)
	global.SERVER_DB.Where("node_id = ?", node.ID).Delete(&model.ServerInstance{})
	provider.AdminConnProvider.BroadcastUpdate()
}

func (s *nodeService) GetInstanceSid(nodeId string, instanceId string) (string, error) {
	var instance model.ServerInstance
	err := global.SERVER_DB.Where("node_id = ? AND c_id = ?", nodeId, instanceId).First(&instance).Error
	if err == nil {
		return instance.SID, err
	} else {
		return "", err
	}
}

func (s *nodeService) CollectLogs(req request.LogsCollect) error {
	s.WsIdMap[req.TraceId] = req.WsId
	return provider.NodeConnProvider.SendToNode(req.NodeId, &message.Message{
		Type: types.ServerCollectNodeLogs,
		Data: req.TraceId,
	})
}

func (s *nodeService) UploadLogs(traceId string, buf *bytes.Buffer) error {
	err := s.BufferCache.SetWithExpire(traceId, buf, time.Second*6)
	if err != nil {
		fmt.Println(err)
		return err
	}
	provider.AdminConnProvider.NotifyDownloadComplete(s.WsIdMap[traceId])
	delete(s.WsIdMap, traceId)
	return nil
}

func (s *nodeService) DownloadLogs(traceId string) (error, *bytes.Buffer) {
	value, err := s.BufferCache.Get(traceId)
	if err != nil {
		return errors.New("缓冲区数据超时被清除"), nil
	} else {
		return err, value.(*bytes.Buffer)
	}
}

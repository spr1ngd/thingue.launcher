package instance

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/labels"
	"thingue-launcher/common/model"
	"thingue-launcher/server/global"
	"time"
)

type ticketService struct {
	cache *expirable.LRU[string, string]
}

var TicketService = ticketService{
	cache: expirable.NewLRU[string, string](500, nil, time.Second*10),
}

func (s *ticketService) GetTicketByLabelSelector(selectCond model.SelectCond) (string, error) {
	selector, err := labels.Parse(selectCond.Selector)
	if err != nil {
		return "", err
	}
	var instances []model.ServerInstance
	global.SERVER_DB.Find(&instances)
	for _, instance := range instances {
		var metaData MetaData
		err := yaml.Unmarshal([]byte(instance.Metadata), &metaData)
		if err != nil {
			continue
		}
		instance.Labels = labels.Set(metaData.Labels)
		if selector.Matches(instance.Labels) {
			//生成ticket
			ticket, _ := uuid.NewUUID()
			//添加缓存
			s.cache.Add(ticket.String(), instance.SID)
			return ticket.String(), nil
		}
	}
	return "", errors.New("找不到合适的实例")
}

func (s *ticketService) GetTicketById(sid string) (string, error) {
	var instance model.ServerInstance
	err := global.SERVER_DB.Where("s_id = ?", sid).First(&instance).Error
	if err == nil {
		ticket, _ := uuid.NewUUID()
		//添加缓存
		s.cache.Add(ticket.String(), instance.SID)
		return ticket.String(), err
	} else {
		return "", err
	}
}

func (s *ticketService) GetSidByTicket(ticket string) (string, error) {
	sid, ok := s.cache.Get(ticket)
	if ok {
		return sid, nil
	} else {
		return "", errors.New("无效ticket")
	}
}

type MetaData struct {
	Labels map[string]string `yaml:"labels"`
}

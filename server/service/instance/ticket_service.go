package instance

import (
	"errors"
	"github.com/google/uuid"
	"github.com/hashicorp/golang-lru/v2/expirable"
	"k8s.io/apimachinery/pkg/labels"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
	"thingue-launcher/server/global"
	"time"
)

type ticketService struct {
	cache *expirable.LRU[string, string]
}

var TicketService = ticketService{
	cache: expirable.NewLRU[string, string](500, nil, time.Second*10),
}

func (s *ticketService) TicketSelect(selectCond request.TicketSelector) (response.InstanceTicket, error) {
	// 数据库查询
	query := global.SERVER_DB.Where("streamer_connected = ?", true)
	if selectCond.SID != "" {
		query = query.Where("s_id = ?", selectCond.SID)
	}
	if selectCond.Name != "" {
		query = query.Where("name = ?", selectCond.Name)
	}
	if selectCond.PlayerCount != 0 {
		query = query.Where("player_count = ?", selectCond.PlayerCount)
	}
	var serverInstances []model.ServerInstance
	query.Find(&serverInstances)
	ticket := response.InstanceTicket{}
	if len(serverInstances) > 0 {
		if selectCond.LabelSelector != "" {
			// label匹配
			selector, err := labels.Parse(selectCond.LabelSelector)
			if err != nil {
				return ticket, err
			}
			for _, instance := range serverInstances {
				if selector.Matches(instance.Labels) {
					//生成ticket
					ticketId, _ := uuid.NewUUID()
					//添加缓存
					s.cache.Add(ticketId.String(), instance.SID)
					ticket.SetInstanceInfo(&instance)
					ticket.Ticket = ticketId.String()
					return ticket, nil
				}
			}
		} else {
			//不需要label匹配，挑选第一个生成ticket
			ticketId, _ := uuid.NewUUID()
			//添加缓存
			s.cache.Add(ticketId.String(), serverInstances[0].SID)
			ticket.SetInstanceInfo(&serverInstances[0])
			ticket.Ticket = ticketId.String()
			return ticket, nil
		}
	}
	return ticket, errors.New("找不到合适的实例")
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

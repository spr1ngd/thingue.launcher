package service

import (
	"errors"
	"github.com/bluele/gcache"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/labels"
	"math"
	"thingue-launcher/common/model"
	"thingue-launcher/common/request"
	"thingue-launcher/common/response"
	"thingue-launcher/server/global"
	"time"
)

type ticketService struct {
	cache gcache.Cache
}

var TicketService = ticketService{
	cache: gcache.New(math.MaxInt64).LRU().Build(),
}

func (s *ticketService) TicketSelect(selectCond request.TicketSelector) (response.InstanceTicket, error) {
	// 数据库查询
	query := global.SERVER_DB.Where("streamer_connected = ? or auto_control = ?", true, true)
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
					s.cache.SetWithExpire(ticketId.String(), instance.SID, 10*time.Second)
					ticket.SetInstanceInfo(&instance)
					ticket.Ticket = ticketId.String()
					return ticket, nil
				}
			}
		} else {
			//不需要label匹配，挑选第一个生成ticket
			ticketId, _ := uuid.NewUUID()
			//添加缓存
			s.cache.SetWithExpire(ticketId.String(), serverInstances[0].SID, 10*time.Second)
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
		err = s.cache.SetWithExpire(ticket.String(), instance.SID, 10*time.Second)
		if err == nil {
			return ticket.String(), err
		}
	}
	return "", err
}

func (s *ticketService) GetSidByTicket(ticket string) (string, error) {
	sid, err := s.cache.Get(ticket)
	if err == nil {
		return sid.(string), nil
	} else {
		return "", errors.New("ticket无效或过期")
	}
}

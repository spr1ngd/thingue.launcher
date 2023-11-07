package service

import (
	"errors"
	"fmt"
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

func (s *ticketService) TicketSelect(selectCond request.SelectorCond) (response.InstanceTicket, error) {
	ticket := response.InstanceTicket{}
	// 数据库查询
	//query := global.SERVER_DB.Where("state_code = ? or auto_control = ?", 1, true)
	query := global.SERVER_DB
	if selectCond.StreamerConnected == true {
		query = query.Where("streamer_connected = ?", selectCond.StreamerConnected)
	}
	if selectCond.SID != "" {
		query = query.Where("s_id = ?", selectCond.SID)
	}
	if selectCond.Name != "" {
		query = query.Where("name = ?", selectCond.Name)
	}
	if selectCond.PlayerCount != nil && *selectCond.PlayerCount >= 0 {
		query = query.Where("player_count = ?", selectCond.PlayerCount)
	}
	var findInstances []*model.ServerInstance
	query.Find(&findInstances)
	// 判断查询后是否有结果
	if len(findInstances) == 0 {
		return ticket, errors.New("没有匹配的实例")
	}
	// 筛选掉未启动且未开启自动启停的实例
	var readyInstances []*model.ServerInstance
	for _, instance := range findInstances {
		if instance.StateCode == 1 || instance.AutoControl == true {
			readyInstances = append(readyInstances, instance)
		}
	}
	if len(readyInstances) == 0 {
		return ticket, errors.New("实例未启动且未开启自动启停")
	}
	if selectCond.LabelSelector != "" {
		// label匹配
		selector, err := labels.Parse(selectCond.LabelSelector)
		if err != nil {
			return ticket, err // label解析失败
		}
		for _, instance := range readyInstances {
			if selector.Matches(instance.Labels) {
				//生成ticket
				ticketId, _ := uuid.NewUUID()
				//添加缓存
				s.cache.SetWithExpire(ticketId.String(), instance.SID, 10*time.Second)
				ticket.SetInstanceInfo(instance)
				ticket.Ticket = ticketId.String()
				return ticket, nil
			}
		}
		return ticket, errors.New(fmt.Sprintf("找不到符合%s的可用实例", selectCond.LabelSelector))
	} else {
		//不需要label匹配，挑选第一个生成ticket
		ticketId, _ := uuid.NewUUID()
		//添加缓存
		s.cache.SetWithExpire(ticketId.String(), readyInstances[0].SID, 10*time.Second)
		ticket.SetInstanceInfo(readyInstances[0])
		ticket.Ticket = ticketId.String()
		return ticket, nil
	}
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

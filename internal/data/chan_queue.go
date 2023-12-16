package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minicloudsky/lianjia/internal/biz"
)

type ChanQueue struct {
	ErshoufangChan chan []biz.Message
	LoupanChan     chan []biz.Message
	CommercialChan chan []biz.Message
	ZufangChan     chan []biz.Message
	logger         log.Logger
}

func (c *ChanQueue) Send(ctx context.Context, msg []biz.Message, houseType biz.HoueseType) error {
	switch houseType {
	case biz.HouseTypeErshoufang:
		c.ErshoufangChan <- msg
	case biz.HouseTypeLoupan:
		c.LoupanChan <- msg
	case biz.HouseTypeCommercial:
		c.CommercialChan <- msg
	case biz.HouseTypeZufang:
		c.ZufangChan <- msg
	default:
		return errors.New(fmt.Sprintf("unknown houseType %s", houseType))
	}
	return nil
}

func (c *ChanQueue) Receive(ctx context.Context) error {
	helper := log.NewHelper(c.logger)
	helper.Infof("channel start receiving messages...")
	openChannels := 3
	for {
		select {
		case msgs, ok := <-c.ErshoufangChan:
			if ok {
				biz.ErshoufangProcessChan <- msgs
			} else {
				openChannels--
				if openChannels == 0 {
					helper.Infof("all channel closed, exiting receive...")
					return nil
				}
			}
		case msgs, ok := <-c.LoupanChan:
			if ok {
				biz.LoupanProcessChan <- msgs
			} else {
				openChannels--
				if openChannels == 0 {
					helper.Infof("all channel closed, exiting receive...")
					return nil
				}
			}
		case msgs, ok := <-c.CommercialChan:
			if ok {
				biz.CommercialProcessChan <- msgs
			} else {
				openChannels--
				if openChannels == 0 {
					helper.Infof("all channel closed, exiting receive...")
					return nil
				}
			}
			//case msgs, ok := <-c.ZufangChan:
			//	if ok {
			//		biz.ZufangProcessChan <- msgs
			//	}
		}
	}
}

package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go"
	"github.com/apache/rocketmq-client-go/consumer"
	"github.com/apache/rocketmq-client-go/primitive"
	"github.com/apache/rocketmq-client-go/producer"
	"model"
)

func InsertStockJnMQ(mpays []model.MyPay)error{
	mpay,err := json.Marshal(mpays)
	if err != nil{
		return err
	}else{
		p,err := rocketmq.NewProducer(
			producer.WithNameServer([]string{"127.0.0.1:9876"}),
			producer.WithRetry(2),
			producer.WithGroupName("GID_1"),
		)
		if err != nil{
			return err
		}else{
			p.Start()
			_,err := p.SendSync(context.Background(),&primitive.Message{
				Topic:         "bookstore",
				Body:          mpay,
			})
			if err != nil{
				return err
			}else{
				return nil
			}
		}
	}
}

//消息队列对订单系统与库存系统解耦

func GetStockJnMQ(){
	c,err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("GID_2"),
	)
	if err != nil{
		fmt.Println(err)
	}else{
		err := c.Subscribe("bookstore",consumer.MessageSelector{},func(ctx context.Context,msgs...*primitive.MessageExt)(consumer.ConsumeResult,error){
			for i,_ := range msgs{
				mpays := make([]model.MyPay,0)
				err := json.Unmarshal(msgs[i].Body,&mpays)
				if err != nil{
				}else{
					BookStockJn(mpays)
				}
			}
			return consumer.ConsumeSuccess,nil
		})
		if err != nil{
			fmt.Println(err)
		}else{
			for{
				c.Start()
			}
		}
	}
}

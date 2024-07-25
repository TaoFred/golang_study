package usage

import (
	"fmt"
	"time"

	"log"

	"github.com/nats-io/nats.go"
)

func GetNatsConn(addr string) (*nats.Conn, error) {
	// 连接到 NATS 服务器
	nc, err := nats.Connect(addr)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

// 发布消息
func Publish(nc *nats.Conn, subject string, data []byte) error {
	err := nc.Publish(subject, data)
	if err != nil {
		return err
	}
	return nil
}

// 发布-订阅模式
func PubAndSub() {
	// 连接到 NATS 服务器
	nc, err := GetNatsConn("127.0.0.1:9277")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// 订阅消息1
	subscription, err := nc.Subscribe("foo", func(m *nats.Msg) {
		// 处理消息
		fmt.Printf("receive foo subject message1: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatalf("subscribe foo subject failed: %v", err)
		return
	}
	fmt.Printf("subscription: %+v\n", subscription)

	// 订阅消息2
	subscription, err = nc.Subscribe("foo", func(m *nats.Msg) {
		// 处理消息
		fmt.Printf("receive foo subject message2: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatalf("subscribe foo subject failed: %v", err)
		return
	}
	fmt.Printf("subscription: %+v\n", subscription)

	nc.Flush()
	// 发布消息
	err = nc.Publish("foo", []byte("hello foo"))
	if err != nil {
		log.Fatalf("publish foo subject message failed: %v", err)
		return
	}
	fmt.Printf("publish foo subject message success\n")
}

// 请求-响应模式
func ReqAndRep() {
	// 连接到 NATS 服务器
	nc, err := GetNatsConn("127.0.0.1:9277")
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	// 订阅消息
	nc.Subscribe("request.foo", func(m *nats.Msg) {
		// 处理消息
		fmt.Printf("receive request.foo subject message: %s\n", string(m.Data))
		// 回复消息
		m.Respond([]byte("reply request.foo"))
	})
	// 发布请求
	resp, err := nc.Request("request.foo", []byte("hello request.foo"), time.Second*5)
	if err != nil {
		log.Fatalf("request request.foo subject message failed: %v", err)
		return
	}
	fmt.Printf("response: %s\n", string(resp.Data))
}

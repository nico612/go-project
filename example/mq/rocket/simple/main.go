package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"net"
	"os"
	"strconv"
	"strings"
)

func getRocketMQBrokerAddress() (string, error) {
	conn, err := net.Dial("udp", "rocketmq-namesrv:9876")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	parts := strings.Split(localAddr, ":")
	ip := parts[0]

	return ip, nil
}

func main() {

	// 获取 RocketMQ Broker 在容器中的地址
	nameServer := "127.0.0.1:9876"

	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{nameServer}),
		producer.WithRetry(2),
	)

	if err != nil {
		fmt.Printf("new producer error: %s\n", err)
		os.Exit(1)
		return
	}

	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s\n", err)
		os.Exit(1)
		return
	}

	topic := "test_topic1"
	for i := 0; i < 10; i++ {
		msg := primitive.NewMessage(topic,
			[]byte("Hello RocketMQ Go Client!"+strconv.Itoa(i)))

		result, err := p.SendSync(context.Background(), msg)

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", result.String())
		}
	}

	// Shutdown producer
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s\n", err)
	}

}

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	url := "nats://192.168.30.59:4222"
	// Connect to a NATS server
	nc, err := nats.Connect(url, nats.UserInfo("idss", "BDsec2022,,-1234567890"))
	if err != nil {
		fmt.Printf("connect '%s' failed: %s\n", url, err)
		return
	}
	fmt.Printf("connect '%s' success, %s\n", url, nc.ConnectedClusterName())
	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		fmt.Printf("make jetstream '%s' failed: %s\n", url, err)
		return
	}
	fmt.Printf("make jetstream '%s' success, %v\n", url, nc.ConnectedServerId())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// js.Publish(ctx, "test", []byte("Hello JS!"), nil)

	// get existing stream handle
	stream, err := js.Stream(ctx, "test")
	if err != nil {
		fmt.Printf("get stream '%s' failed: %s\n", "test", err)
		return
	}
	fmt.Printf("get stream '%s' success, %v\n", "test", stream.CachedInfo())

	stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
		Name: "consumer1",
		// Durable: "consumer1",
	})

	// retrieve consumer handle from a stream
	consumer, err := stream.Consumer(ctx, "consumer1")
	if err != nil {
		fmt.Printf("Get Consumer '%s' failed: %s\n", "consumer1", err)
		return
	}
	fmt.Printf("Get Consumer '%s' success, %v\n", "consumer1", nc.ConnectedServerName())

	// consume messages from the consumer in callback
	cc, err := consumer.Consume(func(msg jetstream.Msg) {
		meta, _ := msg.Metadata()
		fmt.Printf("Received jetstream messa Stream Sequence  : %v, Consumer Sequence: %v\n",
			meta.Sequence.Stream, meta.Sequence.Consumer)
		fmt.Printf("\tmessage: %s\n", string(msg.Data()))
		msg.Ack()
	})
	if err != nil {
		fmt.Printf("Consume '%s' failed: %s\n", "consumer1", err)
		return
	}
	defer cc.Stop()

	time.Sleep(10 * time.Second)
}

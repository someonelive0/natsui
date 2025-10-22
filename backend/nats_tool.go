package backend

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NatsTool struct {
	Conf *NatsConfig
}

func NewNatsTool(conf *NatsConfig) *NatsTool {
	return &NatsTool{Conf: conf}
}

/*
 * make a nats.Conn client object. remember close it when you don't need it.
 */
func (p *NatsTool) MakeConn() (nc *nats.Conn, err error) {
	if len(p.Conf.User) == 0 {
		nc, err = nats.Connect(strings.Join(p.Conf.Servers, ","))
	} else {
		nc, err = nats.Connect(strings.Join(p.Conf.Servers, ","),
			nats.UserInfo(p.Conf.User, p.Conf.Password))
	}
	if err != nil {
		fmt.Printf("connect '%v' failed: %s\n", p.Conf.Servers, err)
		return nil, err
	}
	fmt.Printf("connect '%v' success, %s\n", p.Conf.Servers, nc.ConnectedClusterName())

	return nc, nil
}

/*
 * make a nats.Jetstream context. remember close js.Conn when you don't need it.
 * jetstream.JetStream is a interface{}
 */
func (p *NatsTool) MakeJetStream() (js jetstream.JetStream, err error) {
	var nc *nats.Conn
	if len(p.Conf.User) == 0 {
		nc, err = nats.Connect(strings.Join(p.Conf.Servers, ","))
	} else {
		nc, err = nats.Connect(strings.Join(p.Conf.Servers, ","),
			nats.UserInfo(p.Conf.User, p.Conf.Password))
	}
	if err != nil {
		fmt.Printf("connect '%v' failed: %s\n", p.Conf.Servers, err)
		return nil, err
	}
	fmt.Printf("connect '%v' success, %s\n", p.Conf.Servers, nc.ConnectedClusterName())

	js, err = jetstream.New(nc)
	if err != nil {
		nc.Close()
		fmt.Printf("make jetstream '%v' failed: %s\n", p.Conf.Servers, err)
		return nil, err
	}

	return js, nil
}

/*
 * get existing stream handle, jetstream.Stream is a interface{}
 */
func (p *NatsTool) GetStream(js jetstream.JetStream, streamName string) (stream jetstream.Stream, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(p.Conf.Timeout)*time.Second)
	defer cancel()

	stream, err = js.Stream(ctx, streamName)
	if err != nil {
		fmt.Printf("get stream '%s' failed: %s\n", "test", err)
		return nil, err
	}
	fmt.Printf("get stream '%s' success, %v\n", "test", stream.CachedInfo())

	return stream, nil
}

/*
 * write msg to stream
 */
func (p *NatsTool) Write2Stream(streamName, data string) error {
	js, err := p.MakeJetStream()
	if err != nil {
		return err
	}
	defer js.Conn().Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(p.Conf.Timeout)*time.Second)
	defer cancel()
	_, err = js.Publish(ctx, streamName, []byte(data))
	if err != nil {
		fmt.Printf("publish to stream '%s' failed: %s\n", streamName, err)
		return err
	}

	return nil
}

/*
 * read msgs from stream
 */
func (p *NatsTool) ReadStream(streamName, msg string) error {
	js, err := p.MakeJetStream()
	if err != nil {
		return err
	}
	defer js.Conn().Close()

	stream, err := p.GetStream(js, streamName)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(p.Conf.Timeout)*time.Second)
	defer cancel()

	stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
		Name: "consumer_natsui",
		// Durable: "consumer_natsui",
	})

	// retrieve consumer handle from a stream
	consumer, err := stream.Consumer(ctx, "consumer1")
	if err != nil {
		fmt.Printf("Get Consumer '%s' failed: %s\n", "consumer1", err)
		return err
	}

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
		return err
	}
	defer cc.Stop()

	return nil
}

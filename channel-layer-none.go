package main

import (
	"log"

	uuid "github.com/satori/go.uuid"

	"github.com/ostcar/geiss/asgi"
)

type ChannelLayer struct {
	prefix   string
	expiry   int
	capacity int
}

func NewChannelLayer(expiry int, prefix string, capacity int) *ChannelLayer {
	// TODO: Work with more then one host
	if expiry == 0 {
		expiry = 60
	}
	if prefix == "" {
		prefix = "asgi:"
	}
	if capacity == 0 {
		capacity = 100
	}
	return &ChannelLayer{prefix: prefix, expiry: expiry, capacity: capacity}
}

func (r *ChannelLayer) NewChannel(channelPrefix string) (channel string, err error) {
	var randChanName string = "channel-123"
	channel = channelPrefix + "-" + randChanName
	return channel, nil // must return a random uuid
}

func (r *ChannelLayer) Send(channel string, message asgi.Message) (err error) {
	messageKey := r.prefix + uuid.NewV4().String()
	channelKey := r.prefix + channel
	log.Println("SEND", messageKey, channelKey)
	return nil
}

func (r *ChannelLayer) Receive(channels []string, block bool) (channel string, message asgi.Message, err error) {
	return "", nil, nil
}

package nochannel

type ChannelLayer struct {
	prefix   string
	expiry   int
	host     string
	capacity int
}

func NewChannelLayer() *ChannelLayer {
	return &ChannelLayer{prefix: "", expiry: 60, host: "", capacity: 20}
}


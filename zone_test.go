package test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newResolver(protocol, addr string) *net.Resolver {
	dialer := net.Dialer{
		Timeout: time.Millisecond * time.Duration(10000),
	}
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.DialContext(ctx, protocol, addr)
		},
	}
}

func TestLocalWwwExampleComA(t *testing.T) {
	resolver := newResolver("udp", "127.0.0.1:53")
	ipAddrs, err := resolver.LookupAddr(context.Background(), "www.example.com.")

	assert.NoError(t, err)
	assert.Equal(t, 1, len(ipAddrs))
	assert.Equal(t, "192.168.1.100", ipAddrs[0])
}

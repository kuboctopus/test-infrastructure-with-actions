package test

import (
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
)

func TestLocalWwwExampleComA(t *testing.T) {

	c := new(dns.Client)

	message := new(dns.Msg)
	message.SetQuestion("www.example.com.", dns.TypeA)

	res, _, err := c.Exchange(message, "127.0.0.1:53")
	assert.NoError(t, err)
	assert.Equal(t, "192.168.1.100", res.Answer[0].(*dns.A).A.To4().String())
}

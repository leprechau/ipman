package ipify

import (
	// core
	"context"
	"net/url"
	"strings"

	// backend definition
	"github.com/leprechau/ipman/common/ip"
)

// Get looks up client address via ipify
func (c *Config) Get(proto ip.IFlag) (string, error) {
	var ctx = context.Background()
	var response = new(Response)
	var url *url.URL
	var err error

	// build url based on requested proto
	switch proto {
	case ip.Inet:
		if url, err = url.Parse(c.v4URL); err != nil {
			return "", err
		}
	case ip.Inet6:
		if url, err = url.Parse(c.v6URL); err != nil {
			return "", err
		}
	}

	// execute the client call
	if err = c.client.Get(ctx, url, "/?format=json", nil, response); err != nil {
		return "", err
	}

	// return the ip payload
	return strings.ToLower(response.IP), nil
}

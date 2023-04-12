package defs

import "net/url"

var (
	BaseUrl = &url.URL{
		Scheme: "https",
		Host:   "adventofcode.com",
	}
)

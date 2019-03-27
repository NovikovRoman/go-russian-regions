package go_russian_regions

import "fmt"

type Client struct {
	regions []*Region
}

func NewClient() *Client {
	c := &Client{}
	c.init()
	return c
}

func (c *Client) Regions() []*Region {
	return c.regions
}

func (c *Client) RegionByCode(code string) (*Region, bool) {
	if len(code) < 2 {
		code = fmt.Sprintf("0%s", code)
	}
	for _, r := range c.regions {
		if r.Code == code {
			return r, true
		}
	}
	return nil, false
}

func (c *Client) DefineRegion(text string) (*Region, bool) {
	for _, r := range c.regions {
		if r.Is(text) {
			return r, true
		}
	}
	return nil, false
}

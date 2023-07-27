package main

import (
	"fmt"
	"time"
)

var dontID = map[string]struct{}{
	"1": {},
	"2": {},
	"4": {},
}

type HealthCheck struct {
	ServiceID string
	status    string
}

type GoMetrClient struct {
	url     string
	timeOut int
}

func NewGoMetrClient(url string, seconds int) *GoMetrClient {
	return &GoMetrClient{url: url, timeOut: seconds}
}

func (g *GoMetrClient) getHealth() HealthCheck {
	if g.GetID() == "1" {
		time.Sleep(2 * time.Second)
	}
	return HealthCheck{
		ServiceID: g.GetID(),
		status:    "",
	}
}

func (g *GoMetrClient) Health() (ok bool) {
	ch := make(chan HealthCheck)
	defer close(ch)
	exit := make(chan struct{})
	defer close(exit)

	go func() {
		health := g.getHealth()

		select {
		case ch <- health:
		case <-exit:
			return
		}
	}()

	select {
	case health := <-ch:
		id := health.ServiceID
		_, ok = dontID[id]
	case <-time.After(time.Duration(g.timeOut) * time.Second):
		fmt.Println("Time out")
		exit <- struct{}{}
		return
	}

	return
}

func (g *GoMetrClient) GetMetrics() string {
	return "gometr.GetMetrics"
}

func (g *GoMetrClient) Ping() error {
	return nil
}

func (g *GoMetrClient) GetID() string {
	return g.url
}

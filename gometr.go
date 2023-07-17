package main

type HealthCheck struct {
	ServiceID string
	status    string
}

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type GoMetrClient struct {
	url     string
	seconds int
}

func (g *GoMetrClient) getHealth() HealthCheck {
	return HealthCheck{
		ServiceID: g.GetID(),
		status:    "",
	}
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

func (g *GoMetrClient) Health() bool {
	id := g.getHealth().ServiceID
	for _, v := range []string{"1", "2", "4"} {
		if v == id {
			return true
		}
	}
	return false
}

func NewGoMetrClient(url string, seconds int) *GoMetrClient {
	return &GoMetrClient{url: url, seconds: seconds}
}

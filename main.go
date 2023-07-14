package main

import "fmt"

func main() {
	checks := GenerateCheck()
	for _, service := range checks {
		if service.status == PassStatus {
			fmt.Println(service.ServiceID)
		}
	}
}

type HealthCheck struct {
	ServiceID int
	status    string
}

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

func GenerateCheck() (checks []HealthCheck) {
	for i := 0; i < 5; i++ {
		check := HealthCheck{
			ServiceID: i,
		}
		check.status = FailStatus
		if i%2 == 0 {
			check.status = PassStatus
		}
		checks = append(checks, check)
	}

	return checks
}

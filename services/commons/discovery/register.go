package discovery

import (
	"fmt"
	"slices"
	"strconv"

	consul "github.com/hashicorp/consul/api"
)

const (
	healthCheckInterval = "10s"
	healthCheckTimeout  = "5s"
)

func Register(id, name string, host, port string, tags []string, meta map[string]string) error {
	config := consul.DefaultConfig()
	config.Address = consulEndpoint

	client, err := consul.NewClient(config)
	if err != nil {
		return fmt.Errorf("service discovery: create consul client error: %v", err)
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("service discovery: convert port to int error: %v", err)
	}

	registration := &consul.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    portInt,
		Address: host,
		Meta:    meta,
	}

	var checks []*consul.AgentServiceCheck
	if slices.Contains(tags, "rest") {
		checks = append(checks, &consul.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%s/healthz", host, port),
			Interval: healthCheckInterval,
			Timeout:  healthCheckTimeout,
		})
	}
	if slices.Contains(tags, "grpc") {
		checks = append(checks, &consul.AgentServiceCheck{
			GRPC:     fmt.Sprintf("%s:%s", host, port),
			Interval: healthCheckInterval,
		})
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		return fmt.Errorf("service discovery: register service failed: %v", err)
	}
	return nil
}

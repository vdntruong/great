package discovery

import (
	"fmt"
	"log"
	"slices"
	"sync"
	"time"

	consul "github.com/hashicorp/consul/api"
)

type ServiceDiscovery struct {
	client *consul.Client

	serviceName string
	serviceTag  string

	instancesLock sync.RWMutex
	instances     []string
}

func NewServiceDiscovery(name string, tag string) (*ServiceDiscovery, error) {
	config := consul.DefaultConfig()
	config.Address = consulEndpoint

	client, err := consul.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("service discovery: error creating consul client: %v", err)
	}

	sd := &ServiceDiscovery{
		client:      client,
		serviceName: name,
		serviceTag:  tag,
		instances:   make([]string, 0),
	}

	go sd.watchServiceChanges()
	return sd, nil
}

func (sd *ServiceDiscovery) watchServiceChanges() {
	params := &consul.QueryOptions{
		WaitIndex: 0,
	}

	for {
		services, meta, err := sd.client.Health().Service(sd.serviceName, sd.serviceTag, true, params)
		if err != nil {
			log.Printf("Error watching services: %s", err)
			time.Sleep(time.Second)
			continue
		}

		params.WaitIndex = meta.LastIndex

		sd.instancesLock.Lock()
		sd.instances = make([]string, len(services))
		for i, service := range services {
			protocol, ok := service.Service.Meta["protocol"]
			if !ok {
				protocol = "http"
			}

			if slices.Contains(service.Service.Tags, "grpc") {
				sd.instances[i] = fmt.Sprintf("%s:%d", service.Service.Address, service.Service.Port)
				continue
			}

			sd.instances[i] = fmt.Sprintf("%s://%s:%d", protocol,
				service.Service.Address,
				service.Service.Port)
		}
		sd.instancesLock.Unlock()
	}
}

func (sd *ServiceDiscovery) GetServiceURL() (string, error) {
	sd.instancesLock.RLock()
	defer sd.instancesLock.RUnlock()

	if len(sd.instances) == 0 {
		return "", fmt.Errorf("no healthy instances available")
	}

	// Simple round-robin selection
	// You could implement more sophisticated load balancing here
	instance := sd.instances[time.Now().UnixNano()%int64(len(sd.instances))]
	log.Println("instances:", sd.instances)

	return instance, nil
}

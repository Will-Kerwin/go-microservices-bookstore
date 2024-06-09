package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	consul "github.com/hashicorp/consul/api"
)

// define a consul registy
type Registry struct {
	client *consul.Client
}

// create a new instance of the registry with the provided address
// Return the pointer to an isntance and error if any
func NewRegistry(address string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = address

	client, err := consul.NewClient(config)

	if err != nil {
		return nil, err
	}

	return &Registry{client: client}, nil
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}

// create a service record in discovery
func (registry *Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
	hpParts := strings.Split(hostPort, ":")

	if len(hpParts) != 2 {
		return errors.New("invalid host:port format provided. Eg. localhost:8081")
	}

	port, err := strconv.Atoi(hpParts[1])

	if err != nil {
		return err
	}

	host := hpParts[0]

	err = registry.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: host,
		Port:    port,
		ID:      instanceID,
		Name:    serviceName,
		Check:   &consul.AgentServiceCheck{CheckID: instanceID, TTL: "5s"},
	})

	return err
}

// deregister a service record from discovery
func (registry *Registry) Deregister(ctx context.Context, instanceID string, _ string) error {
	err := registry.client.Agent().ServiceDeregister(instanceID)
	return err
}

// Health Check push to update status of service instance
func (registry *Registry) HealthCheck(instanceID string, _ string) error {
	err := registry.client.Agent().UpdateTTL(instanceID, "", "pass")
	return err
}

// Discover a list of active addresses of a given service name

func (registry *Registry) Discover(ctx context.Context, serviceName string) ([]string, error) {
	enteries, _, err := registry.client.Health().Service(serviceName, "", true, nil)

	if err != nil {
		return nil, err
	}

	var instaces []string

	for _, entry := range enteries {
		instaces = append(instaces, fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port))
	}

	return instaces, nil
}

package common

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"time"

	consul "github.com/hashicorp/consul/api"
)

//Client provides an interface for getting data out of Consul
type Client interface {
	// Get a Service from consul
	Service(string, string) ([]string, error)
	// Register a service with local agent
	Register(string, int) error
	// Deregister a service with local agent
	DeRegister(string) error
}

type ConsulClient struct {
	consul *consul.Client
}

//NewConsulClient returns a Client interface for given consul address
func NewConsulClient() (*ConsulClient, error) {
	config := consul.DefaultConfig()
	addr := os.Getenv("CONSUL_HOST")
	if len(addr) == 0 {
		addr = "127.0.0.1:8500"
	}
	config.Address = addr
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &ConsulClient{consul: c}, nil
}

// Get preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}

// Register a service with consul local agent
func (c *ConsulClient) Register(name string, port int) error {
	ipAddress := GetOutboundIP()
	log.Println("Register with IP:" + ipAddress)
	check := &consul.AgentServiceCheck{
		HTTP:                           "http://" + ipAddress + ":8080/health",
		Interval:                       "15s",
		TLSSkipVerify:                  true,
		DeregisterCriticalServiceAfter: "15s",
	}

	var idService string
	idService = name + ipAddress

	var listTags []string
	listTags = append(listTags, "traefik.backend="+name)
	listTags = append(listTags, "traefik.frontend.rule=Host:"+name+".localhost")
	listTags = append(listTags, "metrics")

	reg := &consul.AgentServiceRegistration{
		ID:      idService,
		Name:    name,
		Address: ipAddress,
		Port:    port,
		Check:   check,
		Tags:    listTags,
	}

	return c.consul.Agent().ServiceRegister(reg)
}

// DeRegister a service with consul local agent
func (c *ConsulClient) DeRegister(name string) error {
	return c.consul.Agent().ServiceDeregister(name)
}

// Service return a service
func (c *ConsulClient) Service(service string, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	passingOnly := true
	addrs, meta, err := c.consul.Health().Service(service, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, err
	}
	return addrs, meta, nil
}

// ConsulManagement about Consul
func ConsulManagement(name string) {
	client, err := NewConsulClient()
	if err != nil {
		fmt.Println("Erreur in consul connexion: ", err)
	}
	client.Register(name, 8080)

	//deregister when Ctrl+C
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		client.DeRegister(name)
		os.Exit(1)
	}()

	go func() {
		addrs, _, err := client.Service("player", "*")
		if err != nil {
			fmt.Println("Erreur in consul list services: ", err)
		}
		for _, addr := range addrs {
			log.Println(addr.Service.Service + "@" + addr.Service.Address)
		}
		time.Sleep(1000 * time.Millisecond)
	}()
}

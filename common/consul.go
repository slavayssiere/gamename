package common

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"time"

	"strconv"

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

//newConsulClient returns a Client interface for given consul address
func newConsulClient() (*ConsulClient, error) {
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
	var clt []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				clt = append(clt, ipnet.IP.String())
				//os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}

	return clt[0]
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

var client *ConsulClient

// ConsulManagement about Consul
func ConsulManagement(name string) (client *ConsulClient) {
	client, err := newConsulClient()
	if err != nil {
		fmt.Println("Erreur in consul connexion: ", err)
	}
	client.Register(name, 8080)

	//deregister when Ctrl+C && exit
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		client.DeRegister(name)
		os.Exit(1)
	}()

	return
}

type ServicesAdresses []string

var ListServices = struct {
	sync.RWMutex
	m map[string]ServicesAdresses
}{m: make(map[string]ServicesAdresses)}

func GetIpForService(name string) (ret string) {
	ListServices.RLock()
	if ListServices.m[name] != nil {
		ret = ListServices.m[name][rand.Intn(len(ListServices.m[name]))]
	}
	ListServices.RUnlock()

	return
}

func ListenService(name string, client *ConsulClient) {
	go func(name string, client *ConsulClient) {
		for {
			addrs, _, err := client.consul.Catalog().Service(name, "", nil)
			if err != nil {
				log.Println("Erreur in consul list services: ", err)
			}
			var listIps ServicesAdresses

			for _, addr := range addrs {
				listIps = append(listIps, addr.ServiceAddress+":"+strconv.Itoa(addr.ServicePort))
			}
			ListServices.Lock()
			ListServices.m[name] = listIps
			ListServices.Unlock()
			time.Sleep(15000 * time.Millisecond)
		}
	}(name, client)
}

package consul

import (
	"strings"
	"sync"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type instanceGroup struct {
	Dc        string
	Instances []*api.CatalogService
}

func FetchDatacenters(env string) ([]string, error) {
	client, err := api.NewClient(consulConfigForEnv(env))
	if err != nil {
		cobra.CheckErr(err)
	}

	return client.Catalog().Datacenters()
}

// TODO: move Map into its own type to improve readability
func GetServicesByDc(dc string) map[string][]string {
	client, err := api.NewClient(consulConfigForDc(dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	q := &api.QueryOptions{
		Datacenter: dc,
	}

	var svcs map[string][]string
	if svcs, _, err = client.Catalog().Services(q); err != nil {
		cobra.CheckErr(err)
	}

	return svcs
}

func GetServicesByEnv(env string) map[string][]string {
	dcs, err := FetchDatacenters(env)
	if err != nil {
		cobra.CheckErr(err)
	}

	ch := make(chan map[string][]string)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(dcs))

		for _, dc := range dcs {
			go listServicesByEnvAsync(dc, ch, &wg)
		}

		wg.Wait()
		close(ch)
	}()

	svcs := make(map[string][]string)
	for res := range ch {
		for k, v := range res {
			svcs[k] = v
		}
	}

	return svcs
}

func listServicesByEnvAsync(dc string, ch chan<- map[string][]string, wg *sync.WaitGroup) {
	defer wg.Done()

	ch <- GetServicesByDc(dc)
}

func ListAllInstancesByEnv(env string, svc string) []*instanceGroup {
	dcs, err := FetchDatacenters(env)
	if err != nil {
		cobra.CheckErr(err)
	}

	ch := make(chan *instanceGroup)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(dcs))

		for _, dc := range dcs {
			go listInstancesByDcAsync(dc, svc, ch, &wg)
		}

		wg.Wait()
		close(ch)
	}()

	var ists []*instanceGroup
	for res := range ch {
		ists = append(ists, res)
	}

	return ists
}

func listInstancesByDcAsync(dc string, svc string, ch chan<- *instanceGroup, wg *sync.WaitGroup) {
	defer wg.Done()

	instances, err := FetchInstancesByDc(dc, svc)
	if err != nil {
		cobra.CheckErr(err)
	}

	ch <- &instanceGroup{
		Dc:        dc,
		Instances: instances,
	}
}

func FetchInstancesByDc(dc string, svc string) ([]*api.CatalogService, error) {
	client, err := api.NewClient(consulConfigForDc(dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	q := &api.QueryOptions{
		Datacenter: dc,
	}

	svcs, _, err := client.Catalog().Service(svc, "", q)

	return svcs, err
}

func ListInstancesByDc(dc string, svc string) []*api.CatalogService {
	instances, err := FetchInstancesByDc(dc, svc)
	if err != nil {
		cobra.CheckErr(err)
	}

	return instances
}

// HTTP client configuration

func consulConfigForEnv(env string) *api.Config {
	var consulUrl string

	if env == "dev" {
		consulUrl = viper.GetString("consul_dev_url")
	} else {
		consulUrl = viper.GetString("consul_prod_url")
	}

	return &api.Config{
		Address:   consulUrl,
		Scheme:    "http",
		Transport: cleanhttp.DefaultTransport(),
	}
}

func consulConfigForDc(dc string) *api.Config {
	var consulUrl string

	if strings.Contains(dc, "dev") {
		consulUrl = viper.GetString("consul_dev_url")
	} else if strings.Contains(dc, "prod") {
		consulUrl = viper.GetString("consul_prod_url")
	} else {
		err := "invalid datacenter"
		cobra.CheckErr(err)
	}

	return &api.Config{
		Address:   consulUrl,
		Scheme:    "http",
		Transport: cleanhttp.DefaultTransport(),
	}
}

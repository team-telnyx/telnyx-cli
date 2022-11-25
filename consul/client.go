/*
Copyright © Telnyx LLC

*/

package consul

import (
	"strings"
	"sync"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type dcInstances struct {
	Dc        string
	Instances []*api.ServiceEntry
}

func FetchDatacenters(env string) ([]string, error) {
	client, err := api.NewClient(consulConfigForEnv(env))
	if err != nil {
		cobra.CheckErr(err)
	}

	return client.Catalog().Datacenters()
}

// TODO: move Map into its own type to improve readability
// TODO: refactor function to use cache
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
	svcs := make(map[string][]string)

	if cacheExists(env) {
		svcNames := readCache(env)

		for _, n := range svcNames {
			svcs[n] = nil
		}

		return svcs
	}

	dcs, err := FetchDatacenters(env)
	if err != nil {
		cobra.CheckErr(err)
	}

	ch := make(chan map[string][]string)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(dcs))

		for _, dc := range dcs {
			go func(dc string, ch chan<- map[string][]string, wg *sync.WaitGroup) {
				defer wg.Done()

				ch <- GetServicesByDc(dc)
			}(dc, ch, &wg)
		}

		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		for k, v := range res {
			svcs[k] = v
		}
	}

	saveCache(env, svcs)

	return svcs
}

func GetInstancesByEnv(env string, svc string) []*dcInstances {
	dcs, err := FetchDatacenters(env)
	if err != nil {
		cobra.CheckErr(err)
	}

	ch := make(chan *dcInstances)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(dcs))

		for _, dc := range dcs {
			go fetchInstancesByDcAsync(dc, svc, ch, &wg)
		}

		wg.Wait()
		close(ch)
	}()

	var ists []*dcInstances
	for res := range ch {
		ists = append(ists, res)
	}

	return ists
}

func fetchInstancesByDcAsync(dc string, svc string, ch chan<- *dcInstances, wg *sync.WaitGroup) {
	defer wg.Done()

	instances := GetInstancesByDc(dc, svc)

	ch <- &dcInstances{
		Dc:        dc,
		Instances: instances,
	}
}

func GetInstancesByDc(dc string, svc string) []*api.ServiceEntry {
	client, err := api.NewClient(consulConfigForDc(dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	q := &api.QueryOptions{
		Datacenter: dc,
	}

	instances, _, err := client.Health().Service(svc, "", false, q)
	if err != nil {
		cobra.CheckErr(err)
	}

	return instances
}

func EnableInstanceMaintenance(svc *api.ServiceEntry, dc string) error {
	client, err := api.NewClient(consulConfigForInstance(svc, dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	return client.Agent().EnableServiceMaintenance(svc.Service.ID, "telnyx-cli")
}

func DisableInstanceMaintenance(svc *api.ServiceEntry, dc string) error {
	client, err := api.NewClient(consulConfigForInstance(svc, dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	return client.Agent().DisableServiceMaintenance(svc.Service.ID)
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

func consulConfigForInstance(svc *api.ServiceEntry, dc string) *api.Config {
	port := consulPort(dc)
	addr := strings.Join([]string{svc.Node.Meta["host-ip"], port}, ":")

	return &api.Config{
		Address:   addr,
		Scheme:    "http",
		Transport: cleanhttp.DefaultTransport(),
	}
}

func consulPort(dc string) string {
	if strings.Contains(dc, "prod") {
		return "28500"
	}
	return "18500"
}

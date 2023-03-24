/*
Copyright © Telnyx LLC

*/

package consul

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type DcInstances struct {
	Dc        string
	Instances []*api.ServiceEntry
	Error     error
}

type ServiceVersion struct {
	Version string `json:"version"`
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

func GetInstancesByEnv(env string, svc string) []*DcInstances {
	dcs, err := FetchDatacenters(env)
	if err != nil {
		cobra.CheckErr(err)
	}

	ch := make(chan *DcInstances)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(dcs))

		for _, dc := range dcs {
			go fetchInstancesByDcAsync(dc, svc, ch, &wg)
		}

		wg.Wait()
		close(ch)
	}()

	var ists []*DcInstances
	for res := range ch {
		ists = append(ists, res)
	}

	return ists
}

func fetchInstancesByDcAsync(dc string, svc string, ch chan<- *DcInstances, wg *sync.WaitGroup) {
	defer wg.Done()

	instances, err := GetInstancesByDc(dc, svc)

	ch <- &DcInstances{
		Dc:        dc,
		Instances: instances,
		Error:     err,
	}
}

func GetInstancesByDc(dc string, svc string) ([]*api.ServiceEntry, error) {
	client, err := api.NewClient(consulConfigForDc(dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	q := &api.QueryOptions{
		Datacenter: dc,
	}

	instances, _, err := client.Health().Service(svc, "", false, q)
	if err != nil {
		return nil, err
	}

	return instances, nil
}

func EnableInstanceMaintenance(svc *api.ServiceEntry, dc, reason string) error {
	client, err := api.NewClient(consulConfigForInstance(svc, dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	return client.Agent().EnableServiceMaintenance(svc.Service.ID, reason)
}

func DisableInstanceMaintenance(svc *api.ServiceEntry, dc string) error {
	client, err := api.NewClient(consulConfigForInstance(svc, dc))
	if err != nil {
		cobra.CheckErr(err)
	}

	return client.Agent().DisableServiceMaintenance(svc.Service.ID)
}

func GetInstanceVersion(svc *api.ServiceEntry) string {
	url := fmt.Sprintf("http://%s:%d/version", svc.Service.Address, svc.Service.Port)
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		cobra.CheckErr(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", "telnyx-cli")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		return "unavailable"
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		cobra.CheckErr(readErr)
	}

	var response ServiceVersion
	if err = json.Unmarshal(body, &response); err != nil {
		return "unavailable"
	}

	return response.Version
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
	addr := strings.Join([]string{consulAgentHost(svc), port}, ":")

	return &api.Config{
		Address:   addr,
		Scheme:    "http",
		Transport: cleanhttp.DefaultTransport(),
	}
}

func consulAgentHost(svc *api.ServiceEntry) string {
	if hostIp, ok := svc.Node.Meta["host-ip"]; ok {
		return hostIp
	} else {
		return svc.Node.Address
	}
}

func consulPort(dc string) string {
	if strings.Contains(dc, "prod") {
		return "28500"
	}
	return "18500"
}

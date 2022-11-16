/*
Copyright © Telnyx LLC

*/

package metaservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// curl -H 'Content-Type: application/json; charset=utf-8' \
//   -X GET 'http://meta.service.ch1b-prod.consul:5984/build/_design/prod-success/_view/prod-success-view?startkey="2022-10-16"'

// {"total_rows":51243,"offset":0,"rows":[
// 	{"id":"dfb850131570bb17953c8ecf50f6f818","key":"2022-11-02T13:55:06.862Z","value":""},
// 	{"id":"dfb850131570bb17953c8ecf50f705c9","key":"2022-11-02T13:53:52.125Z","value":"telephony-backend-consul-kv"},

// TODO: use a generic struct so rows can be of the given type
// see https://itnext.io/how-to-use-golang-generics-with-structs-8cabc9353d75
type Response struct {
	TotalRows int                    `json:"total_rows"`
	Offset    int                    `json:"offset"`
	Rows      []SuccessfulDeployment `json:"rows"`
}

type SuccessfulDeployment struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ResponseTwo struct {
	Bookmark    string       `json:"bookmark"`
	Deployments []Deployment `json:"docs"`
}
type Deployment struct {
	Service         string          `json:"service"`
	Squad           string          `json:"squad"`
	Version         string          `json:"version"`
	StartedByUser   string          `json:"startedByUser"`
	StartTimeStamp  string          `json:"startTimeStamp"`
	GitCommit       string          `json:"gitCommit"`
	BuildNumber     int             `json:"buildNumber"`
	BuildParameters BuildParameters `json:"buildParameters"`
	Duration        int             `json:"duration"`
}

type BuildParameters struct {
	ComponentsToDeploy                  string `json:"componentsToDeploy"`
	ChangeTitle                         string `json:"changeTitle"`
	Type                                string `json:"type"`
	EnableSingleRegionKubernetesCluster string `json:"enableSingleRegionKubernetesCluster"`
	SingleRegionKubernetesCluster       string `json:"singleRegionKubernetesCluster"`
	Description                         string `json:"description"`
	Tags                                string `json:"tags"`
	HostGroup                           string `json:"HOST_GROUP"`
}

func FetchDeployments(svc string, startDate string, endDate string) ([]Deployment, error) {
	successfulDeployments := fetchDeploymentsFromView(startDate, endDate)
	svcDeployments := filterDeploymentsByService(successfulDeployments, svc)

	var deployments []Deployment
	for _, dp := range svcDeployments {
		x := fetchDeploymentById(dp.Id)
		deployments = append(deployments, x)
	}

	return deployments, nil
}

func fetchDeploymentsFromView(startDate string, endDate string) *Response {
	url := fmt.Sprintf(
		"http://meta.service.ch1b-prod.consul:5984/build/_design/prod-success/_view/prod-success-view?startkey=\"%s\"&endKey=\"%s\"&limit=1000",
		startDate,
		endDate,
	)

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
		cobra.CheckErr(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		cobra.CheckErr(readErr)
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		cobra.CheckErr(err)
	}

	return &response
}

func filterDeploymentsByService(response *Response, svc string) []SuccessfulDeployment {
	var deployments []SuccessfulDeployment

	for _, d := range response.Rows {
		if d.Value == svc {
			deployments = append(deployments, d)
		}
	}

	return deployments
}

func fetchDeploymentById(id string) Deployment {
	url := "http://meta.service.ch1b-prod.consul:5984/build/_find"

	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	bodyMap := map[string]interface{}{
		"selector": map[string]string{"_id": id},
		"limit":    1,
		"update":   false,
	}

	body, err := json.Marshal(bodyMap)
	if err != nil {
		cobra.CheckErr(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		cobra.CheckErr(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "telnyx-cli")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		cobra.CheckErr(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		cobra.CheckErr(readErr)
	}

	var response ResponseTwo
	if err = json.Unmarshal(body, &response); err != nil {
		cobra.CheckErr(err)
	}

	return response.Deployments[0]
}

/*
Copyright © Telnyx LLC

Package metaservice implements routines for interacting with infra's
meta service, responsible for storing Jenkins build metadata.

Check https://github.com/team-telnyx/couchdb-build-metadata for more details.
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

type ViewResponse struct {
	TotalRows int                    `json:"total_rows"`
	Offset    int                    `json:"offset"`
	Rows      []SuccessfulDeployment `json:"rows"`
}

type SuccessfulDeployment struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DeploymentResponse struct {
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

func fetchDeploymentsFromView(startDate string, endDate string) *ViewResponse {
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

	var response ViewResponse
	if err = json.Unmarshal(body, &response); err != nil {
		cobra.CheckErr(err)
	}

	return &response
}

func filterDeploymentsByService(response *ViewResponse, svc string) []SuccessfulDeployment {
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

	var response DeploymentResponse
	if err = json.Unmarshal(body, &response); err != nil {
		cobra.CheckErr(err)
	}

	return response.Deployments[0]
}

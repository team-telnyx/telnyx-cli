package consul

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"

	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/config"
	"gopkg.in/yaml.v3"
)

func defaultCacheFolder() string {
	return path.Join(config.DefaultConfigFolder(), config.CacheFolder)
}

func consulCacheFile(env string) string {
	filename := fmt.Sprintf("%s_consul.yml", env)
	return path.Join(defaultCacheFolder(), filename)
}

func cacheExists(env string) bool {
	if _, err := os.Stat(consulCacheFile(env)); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func readCache(env string) []string {
	file, err := ioutil.ReadFile(consulCacheFile(env))
	if err != nil {
		cobra.CheckErr(err)
	}

	data := []string{}
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		cobra.CheckErr(err)
	}

	return data
}

func saveCache(env string, svcs map[string][]string) {
	var svcNames []string
	for k, _ := range svcs {
		svcNames = append(svcNames, k)
	}

	sort.Strings(svcNames)

	data, err := yaml.Marshal(svcNames)
	if err != nil {
		cobra.CheckErr(err)
	}

	err = ioutil.WriteFile(consulCacheFile(env), data, 0777)
	if err != nil {
		cobra.CheckErr(err)
	}
}

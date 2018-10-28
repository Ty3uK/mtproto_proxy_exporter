package stats

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Ty3uK/mtproto_proxy_exporter/config"
)

// Stats represents mtproto_proxy stats data
type Stats struct {
	data []string
}

// GetData obtains mtproto_proxy stats data
func (stats *Stats) GetData(url string) error {
	httpClient := http.Client{
		Timeout: time.Duration(
			time.Duration(config.Config.RequestTimeout) * time.Second,
		),
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("could not prepare new http request: %v", err)
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("could not perform http query: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return fmt.Errorf("could not read http response: %v", err)
	}

	stats.data = strings.Split(string(body), "\n")
	return nil
}

// GetNumberItem finds item in mtproto_proxy stats data
func (stats *Stats) GetNumberItem(name string) (data float64) {
	for _, statItem := range stats.data {
		if strings.HasPrefix(statItem, name) {
			if count, err := strconv.ParseFloat(strings.Split(statItem, "\t")[1], 64); err == nil {
				return count
			}

			return -1
		}
	}

	return -1
}

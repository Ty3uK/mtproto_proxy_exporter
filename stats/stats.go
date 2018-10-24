package stats

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Stats represents mtproto_proxy stats data
type Stats struct {
	data []string
}

// GetData obtains mtproto_proxy stats data
func (stats *Stats) GetData(url string) error {
	response, err := http.Get(url)

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

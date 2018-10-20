package stats

import (
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
func (stats *Stats) GetData(url string) {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	stats.data = strings.Split(string(body), "\n")
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

package alexa

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ALEXA_URL = "http://data.alexa.com/data?cli=10&url="
)

func GetRank(domain string) (*Rank, error) {
	if strings.TrimSpace(domain) == "" {
		return nil, errors.New("url is empty.")
	}
	resp, err := http.Get(fmt.Sprintf("%s%s", ALEXA_URL, domain))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request status error: %v", resp.StatusCode))
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var alexaData XAlexa
	err = xml.Unmarshal(data, &alexaData)
	if err != nil {
		return nil, err
	}

	return &Rank{WorldRank: alexaData.Sd.Reach.Rank,
		CountryCode: alexaData.Sd.Country.Code,
		CountryName: alexaData.Sd.Country.Name,
		CountryRank: alexaData.Sd.Country.Rank,
	}, nil
}

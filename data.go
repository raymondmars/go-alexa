package alexa

import (
	"encoding/json"
	"encoding/xml"
)

type Rank struct {
	WorldRank   int    `json:"world_rank"`   //The domain rank in the world
	CountryCode string `json:"country_code"` //The country (code) to which the domain name belongs
	CountryName string `json:"country_name"` //The country (name) to which the domain name belongs
	CountryRank int    `json:"country_rank"` //The domain rank in this country
}

func (r *Rank) String() string {
	rj, _ := json.Marshal(r)
	return string(rj)
}

type XAlexa struct {
	XMLName xml.Name `xml:"ALEXA"`
	Sd      Xsd      `xml:"SD"`
}
type Xsd struct {
	XMLName xml.Name `xml:"SD"`
	Reach   XReach   `xml:"REACH"`
	Country XCountry `xml:"COUNTRY"`
}
type XReach struct {
	XMLName xml.Name `xml:"REACH"`
	Rank    int      `xml:"RANK,attr"`
}
type XCountry struct {
	XMLName xml.Name `xml:"COUNTRY"`
	Code    string   `xml:"CODE,attr"`
	Name    string   `xml:"NAME,attr"`
	Rank    int      `xml:"RANK,attr"`
}

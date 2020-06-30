# go-alexa
To get alexa rank data

## How to use  
- to get domain rank   
call GetRank function, pass your domain name, after done, a rank struct object will return.      
The rank object contains these information:    
```go 
type Rank struct {
	WorldRank   int    `json:"world_rank"`       //The domain rank in the world
	CountryCode string `json:"country_code"`     //The country (code) to which the domain name belongs
	CountryName string `json:"country_name"`     //The country (name) to which the domain name belongs
	CountryRank int    `json:"country_rank"`     //The domain rank in this country
}
```   


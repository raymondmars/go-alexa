package alexa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRank(t *testing.T) {
	rt, err := GetRank("google.com")
	assert.True(t, err == nil)
	assert.True(t, rt.CountryName == "United States")
	assert.True(t, rt.WorldRank == 1)
	assert.True(t, rt.CountryRank == 1)
}

func TestSomeDomains(t *testing.T) {
	rt, _ := GetRank("amazon.com")
	fmt.Println(rt.String())
	rt1, _ := GetRank("facebook.com")
	fmt.Println(rt1.String())
	rt2, _ := GetRank("twitter.com")
	fmt.Println(rt2.String())
}

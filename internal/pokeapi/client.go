package pokeapi

import (
	"io"
	"fmt"
	"net/http"
	"time"
	"github.com/CaptianRedBeard/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) getFromApi(url string) ([]byte, error) {


	if cachedData, found := c.cache.Get(url); found {
		return cachedData, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data: %s", resp.Status)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, responseData)
	return responseData, nil
}
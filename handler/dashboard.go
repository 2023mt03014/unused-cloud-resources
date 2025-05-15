package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Percentage  int      `json:"percentage"`
	ResourceIDs []string `json:"resource_ids"`
	TotalCount  int      `json:"total_count"`
	UnusedCount int      `json:"unused_count"`
}

var apiEndpoints = []string{
	"http://localhost:9090/aws/lambda",
	"http://localhost:9090/aws/ec2",
	"http://localhost:9090/aws/s3",
}

func fetchAPIData(url string, ch chan<- APIResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		ch <- APIResponse{}
		return
	}
	defer resp.Body.Close()

	var data APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		ch <- APIResponse{}
		return
	}
	ch <- data
}

func DashboardHandler(c *gin.Context) {
	var wg sync.WaitGroup
	ch := make(chan APIResponse, len(apiEndpoints))

	for _, url := range apiEndpoints {
		wg.Add(1)
		go fetchAPIData(url, ch, &wg)
	}

	wg.Wait()
	close(ch)

	var results []APIResponse
	for r := range ch {
		results = append(results, r)
	}

	c.JSON(http.StatusOK, gin.H{
		"dashboard": results,
	})
}

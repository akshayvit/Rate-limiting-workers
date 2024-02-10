package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Market struct {
	MarketId  string `json:"marketId"`
	BestBid   string `json:"bestBid"`
	TimeStamp string `json:"timestamp"`
}

type RTLimiter struct {
	client      *http.Client
	RateLimiter *rate.Limiter
}

func NewClient(rl *rate.Limiter) *RTLimiter {
	return &RTLimiter{
		client:      http.DefaultClient,
		RateLimiter: rl,
	}
}

func (c *RTLimiter) Do(req *http.Request) (*http.Response, error) {
	cont := context.Background()
	err := c.RateLimiter.Wait(cont)
	if err != nil {
		panic(err.Error())
	}
	resp, err1 := c.client.Do(req)
	if err1 != nil {
		panic(err1.Error())
	}
	return resp, nil
}

func worker(id int, jobs <-chan *RTLimiter, results chan<- Market,req *http.Request) {
	
	for client:=range jobs {
		res, err1 := client.Do(req)
	if err1 != nil {
		panic(err1.Error())
	}
	resBody, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		panic(err2.Error())
	}
	var marketBody Market
	err3 := json.Unmarshal(resBody, &marketBody)
	if err3 != nil {
		panic(err3.Error())
	}
	//fmt.Printf("Job %d: The best bid for %s at %s body is %s\n",id, marketBody.MarketId, marketBody.TimeStamp, marketBody.BestBid)
		results<-marketBody
	}
}

func main() {
	rl := rate.NewLimiter(rate.Every(10*time.Second), 5)
	
	reqURL := "https://api.btcmarkets.net/v3/markets/BTC-AUD/ticker"
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		panic(err.Error())
	}
	const jobsn=30
    jobs:=make(chan *RTLimiter,jobsn)
    results:=make(chan Market,jobsn)
	for i := 1; i <= 5; i++ {
		go worker(i,jobs,results,req)
	}
	t1:=time.Now()
   for j:=1;j<=jobsn;j++ {
	   client := NewClient(rl)
       jobs<-client
   }
   
   close(jobs)
   for j:=1;j<=jobsn;j++ {
	marketBody:=<-results
	fmt.Printf("The best bid for %s at %s body is %s\n", marketBody.MarketId, marketBody.TimeStamp, marketBody.BestBid)
       
   }
   t2:=time.Now()
   fmt.Println("Time needed is :,",t2.Sub(t1))
}

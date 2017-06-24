package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	INIT_DELAY      = 3000
	MAX_DELAY       = 6000
	MAX_RETRY       = 4
	DELAY_INCREMENT = 5000
)

type Server struct {
	Name        string
	URL         string
	LastChecked time.Time
	Status      bool
	StatusCode  int
	Delay       int
	Retries     int
	Channel     chan bool
}

var Servers []Server

func (s *Server) checkServer(sc chan *Server) {
	var previousStatus string
	if s.Status == true {
		previousStatus = "OK"
	} else {
		previousStatus = "Down"
	}
	fmt.Println("Server was " + previousStatus + " on last checked - " + s.LastChecked.String())
	fmt.Println("Checking server again:", s.Name)

	resp, err := http.Get(s.URL)
	if err != nil {
		fmt.Println("Error:", err)
		s.Status = false
		s.StatusCode = 0
	} else {
		fmt.Println(resp.Status)
		s.Status = true
		s.StatusCode = resp.StatusCode
	}

	s.LastChecked = time.Now()
	sc <- s
}

func checkServers(sc chan *Server) {
	for i := 0; i < len(Servers); i++ {
		Servers[i].Channel = make(chan bool)
		go Servers[i].checkServer(sc)
		go Servers[i].updateDelay(sc)
	}
}

func (s *Server) updateDelay(sc chan *Server) {
	for {
		select {
		case d := <-s.Channel:

			if d == false {
				s.Delay = s.Delay + DELAY_INCREMENT
				s.Retries++
				if s.Delay >= MAX_DELAY {
					s.Delay = INIT_DELAY
				}
			} else {
				s.Delay = INIT_DELAY
			}
			newDuration := time.Duration(s.Delay)

			if s.Retries >= MAX_RETRY {
				fmt.Println("Server is not reachable after ", MAX_RETRY, " retries.")
			} else {
				fmt.Println("Will check `" + s.Name + "` server again.")
				time.Sleep(newDuration)
				s.checkServer(sc)
			}
		default:
		}
	}
}

func main() {
	sc := make(chan *Server)
	ec := make(chan bool)

	Servers = []Server{
		{Name: "Google", URL: "http://google.com", Status: true, Delay: INIT_DELAY},
		{Name: "Yahoo", URL: "http://yahoo.com", Status: true, Delay: INIT_DELAY},
		{Name: "Amazon", URL: "http://amazon.zom", Status: true, Delay: INIT_DELAY},
	}

	checkServers(sc)

	for {
		select {
		case s := <-sc:
			s.Channel <- false
		default:
		}
	}

	<-ec
}

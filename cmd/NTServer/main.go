package main

import (
	"fmt"
	"time"

	"github.com/technomancers/goNTCore"
)

func main() {
	s, err := goNTCore.NewServer("Test")
	if err != nil {
		panic(err)
	}
	defer s.Close()
	go func() {
		for l := range s.Log {
			if l.Err != nil {
				fmt.Println(l.Err)
			}
			if l.Message != "" {
				fmt.Println(l.Message)
			}
		}
	}()
	s.StartPeriodicClean(5 * time.Second)
	s.Listen()
}

package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/Skarlso/go-furnace/config"
	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

// WaitForFunctionWithStatusOutput waits for a function to complete its action.
func WaitForFunctionWithStatusOutput(state string, freq int, f func()) {
	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan bool)
	go func() {
		defer wg.Done()
		f()
		done <- true
	}()
	go func() {
		counter := 0
		for {
			counter = (counter + 1) % len(Spinners[config.SPINNER])
			fmt.Printf("\r[%s] Waiting for state: %s", yellow(string(Spinners[config.SPINNER][counter])), red(state))
			time.Sleep(time.Duration(freq) * time.Second)
			select {
			case <-done:
				fmt.Println()
				break
			default:
			}
		}
	}()

	wg.Wait()
}

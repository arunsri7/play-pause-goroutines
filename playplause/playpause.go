package playpause

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func Routine(command chan string, response chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	response <- "started"
	for {
		select {
		case cmd := <-command:
			switch cmd {
			case "Stop":
				return
			case "Pause":
				fmt.Println("Trying to Pause")
				response <- "Pasued"
				status = "Pause brooooo"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

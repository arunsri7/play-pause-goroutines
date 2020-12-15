package playpause

import (
	"fmt"
	"time"
)

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func Routine(command chan string, response chan string) {
	var status = "Play"
	response <- "started"
	for {
		select {
		case cmd := <-command:
			switch cmd {
			case "Stop":
				response <- "Stopped"
				return
			case "Pause":
				fmt.Println("Pausing")
				status = "Pause"
				response <- "Paused"
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

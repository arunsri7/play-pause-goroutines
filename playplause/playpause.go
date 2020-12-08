package playpause

import (
	"fmt"
	"sync"
	"time"
)

var i int
var Recieve chan string

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func Routine(command chan string, response chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	fmt.Println("Go")
	response <- "started"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
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

// func PlaypauseFunction(input string) {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	command := make(chan string)
// 	response := make(chan string)
// 	go Routine(command, response, &wg)
// 	for input != "Stop" {
// 		_, err := fmt.Scanln(&input)
// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			return
// 		}
// 		time.Sleep(1 * time.Second)
// 		command <- input
// 	}
// 	wg.Wait()
// }

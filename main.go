package main

import (
	"./changingroom"
	"time"
	"rand"
	"fmt"
)

func randSleep(max int64) {
	time.Sleep(rand.Int63n(max))
}

func main() {
	room := changingroom.NewChangingRoom()

	go func() {
		for i := 0; i < 100; i++ {
			randSleep(1e9)
			go func(i int) {
				room.WomanIn()
				fmt.Printf("Woman %d in room [Status: %v].\n", i, room)
				randSleep(1e10)
				fmt.Printf("Woman %d out of room [Status: %v].\n", i, room)
				room.WomanOut()
			}(i)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			randSleep(1e9)
			go func(i int) {
				room.ManIn()
				fmt.Printf("Man %d in room [Status: %v].\n", i, room)
				randSleep(1e10)
				fmt.Printf("Man %d out of room [Status: %v].\n", i, room)
			}(i)
		}
	}()

	time.Sleep(1e12)

	fmt.Println("Done.")
}
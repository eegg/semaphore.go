package main

import (
	"./twophasebarrier"
	"time"
	"fmt"
)


func main() {
	rvs := twophasebarrier.NewTwoPhaseBarrier([]func(bar *twophasebarrier.TwoPhaseBarrier) {

		func(bar *twophasebarrier.TwoPhaseBarrier) {
			for i := 0; i < 10; i++ {
				bar.StartOfLoop()
				fmt.Printf("One started loop %d\n", i)
				time.Sleep(0.5e9)
				fmt.Printf("One ended loop %d\n", i)
				bar.EndOfLoop()
			}
			bar.End()
		},

		func(bar *twophasebarrier.TwoPhaseBarrier) {
			for i := 0; i < 10; i++ {
				bar.StartOfLoop()
				fmt.Printf("Two started loop %d\n", i)
				time.Sleep(1e9)
				fmt.Printf("Two ended loop %d\n", i)
				bar.EndOfLoop()
			}
			bar.End()
		},

	})

	rvs.Run().P()

	fmt.Println("Done.")
}
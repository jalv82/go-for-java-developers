package main

import "time"

type Ball struct {
	hints uint8
}

func main() {
	table := make(chan *Ball)

	go play("ping", table)
	go play("pong", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<-table
}

func play(s string, table chan *Ball) {
	for {
		ball := <-table
		ball.hints++
		println(s, ball.hints)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

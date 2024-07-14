package main

import "time"

func main() {
	mySleep(1 * time.Second)
}

func mySleep(d time.Duration) {
	<-time.NewTimer(d).C
}

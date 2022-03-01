package main

import (
	"sync"
	"time"
)

// Lets have a shared memory here money which can be accessed by both stingy and spendy
var (
	money = 100
	lock  = sync.Mutex{}
)

// When stingy is working with variable to avoid race condition he wil put a lock on the money
// var and once done he will unlock it
func stingy() {

	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}

	println("Stingy Done!")

}

func spendy() {

	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}

	println("Spendy Done!")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3 * time.Millisecond)
	print(money)
}

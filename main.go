package main

import (
	"fmt"
	"sync"
	"time"
)

const numFilo = 5

var gargo [numFilo]*sync.Mutex
var garcom = sync.NewCond(&sync.Mutex{})
var available = [numFilo]bool{true, true, true, true, true}
var wg sync.WaitGroup

func filosofo(id int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		pensa(id)
		solicitaG(id)
		come(id)
		liberaG(id)
	}
}

func pensa(id int) {
	fmt.Printf("filosofo %d pensando\n", id)
	time.Sleep(time.Millisecond * 1000)
}

func come(id int) {
	fmt.Printf("filosofo %d comendo\n", id)
	time.Sleep(time.Millisecond * 1000)
}

func solicitaG(id int) {
	garcom.L.Lock()
	for !available[id] || !available[(id+1)%numFilo] {
		garcom.Wait()
	}
	available[id] = false
	available[(id+1)%numFilo] = false
	time.Sleep(time.Millisecond * 10)
	garcom.L.Unlock()
}

func liberaG(id int) {
	garcom.L.Lock()
	available[id] = true
	available[(id+1)%numFilo] = true
	garcom.Broadcast()
	garcom.L.Unlock()
}

func main() {
	for i := 0; i < numFilo; i++ {
		gargo[i] = &sync.Mutex{}
	}

	for i := 0; i < numFilo; i++ {
		wg.Add(1)
		go filosofo(i)
	}

	wg.Wait()
}

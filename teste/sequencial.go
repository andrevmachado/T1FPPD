package main

import (
	"fmt"
	"time"
)

const numFilo = 5

var garfos [numFilo]bool

func main() {
	for i := 0; i < numFilo; i++ {
		garfos[i] = true
	}

	for i := 0; i < numFilo; i++ {
		filosofo(i)
	}
}

func filosofo(id int) {
	for i := 0; i < 10; i++ {
		pensa(id)
		solicitaG(id)
		come(id)
		liberaG(id)
	}
}

func pensa(id int) {
	fmt.Printf("Filósofo %d pensando\n", id)
	time.Sleep(time.Millisecond * 1000)
}

func come(id int) {
	fmt.Printf("Filósofo %d comendo\n", id)
	time.Sleep(time.Millisecond * 1000)
}

func solicitaG(id int) {
	for !garfos[id] || !garfos[(id+1)%numFilo] {
		time.Sleep(time.Millisecond * 10)
	}

	garfos[id] = false
	garfos[(id+1)%numFilo] = false
}

func liberaG(id int) {
	// Libera os garfos e marca como disponíveis
	garfos[id] = true
	garfos[(id+1)%numFilo] = true
}

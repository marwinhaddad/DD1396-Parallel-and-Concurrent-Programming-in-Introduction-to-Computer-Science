// Stefan Nilsson 2013-03-13

// This is a testbed to help you understand channels better.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 4

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgc := new(sync.WaitGroup)
	wgp.Add(producers)
	wgc.Add(consumers)

	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch, wgc)
	}

	wgp.Wait() // Wait for all producers to finish.

	close(ch)
	wgc.Wait() // Wait for all consumers to finish.
	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
	wg.Done()
}

// RandomSleep waits for x ms, where x is a random number, 0 < x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}

// 1.
// När vi byter plats på close och wait signalerar vi att inga fler värden ska skickas på kanalen.
// När vi har wait före close väntar vi in värdena som skickas på kanalen. Vi får panik.

// 2.
// Det bli panik eftersom vi kommer försöka skicka värden in i stängd kanal. Det som kommer att hända är att
// kanalen stängs när första producenten är färdig, då kommer de andra producenterna försöka skicka till en
// stängd kanal.

// 3.
// Vi stänger aldrig kanalen och programmet kommer att köra klart som vanligt.
// Pga att vi hade close i slutet i programmet gör det ingen skillnad att ta bort det i och med att
// programmet kommer att sluta direkt efter.

// 4.
// Ökar vi antalet konsumenter till 4 kommer vi att läsa av från fler producenter samtidigt
// och därmed får vi kortare körtid.

// 5.
// Nej det är lite som i föregående problem att vi inte kan vara säkra på att all hinner printas innan programmet
// körs klart.

// Sakapade en waitGroup, wgc, för konsumenterna som inväntar alla konsumenter innan programmet
// blir färdigt. Se kod för ändringar.

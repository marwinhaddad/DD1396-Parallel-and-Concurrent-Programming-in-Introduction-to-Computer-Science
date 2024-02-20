// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"fmt"
	"sync"
)

// This programs demonstrates how a channel can be used for sending and
// receiving by any number of goroutines. It also shows how  the select
// statement can be used to choose one out of several communications.
func main() {
	people := []string{"Anna", "Bob", "Cody", "Dave", "Eva"}
	match := make(chan string, 1) // Make room for one unmatched send.
	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
	}
	wg.Wait()
	select {
	case name := <-match:
		fmt.Printf("No one received %sâ€™s message.\n", name)
	default:
		// There was no pending send operation.
	}
}

// Seek either sends or receives, whichever possible, a name on the match
// channel and notifies the wait group when done.
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// Wait for someone to receive my message.
	}
	wg.Done()
}

/*
1. 	Hypotes: Vi plockar namn ur listan samtidigt pga go-rutinen. Om vi tar bort go-rutinen kommer vi gå igenom
listan sekventiellt och vi kommer ha samma senders och receivers etc
   	Resultat: Få får samma sender-receiver-par varje körning.

2. 	Hypotes: Vet inte om det är någon skillnad mellan "var wg sync.WaitGroup" och "wg := new(sync.WaitGroup)".
Tror däremot att byta "wg *sync.WaitGroup" mot "wg sync.WaitGroup" kommer att resultera i att programmet kraschar
eller inte kan peka till wg, vilket kommer resultera i att vi aldrig ändrar wg, dvs att alla kommer ge och få något.
 	Resultat: Vi får deadlock pga att Eva inte har någon att skicka till så vi fastnar på
select{... match <- name}, detta pga att wg aldrig blir klar.

3.	Hypotes: Programmet kommer att fastna på wg.Wait() pga att gorutinen kommer att blocka tills att sista sender får
en receiver. Eftersom detta aldrig händer när vi har ett udda antal personer kommer vi att få deadlock.
	Resultat: Vi fick deadlock pga gorutinen blockar. Eftersom vi har en obuffrad kanal blockar sender tills att
receiver har mottagit värdet från kanalen.

4.	Hypotes: Inget kommer att hända eftersom vi har udda antalet personer, så det kommer alltid finnas
någon som inte har en receiver, dvs det kommer alltid finnas någon som återstår i match-kanalen så att vi
hoppar ur select{case name := <- match ...}. Skulle vi ha jämnt antal personer skulle vi få deadlock pga
match-kanalen skulle vara tom då alla senders har en receiver.
	Resultat: Inget händer som förutspått. Lägger vi till personer så att vi har ett jämnt antal får vi deadlock utan
default case.

*/

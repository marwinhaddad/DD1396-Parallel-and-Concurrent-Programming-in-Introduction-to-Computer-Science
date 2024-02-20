// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.

func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	// TODO: Answer questions.
	go answerQuestions(questions, answers)

	// TODO: Make prophecies.
	go makeProphecies(answers)

	// TODO: Print answers.
	go printAnsweres(answers)

	return questions
}

func answerQuestions(questions chan string, answer chan string) {
	for q := range questions {

		listImature := []string{"peepoopeepee.", "...really?", "Childish.", "LOL"}
		listBoring := []string{"12 degrees Celsius.", "4 PM.", "The Eiffel tower."}
		listPersonal := []string{"Weird...", "No.", "Yes.", "Spaghetti with meatballs.", "mmm...Perhaps."}
		listSmart := []string{"Euler-Lagrange.", "Pythagoras.", "The canonical momentum.", "Lax-Millgram, yes?", "The answer you seek is 'Who is Gauss'."}
		listCandice := []string{"Do you know Candice?", "Have you heard of ligma?", "Do you listen to Imagine Dragons?"}
		listElse := []string{"I'm having fun, are you having fun?", "Have you seen Will Smith smack Chris Rock? Crazy.",
			"Does your milkshake bring all the boys to the yard?", "Does the little star twinkle because it wants, or because it needs to?",
			"What äre you thinking?", "Always.", "Don't forget to blow out the candles before you leave home.", ":)",
		}
		listYesNo := []string{"Yes.", "No.", "Maybe."}

		if sliceString(q, []string{"poop", "pee", "fart"}) {
			answer <- listImature[rand.Intn(len(listImature))]
		} else if sliceString(q, []string{"weather?", "time", "where", "temperature?"}) {
			answer <- listBoring[rand.Intn(len(listBoring))]
		} else if sliceString(q, []string{"name?", "old", "color?", "food?"}) {
			answer <- listPersonal[rand.Intn(len(listPersonal))]
		} else if sliceString(q, []string{"Euler", "math", "1", "KTH", "school", "homework"}) {
			answer <- listSmart[rand.Intn(len(listSmart))]
		} else if sliceString(q, []string{"joke?", "joke", "funny"}) {
			answer <- listCandice[rand.Intn(len(listCandice))]
		} else if sliceString(q, []string{"yes", "Yes", "No", "no"}) {
			answer <- listYesNo[rand.Intn(len(listYesNo))]
		} else {
			answer <- listElse[rand.Intn(len(listElse))]
		}
		go prophecy(q, answer)
	}
}

func printAnsweres(answers chan string) {
	for a := range answers {
		time.Sleep(time.Duration(rand.Intn(5-3)+3) * time.Second)
		fmt.Printf("The Great Oracle %s, at %s, says: ", star, venue)
		for _, char := range a {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Print(string(char))
		}
		fmt.Println("")
		fmt.Print(prompt)
	}
}

func makeProphecies(answers chan string) {
	for range time.Tick(time.Duration(rand.Intn(30-20)+20) * time.Second) {
		prophecy("", answers)
	}
}

func sliceString(q string, qlist []string) bool {
	for _, x := range qlist {
		if strings.Contains(q, x) {
			return true
		}
	}
	return false
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(30-20)+20) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"You smell that..?",
		"Have you ever programmed in golang? Yeah, me neither.",
		"Has anyone ever shared a Kitkat?",
		"The flavor of water is its temperature.",
		"It's impossible to hum while holding your nose... go ahead, try it.",
		"Some pointless nonsense.",
		"You look funny.",
		"The right way is never left.",
		"Never microwave banans. Never.",
		"The thing does indeed go skrrah, pap pap, ka-ka-ka.",
		"It is impolite to ask a fox what it says.",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}

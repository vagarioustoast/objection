package main

import (
	"fmt"

	"github.com/vagarioustoast/objection/trial"
)

type Question struct {
	Prompt        string
	Stem          string
	Choices       []string
	CorrectAnswer int
}

func main() {

	exchange := trial.Exchange{
		Question:          "Where were you at approximately 9.00 p.m.?",
		Answer:            "I Was standing outside the grocery store.",
		PossibleObjection: trial.Hearsay,
	}

	fmt.Println(exchange.Question)
	fmt.Println(exchange.Answer)
	fmt.Println("Possible objection:", exchange.PossibleObjection)
	// question := Question{
	// 	Prompt: "After a town installed additional streetlights, reported crime declined. Therefore, the new streetlights caused the decline in crime.",
	// 	Stem:   "Which flaw appears in the argument?",
	// 	Choices: []string{
	// 		"It assumes that reported crime is always accurately measured.",
	// 		"It concludes that one event caused another merely because the events occurred together.",
	// 		"It relies on the testimony of an unqualified expert.",
	// 		"It attacks an opposing position instead of addressing the position itself.",
	// 	},
	// 	CorrectAnswer: 2,
	// }

	// fmt.Println("Welcome to Objection.")
	// fmt.Println()

	// fmt.Println(question.Prompt)
	// fmt.Println()
	// fmt.Println(question.Stem)
	// fmt.Println()

	// for index, choice := range question.Choices {
	// 	fmt.Printf("%d. %s\n", index+1, choice)
	// }

	// var userAnswer int

	// fmt.Println()
	// fmt.Print("Your answer: ")
	// _, err := fmt.Scan(&userAnswer)

	// if err != nil {
	// 	fmt.Println("Please enter a number.")
	// 	return
	// }

	// if userAnswer < 1 || userAnswer > len(question.Choices) {
	// 	fmt.Printf("Please enter a number between 1 and %d\n", len(question.Choices))
	// 	return
	// }

	// if userAnswer == question.CorrectAnswer {
	// 	fmt.Println("Correct!")
	// } else {
	// 	fmt.Printf(
	// 		"Incorrect. The answer was %d.\n",
	// 		question.CorrectAnswer,
	// 	)
	// }
}

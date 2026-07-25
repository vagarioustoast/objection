package main

import "fmt"

type Question struct {
	Prompt        string
	Stem          string
	Choices       []string
	CorrectAnswer int
}

func main() {
	question := Question{
		Prompt: "After a town installed additional streetlights, reported crime declined. Therefore, the new streetlights caused the decline in crime.",
		Stem:   "Which flaw appears in the argument?",
		Choices: []string{
			"It assumes that reported crime is always accurately measured.",
			"It concludes that one event caused another merely because the events occurred together.",
			"It relies on the testimony of an unqualified expert.",
			"It attacks an opposing position instead of addressing the position itself.",
		},
		CorrectAnswer: 2,
	}

	fmt.Println("Welcome to Objection.")
	fmt.Println()

	fmt.Println(question.Prompt)
	fmt.Println()
	fmt.Println(question.Stem)
	fmt.Println()

	for index, choice := range question.Choices {
		fmt.Printf("%d. %s\n", index+1, choice)
	}

	var userAnswer int

	fmt.Println()
	fmt.Print("Your answer: ")
	fmt.Scan(&userAnswer)

	if userAnswer == question.CorrectAnswer {
		fmt.Println("Correct!")
	} else {
		fmt.Printf(
			"Incorrect. The answer was %d.\n",
			question.CorrectAnswer,
		)
	}
}

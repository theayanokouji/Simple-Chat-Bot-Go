package main

import "fmt"

func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

func showName() {
	var name string
	fmt.Println("Please, remind me of your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Printf("Your age is %d; that's a good time to start programming!\n", age)
}

func count() {
	var n int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d !\n", i)
	}
}

func startQuiz() {
	fmt.Println("Let's test your programming knowledge.")
	fmt.Println("Why do we use methods?")
	fmt.Println("1. To repeat a statement multiple times.")
	fmt.Println("2. To decompose a program into several small subroutines.")
	fmt.Println("3. To determine the execution time of a program.")
	fmt.Println("4. To interrupt the execution of a program.")

	// use an infinite for loop until the user enters the correct answer
	for {
		var answer int
		fmt.Scan(&answer)
		if answer == 2 {
			return
		}
		fmt.Println("Please, try again.")
	}
}

func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

func main() {
	greet("VortexBot", "2024")
	showName()
	guessAge()
	count()
	startQuiz()
	sayGoodbye()
}

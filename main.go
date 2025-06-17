package main

import (
	"fmt"
	"os"
)

// Дефинираме типове за състояние и вход
type State int
type Input rune

// Състояния на автомата
const (
	Start State = iota     // Начално състояние
	GotFirstA              // Получено е едно 'a'
	GotSecondA             // Получени са две 'a'
)

// Цветове за отпечатване в терминала
const (
	COLOUR_DEFAULT  = "\033[0m"
	COLOUR_YELLOW   = "\033[1;33m"
	COLOUR_GREEN  = "\033[1;32m"
	COLOUR_RED    = "\033[1;31m"
)

// Връща началното състояние
func getStartState() State { return Start }

// Проверява дали състоянието е крайно
func isFinalState(s State) bool { return s == GotSecondA }

// Преходна функция — определя новото състояние според текущото състояние и входен символ
func move(state State, input Input) State {
	switch state {
		case Start:
			if input == 'a' { return GotFirstA } // от Start към GotFirstA
			if input == 'b' { return Start }     // остава в Start
		case GotFirstA:
			if input == 'a' { return GotSecondA } // от GotFirstA към GotSecondA
			if input == 'b' { return GotFirstA }  // остава в GotFirstA
		case GotSecondA:
			if input == 'a' { return -1 }         // невалиден вход след две 'a' — грешка
			if input == 'b' { return GotSecondA } // остава в GotSecondA
	}
	
	return -1 // невалидно състояние
}

// Разпознава дали входният низ се приема от автомата
func recognize(inputStr string) bool {
	state := getStartState() // започваме от началното състояние

	for _, ch := range inputStr {
		// само 'a' и 'b' са позволени
		if ch != 'a' && ch != 'b' {
			fmt.Printf("Invalid input: %c\n", ch)
			return false
		}
		state = move(state, Input(ch)) // извършваме преход

		if state == -1 { return false } // невалидно състояние — прекратяваме
		
		// Показваме прехода в терминала
		fmt.Printf(
		"%s➤  Processing input: '%c' ➜  Transitioning to state: %d %s\n", 
		COLOUR_YELLOW, 
		ch, 
		state, 
		COLOUR_DEFAULT)
	}

	// приемаме низа само ако сме в крайно състояние
	return isFinalState(state)
}

// Главна функция — стартира програмата
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_string>") // указания за ползване
		return
	}

	input := os.Args[1]
	fmt.Printf("%s■ Processing input string: %s%s\n", COLOUR_YELLOW, input, COLOUR_DEFAULT)
	if recognize(input) {
		fmt.Printf("%s■ String accepted%s\n", COLOUR_GREEN, COLOUR_DEFAULT) // низът е приет
	} else {
		fmt.Printf("%s■ String not accepted%s\n", COLOUR_RED, COLOUR_DEFAULT) // низът е отхвърлен
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/Doniblaze/exercise-repo/internal/section1"
	"github.com/Doniblaze/exercise-repo/internal/section2"
	"github.com/Doniblaze/exercise-repo/internal/section3"
)

var numbers []int

func main() {
	displayMainMenu()
	var userSelection int
	_, err := fmt.Scanln(&userSelection)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	switch userSelection {
	case 1:
		//Variables and Types
		section1.SwapWithoutTemp()
	case 2:
		//Slices
		fmt.Println("Enter a list of numbers separated by spaces:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()

			// Split the input string into individual numbers
			for _, strNum := range strings.Fields(input) {
				num, err := strconv.Atoi(strNum)
				if err != nil {
					fmt.Println("Error converting to integer:", err)
					return
				}
				numbers = append(numbers, num)
			}

			sum := section1.SumAllEvenNumber(numbers)
			fmt.Printf("Sum of Even Numbers: %d\n", sum)
		}
	case 3:
		//Interface
		fmt.Println("Creating a file called logfile.txt......")
		// FileLogger implementation
		fileLogger := section1.FileLogger{FileName: "logfile.txt"}
		fmt.Println("Enter the Message u want to be log into the file and the console:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			message := scanner.Text()

			// Log to the file
			fileLogger.Log(message)

			// ConsoleLogger implementation
			consoleLogger := section1.ConsoleLogger{}
			consoleLogger.Log(message)
		} else {
			fmt.Println("Error reading user input:", scanner.Err())
		}
	case 4:
		//goroutines
		fmt.Println("a Golang program that uses goroutines to calculate the sum of elements in a large array concurrently. Ensuring proper synchronization.")
		section2.RunGorountineExercise()
	case 5:
		//channels
		fmt.Println("a program that uses channels to simulate a simple producer-consumer scenario. The producer generate random numbers and send them to the consumer, which calculates and prints their square.")
		section2.RunChannelExercise()
	case 6:
		//HTTP SERVER
		fmt.Println("a simple Golang HTTP server that listens on port 8080. Responding with 'Hello', 'World'! for any incoming request")
		section3.RunHttpServerExercise()
	case 7:
		//RESTful APIs
		fmt.Println("a basic RESTful API for a todo list. Include operations to create, read, update, and delete tasks. Used in-memory storage for simplicity.")
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
		// Run the RESTful API in a separate Goroutine
		go section3.RunRESTfulApiExercise()

		<-shutdown

		fmt.Println("Exiting the application.")

	}

}

func displayMainMenu() {
	fmt.Println("-------------------")
	fmt.Println("The Assignment")
	fmt.Println("Choose from the options to continue")
	fmt.Println("-------------------")
	fmt.Println("[1] Variables and Types")
	fmt.Println("[2] Slices")
	fmt.Println("[3] Interfaces")
	fmt.Println("[4] Goroutines")
	fmt.Println("[5] Channels")
	fmt.Println("[6] HTTP Server")
	fmt.Println("[7] Rest API")
}

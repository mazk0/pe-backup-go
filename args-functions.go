package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

func printInstructions() {
	fmt.Println("Usage: dotnet run <source:string> <destination:string> <IncludeDateInFileName:bool (optional)>")
	fmt.Println("Example: dotnet run --configuration Release https://pe.makra.dev/api/event/getall backup.json")
}

func formatDestination(destination string) (formattedDestination string) {
	index := strings.LastIndex(destination, ".")
	formattedDestination = fmt.Sprintf("%s_%s%s", destination[:index], time.Now().Format("2006-01-02"), destination[index:])
	return
}

func mapArguments(args []string) (couldMap bool, source string, destination string) {
	if len(os.Args) != 3 {
		fmt.Println("Incorrect number of arguments passed")
		printInstructions()
		return
	}

	source = os.Args[1]
	destination = os.Args[2]

	if _, err := url.ParseRequestURI(source); err != nil {
		fmt.Printf("Incorrectly formatted source Has to be a valid url. Error %s\n", err.Error())
		printInstructions()
		return
	}

	couldMap = true
	return
}

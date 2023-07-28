package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	couldMap, source, destination := mapArguments(os.Args)
	if !couldMap {
		return
	}

	resp, err := http.Get(source)
	if err != nil {
		fmt.Printf("Failed to get event from Api. Error %s\n", err.Error())
	}
	defer resp.Body.Close()

	formattedDestination := formatDestination(destination)

	out, err := os.Create(formattedDestination)
	if err != nil {
		fmt.Printf("Failed to create backup file. Error %s\n", err.Error())
	}
	defer out.Close()

	fmt.Printf("Writing backup to %s\n", formattedDestination)
	bytesWritten, err := io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("Failed to copy Api body to file. Error %s\n", err.Error())
	}

	fmt.Printf("%d was written to %s\n", bytesWritten, formattedDestination)
}

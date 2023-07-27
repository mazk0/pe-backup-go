package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
    "time"
)

func formatDestination(destination string) (formattedDestination string) {
    index := strings.LastIndex(destination, ".")
    now := time.Now()
    formattedDestination = fmt.Sprintf("%s_%s%s", destination[:index], now.Format("2006-01-02"), destination[index:])
    return
}

func mapArguments(args []string) (couldMap bool, source string, destination string) {
    if len(os.Args) != 3 {
        fmt.Println("Incorrect number of arguments passed. Use following command to run application")
        fmt.Println("go run . https://pe.makra.dev/api/event/getall filename.json")
        couldMap = false
        return
    }

    couldMap = true
    source = os.Args[1]
    destination = os.Args[2]
    return
}

func main() {
    couldMap, source, destination := mapArguments(os.Args)
    if !couldMap {
        return
    }

    resp, err := http.Get(source)
    if err != nil {
        fmt.Println(fmt.Sprintf("Failed to get event from Api. Error %s", err.Error()))
    }
    defer resp.Body.Close()

    out, err := os.Create(formatDestination(destination))
    if err != nil {
        fmt.Println(fmt.Sprintf("Failed to create backup file. Error %s", err.Error()))
    }
    defer out.Close()

    bytesWritten, err := io.Copy(out, resp.Body)
    if err != nil {
        fmt.Println(fmt.Sprintf("Failed to copy Api body to file. Error %s", err.Error()))
    }

    fmt.Println(fmt.Sprintf("%d was written to %s", bytesWritten, formattedDestination))
}
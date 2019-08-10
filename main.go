package main

import (
	// "io"
	// "net/http"
    // "os"
    // "strconv"
    "fmt"
    "flag"
    //"io/ioutil"
    "gopl.io/ch4/xkcd/pkg/xkcd"
)

func main() {
    title := flag.String("title", "", "Search for xkcd title")
    year := flag.String("year", "", "Search for xkcd comics from certain year")
    flag.Parse()

    err := xkcd.LoadIndex("./xkcd.info.json")
    if err != nil {
        panic(err)
    }

    if *title != "" {
        result := xkcd.SearchByTitle(*title)
        if result != nil {
            fmt.Printf("\nURL: %s\nTranscript: %s\n", result.URL, result.Transcript)
        }
    }

    if *year != "" {
        result := xkcd.SearchByYear(*year)
        if len(result) > 0 {
            for _, v := range result {
                fmt.Printf("\nURL: %s\nTranscript: %s\n", v.URL, v.Transcript)
            }
        }
    }

    fmt.Println(flag.Args())
    if len(flag.Args()) > 0 {
        result, _ := xkcd.QueryIndex(flag.Args())
        for _, v := range result {
            fmt.Printf("\nURL: %s\nTranscript: %s\n", v.URL, v.Transcript)
        }
    }
}
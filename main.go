package main

import (
    "fmt"
    "net/url"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: url-decode <encoded-url>")
        return
    }

    encodedURL := os.Args[1]

    // Decode the URL
    decodedURL, err := url.QueryUnescape(encodedURL)
    if err != nil {
        fmt.Println("Error decoding URL:", err)
        return
    }

    fmt.Println("Decoded URL:", decodedURL)
}

package decode

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"os"
	"github.com/fatih/color"
)

func DecodeBase64(encodedStr string) string {
    decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", encodedStr)
        os.Exit(1)
    }

    return string(decodedStr)
}

func DecodeBase32(encodedStr string) string {
    decodedStr, err := base32.StdEncoding.DecodeString(encodedStr)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", encodedStr)
        os.Exit(1)
    }

    return string(decodedStr)
}

func IsBase64(encodedStr string) bool {
    _, err := base64.StdEncoding.DecodeString(encodedStr)

    if err != nil {
        return false
    }

    return true
}

func IsBase32(encodedStr string) bool {
    _, err := base64.StdEncoding.DecodeString(encodedStr)

    if err != nil {
        return false
    }

    return true
}

func Echo(encoded, decode string) {
    color.Green("===================")
    fmt.Printf("Input:\t%s\n", encoded)
    fmt.Printf("Result:\t")
    color.Cyan(decode)
    color.Green("===================")
}


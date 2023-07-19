package main

import (
	"bufio"
	"d4r10us/decode/pkg/decode"
	"flag"
	"fmt"
	"os"
)

func getArg() string {
    if len(os.Args) >= 2 {
        encodedStr := os.Args[1]

        return encodedStr
    } 

    r := bufio.NewReader(os.Stdin)
    encodedStr, err := r.ReadString('\n')

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", err)
        os.Exit(1)
    }

    return encodedStr
}

func main() {
    b64 := flag.Bool("b64", true, "Decode base64 string.") 
    b32 := flag.Bool("b32", false, "Decode base32 string.") 
    flag.Parse()

    encodedStr := getArg()

    if *b64 {
        decoded := decode.DecodeBase64(encodedStr)     
        decode.Echo(encodedStr, decoded) 
        os.Exit(0)
    }
    
    if *b32 {
        decoded := decode.DecodeBase32(encodedStr)
        decode.Echo(encodedStr, decoded)
        os.Exit(0)
    }
}


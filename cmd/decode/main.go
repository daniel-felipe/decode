package main

import (
	"bufio"
	"d4r10us/decode/pkg/decode"
	"flag"
	"fmt"
	"os"
)

func main() {
    b64 := flag.String("base64", "", "Decode base64 string.") 
    b32 := flag.String("base32", "", "Decode base32 string.") 
    flag.Parse()

    if *b64 != "" {
        decode.Echo(*b64, decode.DecodeBase64(*b64))
        os.Exit(0)
    } else if *b32 != "" {
        decode.Echo(*b32, decode.DecodeBase32(*b32))
        os.Exit(0)
    }
    
    if *b64 == "" && *b32 == "" {
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        
        encodedStr := scanner.Text()
        
        if (decode.IsBase64(encodedStr)) {
            decode.Echo(encodedStr, decode.DecodeBase64(encodedStr))
            os.Exit(0)
        } else if (decode.IsBase32(encodedStr)) {
            decode.Echo(encodedStr, decode.DecodeBase32(encodedStr))
            os.Exit(0)
        }
    }

    os.Exit(1)
}


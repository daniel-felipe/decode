package decode

import "testing"

func TestDecodeBase64(t *testing.T) {
    encoded := "Sm9obiBEb2U="

    exp := "John Doe"

    res := DecodeBase64(encoded)

    if res != exp {
        t.Errorf("Expected %s, got %s instead.\n", exp, res)
    }
}

func TestDecodeBase32(t *testing.T) {
    encoded := "JJXWQ3RAIRXWK==="

    exp := "John Doe"

    res := DecodeBase32(encoded)

    if res != exp {
        t.Errorf("Expected %s, got %s instead.\n", exp, res)
    }
}

func TestIsBase64(t *testing.T) {
    b64String := "Sm9obiBEb2U="

    exp := true

    res := IsBase64(b64String)

    if res != exp {
        t.Errorf("Expected to be true, got false instead.\n")
    }
}

func TestIsBase32(t *testing.T) {
    b32String := "JJXWQ3RAIRXWK===" 

    exp := true

    res := IsBase32(b32String)

    if res != exp {
        t.Errorf("Expeceted to be true, got false instead.\n")
    }
}


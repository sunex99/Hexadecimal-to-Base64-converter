package main

import (
	"fmt"
	"math/big"
	"strconv"

	"k8s.io/klog/v2"
)

const simpleHex = "eb3d8"

var hexInput string

// useBigIntiger
func useBigIntiger(h string) string {
	// Create a new big.Int
	bigInt := new(big.Int)

	// Set the string value in base 16
	_, ok := bigInt.SetString(h, 16)
	if !ok {
		fmt.Println("Error: could not parse string to big.Int")
		return ""
	}
	// Print the big.Int value
	fmt.Printf("Parsed value is: %s\n", bigInt.String())
	return bigInt.String()
}

func main() {
	/*
		In Go, a string is in effect a read-only slice of bytes. That means that when we store a character value in a string, we store its byte-at-a-time representation.
	*/

	klog.Info("Starting hex to base64 converter. Insert your hex number:")
	//fmt.Scanln(&hexInput)
	hexInput = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	//Lets add prefix
	//hex := fmt.Sprint("0x", "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	// Convert hec into dec
	dec := useBigIntiger(hexInput)

	int, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		// s is not a valid
		klog.Errorf("hex input is not a valid:", err)
	} else {
		klog.Info("Converted to intiger:", int)
	}

}

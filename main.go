package main

import "fmt"

const hexNum = "0x1A3F" // hexNum is a hexadecimal number represented as a string (format: 0x...). "0x" is the prefix for hexadecimal numbers in Go.

// hexToBinary converts a hexadecimal number (as a string) to its binary representation.
func hexToBinary(hex string) string {
	fmt.Printf("Converting hex %s to binary...\n", hex)
	var binary string
	for i := 2; i < len(hex); i++ { // Start from index 2 to skip the "0x" prefix
		switch hex[i] {
		case '0':
			binary += "0000"
		case '1':
			binary += "0001"
		case '2':
			binary += "0010"
		case '3':
			binary += "0011"
		case '4':
			binary += "0100"
		case '5':
			binary += "0101"
		case '6':
			binary += "0110"
		case '7':
			binary += "0111"
		case '8':
			binary += "1000"
		case '9':
			binary += "1001"
		case 'A', 'a':
			binary += "1010"
		case 'B', 'b':
			binary += "1011"
		case 'C', 'c':
			binary += "1100"
		case 'D', 'd':
			binary += "1101"
		case 'E', 'e':
			binary += "1110"
		case 'F', 'f':
			binary += "1111"
		default:
			return "" // Return empty string for invalid characters
		}
	}
	fmt.Printf("Hex %s converted to binary: %s\n", hex, binary)
	return binary
}

// binToDecimal converts a binary string to its decimal representation.
func binToDecimal(binary string) int {
	var decimal int
	for i, bit := range binary {
		if bit == '1' {
			decimal += 1 << (len(binary) - 1 - i) // Calculate the value of the bit and add it to the total
		}
	}
	return decimal
}

func binaryToBase64(binary string) string {
	// binaryToBase64 converts a binary string to its Base64 representation.
	const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	var base64 string

	// add padding to make the length of binary a multiple of 6
	for len(binary)%6 != 0 {
		binary = "0" + binary // Prepend zeros to the binary string until its length is a multiple of 6
	}

	for i := 0; i < len(binary); i += 6 {
		// Take 6 bits at a time
		bitSegment := binary[i : i+6]
		fmt.Printf("Processing binary segment: %s, with length %d\n", bitSegment, len(bitSegment))

		// Convert the 6-bit segment to a decimal number
		decimalValue := binToDecimal(bitSegment)
		fmt.Printf("Decimal value of segment %s: %d\n", bitSegment, decimalValue)

		// Map the decimal value to a Base64 character
		if decimalValue < 0 || decimalValue >= len(base64Chars) {
			fmt.Printf("Error: Decimal value %d is out of range for Base64 characters.\n", decimalValue)
			return ""
		}
		base64 += string(base64Chars[decimalValue])
	}
	fmt.Printf("Converting binary %s to Base64...\n", binary)
	return base64
}

func main() {
	fmt.Printf("Starting the program...\n")
	b := hexToBinary(hexNum)
	b64 := binaryToBase64(b)
	if b64 == "" {
		fmt.Printf("Conversion failed due to an error.\n")
		return
	}
	fmt.Printf("Hexadecimal %s in Base64 is: %s\n", hexNum, b64)
	fmt.Printf("Program finished.\n")
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create a new []byte b32
	b32 := []byte("float32:")

	// AppendFloat appends the string form of the floating-point number f,
	// as generated by FormatFloat, to dst and returns the extended buffer.
	// func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)

	// Print b32
	fmt.Println(string(b32))

	// Create a new []byte b64
	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)

	// Print b64
	fmt.Println(string(b64))
}
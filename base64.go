package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"strconv"
)

const b64chars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func prepareString(str string) string {

	// Convert to binary
	var outstr string
	for _, c := range str {
		outstr = fmt.Sprintf("%s%.8b", outstr, c)
	}

	// Split to 6-bit and convert to decimal - then b64
	var decstr, b64string string
	var i int
	for i = 0; i+6 < len(outstr); i = i+6 {
		decstr = outstr[i:i+6]
		decnum, _ := strconv.ParseInt(decstr, 2, 0)
		b64string += string(b64chars[decnum])
	}
	decstr = outstr[i:]
	for math.Mod(float64(len(decstr)), 6) != 0 {
		decstr += "0"
	}
	decnum, _ := strconv.ParseInt(decstr, 2, 0)
	b64string += string(b64chars[decnum])

	// Output padding
	for math.Mod(float64(len(b64string)), 4) != 0 {
		b64string += "="
	}

	return b64string
}

func main() {
	var instr string
	for {
		fmt.Print("Enter text: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		instr = scanner.Text()
		fmt.Println(prepareString(instr))
	}
}

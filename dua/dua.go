package main

import (
	"fmt"
	"strings"
)

func cekInputString(input1 string, input2 string) bool {
	res := 0
	if input1 == input2 {
		return true
	}
	for i := 0; i < len(input1); i++ {
		if input1[:i] != input2[:i] {
			if i < (len(input2))-1 && i < (len(input1))-1 {

				if input1[i] == input2[i+1] { // cek apakah delete
					res++
					// fmt.Println("delet", string(input1), string(input2))
					input1 = strings.Replace(input1, string(input1[i]), string(input1[i+1]), -1)

				} else if input1[i+1] == input2[i] {
					input1 = strings.Replace(input1, string(input1[:i+1]), string(input2[:i]), -1)
					res++
				} else if input1[i+1] == input2[i+1] { // cek apakah eidt
					input1 = strings.Replace(input1, string(input1[:i]), string(input2[:i]), -1)
					// fmt.Println("edit", string(input1), string(input2))
					res++
				}
				if input1[i] != input2[i] {
					// fmt.Println("else", string(input1), string(input2))
					res++
				}
			}
		}

	}
	// fmt.Println(res)
	if res > 1 {
		return false
	} else {
		return true

	}
}

func main() {
	fmt.Println(cekInputString("telkom", "telecom")) //expect false
	fmt.Println(cekInputString("telkom", "tlkom"))   // expect true
	fmt.Println(cekInputString("telkom", "tilkom"))  // expect true
	fmt.Println(cekInputString("telkom", "telkom"))  // expect true
}

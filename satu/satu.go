package main

import (
	"encoding/json"
	"fmt"
)

func getbanyakmatauang(input int, nominal int) int {
	inputintVar := input
	res := 0
	if inputintVar-nominal >= 0 {
		res++
		return res + getbanyakmatauang(input-nominal, nominal)
	} else {

		return 0
	}

	return res
}

func resJSON(input int) string {
	// input := 260000
	var nominal = map[string]int{

		"Rp. 100.000": 0,
		"Rp. 50.000":  0,
		"Rp. 20.000":  0,
		"Rp. 10.000":  0,
		"Rp. 5.000":   0,
		"Rp. 2.000":   0,
		"Rp. 1.000":   0,
		"Rp. 500":     0,
		"Rp. 200":     0,
		"Rp. 100":     0,
	}

	nominal["Rp. 100.000"] = getbanyakmatauang(input, 100000)
	input -= (nominal["Rp. 100.000"] * 100000)
	nominal["Rp. 50.000"] = getbanyakmatauang(input, 50000)
	input -= (nominal["Rp. 50.000"] * 50000)
	nominal["Rp. 20.000"] = getbanyakmatauang(input, 20000)
	input -= (nominal["Rp. 20.000"] * 20000)
	nominal["Rp. 10.000"] = getbanyakmatauang(input, 10000)
	input -= (nominal["Rp. 10.000"] * 10000)
	nominal["Rp. 5.000"] = getbanyakmatauang(input, 5000)
	input -= (nominal["Rp. 5.000"] * 5000)
	nominal["Rp. 2.000"] = getbanyakmatauang(input, 2000)
	input -= (nominal["Rp. 2.000"] * 2000)
	nominal["Rp. 1.000"] = getbanyakmatauang(input, 1000)
	input -= (nominal["Rp. 1.000"] * 1000)
	nominal["Rp. 500"] = getbanyakmatauang(input, 500)
	input -= (nominal["Rp. 500"] * 500)
	nominal["Rp. 200"] = getbanyakmatauang(input, 200)
	input -= (nominal["Rp. 200"] * 200)
	nominal["Rp. 100"] = getbanyakmatauang(input, 100)
	input -= (nominal["Rp. 100"] * 100)

	// bs, _ := json.Marshal(nominal)
	res, _ := json.MarshalIndent(nominal, "", " ")
	return string(res)
}

func main() {
	// fmt.Println("Hello, World")
	fmt.Println(resJSON(100000))
}

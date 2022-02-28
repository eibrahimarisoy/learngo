// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare empty arrays
//
//  1. Declare and print the following arrays with their types:
//
//    1. The names of your three best friends (names array)
//
//    2. The distances to five different locations (distances array)
//
//    3. A data buffer with five bytes of capacity (data array)
//
//    4. Currency exchange ratios only for a single currency (ratios array)
//
//    5. Up/Down status of four different web servers (alives array)
//
//    6. A byte array that doesn't occupy memory space (zero array)
//
//  2. Print only the types of the same arrays.
//
//  3. Print only the elements of the same arrays.
//
// HINT
//   When printing the elements of an array, you can use the usual Printf verbs.
//
//   For example:
//     When printing a string array, you can use "%q" verb as usual.
//
// ---------------------------------------------------------

func main() {
	// names := [3]string{}
	// fmt.Printf("names: %q\n", names)

	// distances := [5]float64{}
	// fmt.Printf("distances: %q\n", distances)

	// data := [5]byte{}
	// fmt.Printf("data: %q\n", data)

	// ratios := [1]float64{}
	// fmt.Printf("ratios: %q\n", ratios)

	// alives := [4]bool{}
	// fmt.Printf("alives: %q\n", alives)

	// zero := [0]byte{}
	// fmt.Printf("zero: %q\n", zero)

	var (
		names     [3]string
		distances [5]float64
		data      [5]byte
		ratios    [1]float64
		alives    [4]bool
		zero      [0]byte
	)

	fmt.Printf("%#v \n", names)
	fmt.Printf("%#v \n", distances)
	fmt.Printf("%#v \n", data)
	fmt.Printf("%#v \n", ratios)
	fmt.Printf("%#v \n", alives)
	fmt.Printf("%#v \n", zero)

}

// EXPECTED OUTPUT
//  names    : [3]string{"", "", ""}
//  distances: [5]int{0, 0, 0, 0, 0}
//  data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
//  ratios   : [1]float64{0}
//  alives   : [4]bool{false, false, false, false}
//  zero     : [0]uint8{}
//
//  names    : [3]string
//  distances: [5]int
//  data     : [5]uint8
//  ratios   : [1]float64
//  alives   : [4]bool
//  zero     : [0]uint8
//
//  names    : ["" "" ""]
//  distances: [0 0 0 0 0]
//  data     : [0 0 0 0 0]
//  ratios   : [0.00]
//  alives   : [false false false false]
//  zero     : []

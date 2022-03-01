// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Right Trim It
//
//  1. Look at the documentation of strings package
//  2. Find a function that trims the spaces from
//     only the right-most part of the given string
//  3. Trim it from the right part only
//  4. Print the number of characters it contains.
//
// RESTRICTION
//  Your program should work with unicode string values.
//
// EXPECTED OUTPUT
//  5
// ---------------------------------------------------------

// func main() {
// 	// currently it prints 17
// 	// it should print 5

// 	name := "7u8u0o0p111p"
// 	count := 0
// 	tmp := ""
// 	for _, i := range name {
// 		if _, err := strconv.Atoi(string(i)); err == nil {
// 			tmp += string(i)
// 		} else {
// 			fmt.Println(tmp)
// 			new, _ := strconv.Atoi(tmp)
// 			count += new
// 			tmp = ""
// 		}
// 	}
// 	new, _ := strconv.Atoi(tmp)
// 	count += new
// 	tmp = ""

// 	fmt.Println(count)
// }

func main() {
	a := "3 04 PM"
	b := "00 00 PM"
	fmt.Println(time.Parse(a, b))

}



for i := 0; i < len(arr); i++{
    for j := i+1 ; j < len(arr); j++{
      if arr[i] == arr[j]{
        maxRepeated[strconv.Itoa(arr[i])]++
      }
    }

  }
    fmt.Println(maxRepeated)


  return arr[0];











  for key, value := range maxRepeated{
    fmt.Println(key, value)
    if value != 1{
      break
    }
  }
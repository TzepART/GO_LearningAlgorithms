package main

import "fmt"

func containsDuplicates(checkingArray []string) bool {

	result := false

	for y := range checkingArray {
		for x := range checkingArray {
			if y != x && checkingArray[y] == checkingArray[x]{
				result = true
			}
		}
	}

	return result
}

func main() {
	checkingArray := []string{"one", "two", "one"}
	if containsDuplicates(checkingArray){
		fmt.Printf("Duplicates exist\n")
	}else{
		fmt.Printf("Duplicates don't exist\n")
	}
}
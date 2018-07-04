package main

import "fmt"

func containsDuplicatesFirstVersion(checkingArray []string) bool {

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

func containsDuplicatesSecondVersion(checkingArray []string) bool {

	result := false
	lengthArray := len(checkingArray)

	for y:=0; y < lengthArray-1; y++ {
		for x:=y+1; x < lengthArray; x++ {
			if checkingArray[y] == checkingArray[x]{
				result = true
			}
		}
	}

	return result
}

func main() {
	checkingArray := []string{"one", "two", "one", "two"}
	if containsDuplicatesFirstVersion(checkingArray){
		fmt.Printf("V1 - Duplicates exist\n")
	}else{
		fmt.Printf("V1 - Duplicates don't exist\n")
	}

	if containsDuplicatesSecondVersion(checkingArray){
		fmt.Printf("V2 - Duplicates exist\n")
	}else{
		fmt.Printf("V2 - Duplicates don't exist\n")
	}
}
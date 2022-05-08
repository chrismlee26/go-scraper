package main

import (
	"encoding/json"
	"fmt"
)

type FruitList struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// fruitSlice := []string{"apple", "peach", "pear"}
	fruitList := &FruitList{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	fruitJSON, _ := json.Marshal(fruitList)
	fmt.Println(string(fruitJSON))

	totalFruitsMap := map[string]int{"apple": 5, "lettuce": 7}
	totalFruitsJSON, _ := json.Marshal(totalFruitsMap)
	fmt.Println(string(totalFruitsJSON))

}

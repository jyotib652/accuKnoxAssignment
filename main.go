package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Data struct {
	EaterId    int `json:"eater_id"`
	FoodMenuId int `json:"foodmenu_id"`
}

// var FoodItems map[int]int
// var FoodItemsCount map[int]int

// var FoodItems = map[int]int{}
// var FoodItemsCount = map[int]int{}

// var TopFoodItems = []int{}

var file = "./log.json"

func main() {
	file := "./log.json"
	data, err := readLogFileData(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("data:", data)

	dataFromError, msg := detectError(data)
	fmt.Printf("dataFromError: %v\n", dataFromError)
	fmt.Printf("error: %s\n", msg)

	TopItems := determineTopFoodItems(dataFromError)
	fmt.Println("TopItems:", TopItems)
}

// detectError detects if the same diner had same food twice
func detectError(data []Data) (map[int]int, error) {
	FoodItemsCount := map[int]int{}
	FoodItems := map[int]int{}
	for _, val := range data {
		FoodItemsCount[val.FoodMenuId] += 1
		_, ok := FoodItems[val.EaterId]
		if !ok {
			FoodItems[val.EaterId] = val.FoodMenuId
		} else {
			if FoodItems[val.EaterId] == val.FoodMenuId {
				return FoodItemsCount, fmt.Errorf("the same diner(id:%d) had same food item(id:%d) more than once", val.EaterId, val.FoodMenuId)
			}

		}
	}
	return FoodItemsCount, nil
}

// determineTopFoodItems finds out the top three food items
func determineTopFoodItems(data map[int]int) []int {
	TopFoodItems := []int{}
	for _, val := range data {
		TopFoodItems = append(TopFoodItems, val)
	}
	// sort.Ints(TopFoodItems)                              //ascending order
	sort.Sort(sort.Reverse(sort.IntSlice(TopFoodItems))) //descending order
	fmt.Println("TopFoodItems:", TopFoodItems)
	var first, second, third int
	for key, val := range data {
		if val == TopFoodItems[0] {
			first = key
		} else if val == TopFoodItems[1] {
			second = key
		} else if val == TopFoodItems[2] {
			third = key
		}
	}
	topThreeFoodItems := []int{first, second, third}

	return topThreeFoodItems
}

// readLogFileData reads log data from a json file and parse it
func readLogFileData(s string) ([]Data, error) {
	jsonFile, err := os.Open(s)
	if err != nil {
		return nil, err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var payload []Data

	err = json.Unmarshal(byteValue, &payload)
	if err != nil {
		return nil, err
	}

	return payload, err
}

package main

import (
	"errors"
	"fmt"
	"testing"
)

var data []Data
var testFoodItemsCount map[int]int

func Test_readLogFileData(t *testing.T) {
	data, _ := readLogFileData(file)
	for _, val := range data {
		if val.EaterId == 0 {
			t.Errorf("expected eater_id other than zero but got: %d\n", val.EaterId)
		}
		if val.FoodMenuId == 0 {
			t.Errorf("expected foodmenu_id other than zero but got: %d\n", val.FoodMenuId)
		}
	}
}

func Test_detectError(t *testing.T) {
	expected := errors.New("the same diner(id:1) had same food item(id:1) more than once")
	_, msg := detectError(data)

	if errors.Is(msg, expected) {
		t.Errorf("expected:%v but got:%v", expected, msg)
	}
}

func Test_determineTopFoodItems(t *testing.T) {
	topItems := determineTopFoodItems(testFoodItemsCount)
	fmt.Println("topItems:", topItems)
	if len(topItems) != 3 {
		t.Errorf("expected:a slice of length 3, but got:%v", topItems)
	}
}

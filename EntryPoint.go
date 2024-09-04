package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int32
}

type Tier string

const (
	Common = iota
	Uncommon
	Rare
	SuperRare
	Epic
	Prestige
	Exotic
	Pinnacle
)

type CommonData struct {
	Name            string
	Description     string
	Tier            Tier
	Id              string
	IsPrestigeProof bool
	Rarity          float32
	IsShopItem      bool
	ItemValue       float64
}

func main() {

	var person Person
	jsonData := `{"name":"Nathan","age":30}`
	json.Unmarshal([]byte(jsonData), &person)
	fmt.Printf("Name: %s, Age: %d", person.Name, person.Age)
}

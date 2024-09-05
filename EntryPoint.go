package main

import (
	"OreForgeWebserver/UserContent"
	"encoding/json"
	"fmt"
)

func main() {

	commonData := UserContent.CommonItemData{
		Name:            "Test Item",
		Description:     "Test Description",
		Tier:            UserContent.EPIC,
		Id:              "testing123",
		IsPrestigeProof: true,
		Rarity:          1,
		IsShopItem:      false,
		ItemValue:       0,
	}

	jsonData, err := json.Marshal(commonData)
	if err != nil {
		return
	}

	var storedData UserContent.CommonItemData
	fmt.Printf("JSON FORM: %s\n", string(jsonData))
	json.Unmarshal(jsonData, &storedData)
	fmt.Printf("Name: %s, Tier: %s\n", storedData.Name, storedData.Tier)
}

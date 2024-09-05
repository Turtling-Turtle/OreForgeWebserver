package UserContent

type Tier string

const (
	COMMON     = "COMMON"
	UNCOMMON   = "UNCOMMON"
	RARE       = "RARE"
	SUPER_RARE = "SUPER_RARE"
	EPIC       = "EPIC"
	PRESTIGE   = "PRESTIGE"
	EXOTIC     = "EXOTIC"
	PINNACLE   = "PINNACLE"
)

type Currency string

type UnlockMethod string

const (
	SPECIAL_POINTS = "SPECIAL_POINTS"
	PRESTIGE_LEVEL = "PRESTIGE_LEVEL"
	QUEST          = "QUEST"
	NONE           = "NONE"
)

type CommonItemData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          string `json:"id"`

	Tier Tier `json:"tier"`

	IsPrestigeProof bool `json:"isPrestigeProof"`

	IsShopItem        bool         `json:"isShopItem"`
	ItemValue         float64      `json:"itemValue"`
	UnlockMethod      UnlockMethod `json:"unlockMethod"`
	UnlockRequirement float64      `json:"unlockRequirement"`
	SellPrice         float64      `json:"sellPrice"`
	Rarity            float32      `json:"rarity"`
}

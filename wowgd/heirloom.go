package wowgd

type HeirloomIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Heirlooms []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"heirlooms"`
}

type Heilroom struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID   int `json:"id"`
	Item struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"item"`
	Source struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"source"`
	SourceDescription string `json:"source_description"`
	Upgrades          []struct {
		Item struct {
			Item struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID int `json:"id"`
			} `json:"item"`
			Context   int   `json:"context"`
			BonusList []int `json:"bonus_list"`
			Quality   struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"quality"`
			Name  string `json:"name"`
			Media struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID int `json:"id"`
			} `json:"media"`
			ItemClass struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"item_class"`
			ItemSubclass struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"item_subclass"`
			InventoryType struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"inventory_type"`
			Binding struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"binding"`
			Weapon struct {
				Damage struct {
					MinValue      int    `json:"min_value"`
					MaxValue      int    `json:"max_value"`
					DisplayString string `json:"display_string"`
					DamageClass   struct {
						Type string `json:"type"`
						Name string `json:"name"`
					} `json:"damage_class"`
				} `json:"damage"`
				AttackSpeed struct {
					Value         int    `json:"value"`
					DisplayString string `json:"display_string"`
				} `json:"attack_speed"`
				Dps struct {
					Value         float64 `json:"value"`
					DisplayString string  `json:"display_string"`
				} `json:"dps"`
			} `json:"weapon"`
			Stats []struct {
				Type struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"type"`
				Value   int `json:"value"`
				Display struct {
					DisplayString string `json:"display_string"`
					Color         struct {
						R int `json:"r"`
						G int `json:"g"`
						B int `json:"b"`
						A int `json:"a"`
					} `json:"color"`
				} `json:"display"`
				IsEquipBonus bool `json:"is_equip_bonus,omitempty"`
			} `json:"stats"`
			Upgrades struct {
				Value         int    `json:"value"`
				MaxValue      int    `json:"max_value"`
				DisplayString string `json:"display_string"`
			} `json:"upgrades"`
			Requirements struct {
				Level struct {
					DisplayString string `json:"display_string"`
				} `json:"level"`
			} `json:"requirements"`
			Level struct {
				Value         int    `json:"value"`
				DisplayString string `json:"display_string"`
			} `json:"level"`
		} `json:"item"`
		Level int `json:"level"`
	} `json:"upgrades"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
}

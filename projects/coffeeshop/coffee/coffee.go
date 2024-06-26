package coffee

import (
	"fmt"

	"github.com/spf13/viper"
)

//var coffees = map[string]float32{"Latte":2.5, "Cappuccino": 2.75, "Flat White": 2.25}

type CoffeeDetails struct {
	Name  string
	Price float32
}

type CoffeeList struct {
	List []CoffeeDetails
}

var Coffees CoffeeList

func GetCoffees() (*CoffeeList, error) {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: %w", err)
		return nil, err
	}

	err = viper.Unmarshal(&Coffees)
	if err != nil {
		return nil, err
	}

	return &Coffees, nil
}

func IsCoffeeAvailable(coffeetype string) bool {
	// search for a specific type of coffee given the input type
	for _, coffee := range Coffees.List {
		if coffee.Name == coffeetype {
			fmt.Println(coffee.Name, "for", coffee.Price)
			return true
		}
	}
	return false
}

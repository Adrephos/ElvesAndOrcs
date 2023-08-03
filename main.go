package main

import (
	"fmt"

	factorymaker "github.com/Adrephos/ElvesAndOrcs/FactoryMaker"
	kingdom "github.com/Adrephos/ElvesAndOrcs/Kingdom"
)

func printKingdom(k *kingdom.Kingdom) {
	fmt.Printf("King: %v \nArmy: %v \nCastle: %v \n",
		k.GetKing().GetDescription(), k.GetArmy().GetDescription(), k.GetCastle().GetDescription())
}

func main() {
	orcKingdom := kingdom.NewKingdom(factorymaker.ORC)
	elfKingdom := kingdom.NewKingdom(factorymaker.ELF)

	printKingdom(orcKingdom)
	fmt.Println()
	printKingdom(elfKingdom)
}

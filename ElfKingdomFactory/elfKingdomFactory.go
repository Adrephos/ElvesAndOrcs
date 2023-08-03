package elfkingdomfactory

import (
	elfarmy "github.com/Adrephos/ElvesAndOrcs/ElfArmy"
	elfcastle "github.com/Adrephos/ElvesAndOrcs/ElfCastle"
	elfking "github.com/Adrephos/ElvesAndOrcs/ElfKing"
	interfaces "github.com/Adrephos/ElvesAndOrcs/Interfaces"
)

type ElfKingdomFactory struct {
}

func (ekf *ElfKingdomFactory) CreateArmy()  interfaces.Army {
	return elfarmy.NewElfArmy()
}

func (ekf *ElfKingdomFactory) CreateCastle()  interfaces.Castle {
	return elfcastle.NewElfCastle()
}

func (ekf *ElfKingdomFactory) CreateKing()  interfaces.King {
	return elfking.NewElfKing()
}

func NewElfKingdomFactory() *ElfKingdomFactory {
	return &ElfKingdomFactory{}
}

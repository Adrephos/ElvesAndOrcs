package orckingdomfactory

import (
	orcarmy "github.com/Adrephos/ElvesAndOrcs/OrcArmy"
	orccastle "github.com/Adrephos/ElvesAndOrcs/OrcCastle"
	orcking "github.com/Adrephos/ElvesAndOrcs/OrcKing"
	interfaces "github.com/Adrephos/ElvesAndOrcs/Interfaces"
)

type OrcKingdomFactory struct {
}

func (okf OrcKingdomFactory) CreateArmy()  interfaces.Army {
	return orcarmy.NewOrcArmy()
}

func (okf OrcKingdomFactory) CreateCastle()  interfaces.Castle {
	return orccastle.NewOrcCastle()
}

func (okf OrcKingdomFactory) CreateKing()  interfaces.King {
	return orcking.NewOrcKing()
}

func NewOrcKingdomFactory() *OrcKingdomFactory {
	return &OrcKingdomFactory{}
}

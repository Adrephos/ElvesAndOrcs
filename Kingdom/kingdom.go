package kingdom

import (
	factorymaker "github.com/Adrephos/ElvesAndOrcs/FactoryMaker"
	interfaces "github.com/Adrephos/ElvesAndOrcs/Interfaces"
)

type Kingdom struct {
	army   interfaces.Army
	castle interfaces.Castle
	king   interfaces.King
}

func (k Kingdom) GetArmy() interfaces.Army { return k.army }

func (k Kingdom) GetCastle() interfaces.Castle { return k.castle }

func (k Kingdom) GetKing() interfaces.King { return k.king }

func (k *Kingdom) SetArmy(army interfaces.Army) {
	k.army = army
}

func (k *Kingdom) SetCastle(castle interfaces.Castle) {
	k.castle = castle
}

func (k *Kingdom) SetKing(king interfaces.King) {
	k.king = king
}

func NewKingdom(kType factorymaker.KingdomType) *Kingdom {
	factoryM := factorymaker.NewFactoryMaker()
	factory := factoryM.MakeFactory(kType)

	return &Kingdom{
		army: factory.CreateArmy(),
		castle: factory.CreateCastle(),
		king: factory.CreateKing(),
	}
}

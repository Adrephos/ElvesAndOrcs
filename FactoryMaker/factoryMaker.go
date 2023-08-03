package factorymaker

import (
	"log"

	elfkingdomfactory "github.com/Adrephos/ElvesAndOrcs/ElfKingdomFactory"
	interfaces "github.com/Adrephos/ElvesAndOrcs/Interfaces"
	orckingdomfactory "github.com/Adrephos/ElvesAndOrcs/OrcKingdomFactory"
)

type KingdomType string

const (
	ELF = "ELF"
	ORC = "ORC"
)

type FactoryMaker struct {
}

func (fm FactoryMaker) MakeFactory(kType KingdomType) interfaces.KingdomFactory {
	switch kType {
	case ELF:
		return elfkingdomfactory.NewElfKingdomFactory()
	case ORC:
		return orckingdomfactory.NewOrcKingdomFactory()
	default:
		log.Printf("%s is not a known kingdom type", kType)
		return nil
	}
}

func NewFactoryMaker() *FactoryMaker {
	return &FactoryMaker{}
}

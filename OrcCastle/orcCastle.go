package orccastle

const DESCRIPTION =	"An ugly orc castle"

type OrcCastle struct {
	description string
}

func (oc *OrcCastle) GetDescription() string { return oc.description }

func NewOrcCastle() *OrcCastle {
	return &OrcCastle{description: DESCRIPTION}
}

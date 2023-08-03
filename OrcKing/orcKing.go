package orcking

const DESCRIPTION =	"King Korc"

type OrcKing struct {
	description string
}

func (ok *OrcKing) GetDescription() string { return ok.description }

func NewOrcKing() *OrcKing {
	return &OrcKing{description: DESCRIPTION}
}

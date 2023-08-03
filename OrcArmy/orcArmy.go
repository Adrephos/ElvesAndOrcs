package orcarmy

const DESCRIPTION = "An orc army"

type OrcArmy struct {
	description string
}

func (oa OrcArmy) GetDescription() string { return oa.description }

func NewOrcArmy() *OrcArmy {
	return &OrcArmy{description: DESCRIPTION}
}

package elfarmy

const DESCRIPTION = "An elf army xd xd"

type ElfArmy struct {
	description string
}

func (ea ElfArmy) GetDescription() string { return ea.description }

func NewElfArmy() *ElfArmy {
	return &ElfArmy{description: DESCRIPTION}
}

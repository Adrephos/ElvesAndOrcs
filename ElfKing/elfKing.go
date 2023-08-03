package elfking

const DESCRIPTION =	"Legolas xd"

type ElfKing struct {
	description string
}

func (ek *ElfKing) GetDescription() string { return ek.description }

func NewElfKing() *ElfKing {
	return &ElfKing{description: DESCRIPTION}
}

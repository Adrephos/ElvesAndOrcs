package elfcastle

const DESCRIPTION =	"A pretty big Elf castle"

type ElfCastle struct {
	description string
}

func (ec *ElfCastle) GetDescription() string { return ec.description }

func NewElfCastle() *ElfCastle {
	return &ElfCastle{description: DESCRIPTION}
}

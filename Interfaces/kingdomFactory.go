package interfaces

type KingdomFactory interface {
	CreateArmy() Army
	CreateCastle() Castle
	CreateKing() King
}

package Postgres

type MigrationDirection int

const (
	Up   MigrationDirection = 0
	Down MigrationDirection = 1
)

func (o *MigrationDirection) String() string {
	switch *o {
	case Up:
		return "up"
	case Down:
		return "down"
	default:
		panic("invalid Migration direction")
	}
}

func (o *MigrationDirection) Valid() bool {
	switch *o {
	case Up, Down:
		return true
	default:
		return false
	}
}

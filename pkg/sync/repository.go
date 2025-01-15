package sync

type Repository interface {
	SetSrc(id string, src string) error
	GetSrc(id string) (string, error)
	SetSec(id string, sec int) error
	GetSec(id string) (int, error)
	SetCmd(id string, cmd string) error
	GetCmd(id string) (string, error)
}

package antibody

type Bundle interface {
	Download() (string, error)
	Update() (string, error)
}

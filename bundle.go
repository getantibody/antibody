package antibody

type Bundle interface {
	Folder() string
	Download() error
	Update() error
	Sourceables() []string
}

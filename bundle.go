package antibody

// Bundle represents a shell plugin.
type Bundle interface {
	Folder() string
	Download() error
	Update() error
	Sourceables() []string
}

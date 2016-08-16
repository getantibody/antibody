package project

// NewLocal Returns a local project, which can be any folder you want to
func NewLocal(folder string) Project {
	return localProject{folder}
}

type localProject struct {
	folder string
}

func (l localProject) Download() error {
	return nil
}

func (l localProject) Update() error {
	return nil
}

func (l localProject) Folder() string {
	return l.folder
}

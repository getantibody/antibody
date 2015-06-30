package main

type Bundle interface {
	Folder() string
	Download() error
	Update() error
	Sourceables() []string
}

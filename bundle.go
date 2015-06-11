package main

type Bundle interface {
	Download() (string, error)
	Update() (string, error)
}

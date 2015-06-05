package antibody

func main() {
	home := Home()
	if ReadStdin() {
		ProcessStdin(home)
	} else {
		ProcessArgs(home)
	}
}

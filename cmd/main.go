package main

import (
	"fmt"

	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody"
)

func main() {
	sh, err := antibody.New(
		gohome.Cache("__antibodyss")+"/",
		[]string{
			"caarlos0/ports kind:path",
			"caarlos0/jvm branch:gh-pages",
			"caarlos0/zsh-open-pr kind:zsh",
		},
	).Bundle()
	if err != nil {
		panic(err)
	}
	fmt.Println(sh)
}

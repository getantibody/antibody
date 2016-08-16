package main

import (
	"fmt"

	"github.com/caarlos0/gohome"
	"github.com/getantibody/antibody"
)

func main() {
	sh, err := antibody.New(
		gohome.Cache("__antibody")+"/",
		[]string{
			"caarlos0/ports kind:path",
			"caarlos0/jvm version:gh-pages",
		},
	).Bundle()
	if err != nil {
		panic(err)
	}
	fmt.Println(sh)
}

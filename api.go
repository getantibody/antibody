package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func Bundle(bundle string, home string) {
	folder, err := Clone(bundle, home)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(folder)
}

func process(bundle string, home string, wg *sync.WaitGroup) {
	defer wg.Done()
	Bundle(bundle, home)
}

func ProcessStdin(stdin io.Reader, home string) {
	var wg sync.WaitGroup
	bundles, _ := ioutil.ReadAll(stdin)
	for _, bundle := range strings.Split(string(bundles), "\n") {
		if bundle != "" {
			wg.Add(1)
			go process(bundle, home, &wg)
		}
	}
	wg.Wait()
}

func Update(home string) ([]string, error) {
	bundles, _ := ioutil.ReadDir(home)
	var sourceables []string
	for _, bundle := range bundles {
		if bundle.Mode().IsDir() && bundle.Name()[0] != '.' {
			updated, err := Pull(bundle.Name(), home)
			if err != nil {
				return sourceables, err
			}
			sourceables = append(sourceables, updated)
		}
	}
	return sourceables, nil
}

func ProcessArgs(args []string, home string) {
	cmd := args[0]
	if cmd == "update" {
		Update(home)
	} else if cmd == "bundle" {
		Bundle(args[1], home)
	} else {
		panic("Invalid command: " + cmd)
	}
}

func ReadStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func Home() string {
	home := os.Getenv("ANTIBODY_HOME")
	if home == "" {
		home = os.Getenv("HOME") + "/.antibody/"
	} else {
		if !strings.HasSuffix(home, "/") {
			home += "/"
		}
	}
	return home
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/alecthomas/kingpin"
	"github.com/getantibody/antibody/antibodylib"
	"github.com/getantibody/antibody/project"
	"github.com/getantibody/antibody/shell"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	version = "dev"

	app         = kingpin.New("antibody", "The fastest shell plugin manager")
	parallelism = app.Flag("parallelism", "max amount of tasks to launch in parallel").
			Short('p').
			Default(strconv.Itoa(runtime.NumCPU())).
			Int()
	bundleCmd = app.Command("bundle", "downloads a bundle and prints its source line")
	bundles   = bundleCmd.Arg("bundles", "bundle list").Strings()
	updateCmd = app.Command("update", "updates all previously bundled bundles")
	homeCmd   = app.Command("home", "prints where antibody is cloning the bundles")
	purgeCmd  = app.Command("purge", "purges a bundle from your computer")
	purgee    = purgeCmd.Arg("bundle", "bundle to be purged").String()
	listCmd   = app.Command("list", "lists all currently installed bundles").Alias("ls")
	initCmd   = app.Command("init", "initializes the shell so Antibody can work as expected")
)

func main() {
	app.Author("Carlos Alexandro Becker <caarlos0@gmail.com>")
	app.Version("antibody version " + version)
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case bundleCmd.FullCommand():
		bundle()
	case updateCmd.FullCommand():
		update()
	case homeCmd.FullCommand():
		fmt.Println(antibodylib.Home())
	case purgeCmd.FullCommand():
		purge()
	case listCmd.FullCommand():
		list()
	case initCmd.FullCommand():
		sh, err := shell.Init()
		app.FatalIfError(err, "failed to init")
		fmt.Println(sh)
	}
}

func bundle() {
	var input io.Reader
	if !terminal.IsTerminal(int(os.Stdin.Fd())) && len(*bundles) == 0 {
		input = os.Stdin
	} else {
		input = bytes.NewBufferString(strings.Join(*bundles, " "))
	}
	sh, err := antibodylib.New(antibodylib.Home(), input, *parallelism).Bundle()
	app.FatalIfError(err, "failed to bundle")
	fmt.Println(sh)
}

func update() {
	var home = antibodylib.Home()
	fmt.Printf("Updating all bundles in %v...\n", home)
	var err = project.Update(home, *parallelism)
	app.FatalIfError(err, "failed to update")
}

func purge() {
	fmt.Println("Removing", *purgee)
	var err = os.RemoveAll(project.New(antibodylib.Home(), *purgee).Folder())
	app.FatalIfError(err, "failed to purge")
}

func list() {
	home := antibodylib.Home()
	projects, err := project.List(home)
	app.FatalIfError(err, "failed to list bundles")
	for _, b := range projects {
		fmt.Println(filepath.Join(home, b))
	}
}

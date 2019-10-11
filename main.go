package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/getantibody/antibody/antibodylib"
	"github.com/getantibody/antibody/project"
	"github.com/getantibody/antibody/shell"
	"github.com/getantibody/folder"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/alecthomas/kingpin.v2"
)

// nolint: gochecknoglobals
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
	purgee    = purgeCmd.Arg("bundle", "bundle to be purged").Required().String()
	listCmd   = app.Command("list", "lists all currently installed bundles").Alias("ls")
	pathCmd   = app.Command("path", "prints the path of a currently cloned bundle")
	pathee    = pathCmd.Arg("bundle", "bundle in which to find and print cloned path").Required().String()
	initCmd   = app.Command("init", "initializes the shell so Antibody can work as expected")
)

// nolint: gochecknoinits
func init() {
	log.SetOutput(os.Stderr)
	log.SetPrefix("antibody: ")
	log.SetFlags(0)
}

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
		fmt.Println(findHome())
	case purgeCmd.FullCommand():
		purge()
	case listCmd.FullCommand():
		list()
	case pathCmd.FullCommand():
		path()
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
	sh, err := antibodylib.New(findHome(), input, *parallelism).Bundle()
	app.FatalIfError(err, "failed to bundle")
	fmt.Println(sh)
}

func update() {
	var home = findHome()
	fmt.Printf("Updating all bundles in %v...\n", home)
	var err = project.Update(home, *parallelism)
	app.FatalIfError(err, "failed to update")
}

func purge() {
	fmt.Printf("Removing %s...\n", *purgee)
	project, err := project.New(findHome(), *purgee)
	if err != nil {
		app.Fatalf(err.Error())
	}
	var path = project.Path()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		app.Fatalf("%s does not exist on expected location: %s", *purgee, path)
	}
	app.FatalIfError(os.RemoveAll(path), "failed to purge")
	fmt.Println("removed!")
}

func list() {
	var home = findHome()
	projects, err := project.List(home)
	app.FatalIfError(err, "failed to list bundles")
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 4, ' ', tabwriter.TabIndent)
	for _, b := range projects {
		fmt.Fprintf(w, "%s\t%s\n", folder.ToURL(b), filepath.Join(home, b))
	}
	app.FatalIfError(w.Flush(), "failed to flush")
}

func path() {
	proj, err := project.New(findHome(), *pathee)
	if err != nil {
		app.Fatalf(err.Error())
	}
	var path = proj.Path()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		app.Fatalf("%s does not exist in cloned paths", *pathee)
	} else {
		fmt.Println(path)
	}
}

func findHome() string {
	h, err := antibodylib.Home()
	if err != nil {
		app.Fatalf("could't get cache folder: %v", err)
	}
	return h
}

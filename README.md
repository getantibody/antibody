<img src="logo.png" align="right" width="192px"/>

A faster and simpler antigen written in Golang.

[![License](https://img.shields.io/github/license/getantibody/antibody.svg?style=flat-square)](/LICENSE.md) [![Build Status](https://img.shields.io/circleci/project/getantibody/antibody/master.svg?style=flat-square)](https://circleci.com/gh/getantibody/antibody) [![Coverage Status](https://img.shields.io/coveralls/getantibody/antibody.svg?style=flat-square)](https://coveralls.io/github/getantibody/antibody?branch=master) [![Go Report Card](http://goreportcard.com/badge/getantibody/antibody)](http://goreportcard.com/report/getantibody/antibody) [![](https://godoc.org/github.com/getantibody/antibody?status.svg)](http://godoc.org/github.com/getantibody/antibody) [![Join the chat at https://gitter.im/getantibody/antibody](https://badges.gitter.im/getantibody/antibody.svg)](https://gitter.im/getantibody/antibody?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

> "Antigen is a small set of functions that help you easily manage your shell
> (zsh) plugins, called bundles. The concept is pretty much the same as
> bundles in a typical vim+pathogen setup. Antigen is to zsh, what Vundle
> is to vim."
>
> Read more: [Antigen](https://github.com/zsh-users/antigen).


### Why?

Antigen is really nice, but it is bloated and it is slow - 5+ seconds to load
on my Mac... that's way too much to wait for a prompt to load!

[![asciicast](https://asciinema.org/a/3ltopcz7dsid3iqgsg58o2sr8.png)](https://asciinema.org/a/3ltopcz7dsid3iqgsg58o2sr8)

I'm aware that there is other attempts, like
[antigen-hs](https://github.com/Tarrasch/antigen-hs), but I don't want to
install a lot of stuff for this to work.

So, why Go, you might ask: Well, the compiled Go program runs anywhere
and doesn't depend on any shared libraries. I also don't need to source it as
it would be necessary with plain simple shell. I also can do stuff in
parallel with Go routines. The little amount of shell written is needed
because I can't source something from inside a Go program (or at least
don't yet know how to do it).

### What works

These are the only antigen commands I ever used:

- `bundle`
- `update`
- `apply`

Antibody does just those three things, but you don't even need to `apply`.
Running `antibody bundle` will already download and apply the given bundle.

`antibody home` also shows where the repositories are being downloaded.

### What doesn't work

- Modules that are not in GitHub (you can open a PR if you wish);
- The `theme` command (although most themes might just work with `bundle`);
- oh-my-zsh support: it looks very ugly to me and I won't do it;

### Install

The simplest way to install Antibody is to run:

```console
$ curl -s https://raw.githubusercontent.com/getantibody/installer/master/install | bash -s
$ echo 'source <(antibody init)' >> ~/.zshrc
```

This will put the binary in `/usr/local/bin/antibody` and setup your `~/.zshrc` to
load what is needed on startup.

### Usage

Now, you can just `antibody bundle` stuff, e.g.,
`antibody bundle caarlos0/jvm`. The repository will be cloned at
your `XDG_CACHE` folder and antibody will try to load files that match:

- `*.plugin.zsh`
- `*.zsh`
- `*.sh`
- `*.zsh-theme`

You can also specify a branch to download, for example,
`antibody bundle caarlos0/jvm v2` will download the `v2` branch of that
repository.

When you decide to update your bundles, just run `antibody update`: it will
update all bundles inside the `antibody home` folder.

### Protips

Prefer to use it like this:

```sh
$ cat plugins.txt
caarlos0/jvm
djui/alias-tips
caarlos0/zsh-mkc
zsh-users/zsh-completions
caarlos0/zsh-open-github-pr
zsh-users/zsh-syntax-highlighting
zsh-users/zsh-history-substring-search

$ antibody bundle < plugins.txt
```

This way antibody can concurrently clone the bundles and find their sourceable
files, so it will probably be faster than call each one separately.

### In the wild

- I did this mostly for myself, so, my
[dotfiles](https://github.com/caarlos0/dotfiles);
- @mkwmms' [dotfiles](https://github.com/mkwmms/dotfiles);
- @oieduardorabelo's [dotfiles](https://github.com/oieduardorabelo/dotfiles);
- @nisaacson's [dotfiles](https://github.com/nisaacson/dotfiles);
- @pragmaticivan's [dotfiles](https://github.com/pragmaticivan/dotfiles);
- @wkentaro's [dotfiles](https://github.com/wkentaro/dotfiles);
- @marceldias' [dotfiles](https://github.com/marceldiass/dotfiles);
- @davidkna's [dotfiles](https://github.com/davidkna/dotfiles);
- and probably [many others](https://github.com/search?q=antibody&type=Code);

### Static loading

You can use antibody in a static-loading manner (so you don't need to exec
antibody every time you open a shell).

```sh
$ antibody bundle < bundles.txt >> sourceables.sh
# In your zshrc (or whatever):
$ source sourceables.sh
```

Beware that antibody does stuff in parallel, so bundle order is not guaranteed.

### Thanks

- [@pragmaticivan](https://github.com/pragmaticivan), for the logo design;
- All the amazing [contributors](https://github.com/getantibody/antibody/graphs/contributors).

<img src="logo.png" align="right" width="192px" height="192px"/>

A faster and simpler version of antigen written in Go.

[![License](https://img.shields.io/github/license/caarlos0/antibody.svg?style=flat-square)](/LICENSE.md) [![Build Status](https://img.shields.io/circleci/project/caarlos0/antibody/master.svg?style=flat-square)](https://circleci.com/gh/caarlos0/antibody) [![Coverage Status](https://coveralls.io/repos/caarlos0/antibody/badge.svg?branch=master&service=github)](https://coveralls.io/github/caarlos0/antibody?branch=master) [![Stories in Ready](https://badge.waffle.io/caarlos0/antibody.png?label=ready&title=Ready)](https://waffle.io/caarlos0/antibody)

> "Antigen is a small set of functions that help you easily manage your shell
> (zsh) plugins, called bundles. The concept is pretty much the same as
> bundles in a typical vim+pathogen setup. Antigen is to zsh, what Vundle
> is to vim."
>
> Read more: [Antigen](https://github.com/zsh-users/antigen).


### Why?

Antigen is really nice, but it is bloated and it is slow - 5+ seconds to load
on my Mac... that's way too much to wait for a prompt to load!

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

### What doesn't work

- Modules that are not in GitHub (you can open a PR if you wish);
- The `theme` command (although some themes might just work with `bundle`);
- oh-my-zsh support: it looks very ugly to me and I won't do it;

### Usage

- Download the [latest release](https://github.com/caarlos0/antibody/releases/).
If are using OSX and homebrew, you can use the
[`homebrew tap`](https://github.com/caarlos0/homebrew-antibody) for that;
- Uncompress it somewhere (uneeded with `homebrew`);
- [`source antibody.zsh`](https://github.com/caarlos0/dotfiles/blob/master/antibody/path.zsh).

Pay attention to not put the `antibody` binary in your `PATH`. This will cause
antibody to malfunction. You just need to source the `antibody.zsh` for it
to work.

Now, you can just `antibody bundle` stuff, e.g.,
`antibody bundle Tarrasch/zsh-autoenv`. The repository will be cloned at
your `XDG_CACHE` folder and antibody will try to load some files that match:

- `*.plugin.zsh`
- `*.zsh`
- `*.sh`
- `*.zsh-theme`

When you decide to update your bundles, just run `antibody update`.

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
[dotfiles](https://github.com/caarlos0/dotfiles/pull/78);
- @mkwmms' [dotfiles](https://github.com/mkwmms/dotfiles);
- @oieduardorabelo's [dotfiles](https://github.com/oieduardorabelo/dotfiles);
- @nisaacson's [dotfiles](https://github.com/nisaacson/dotfiles);
- @pragmaticivan's [dotfiles](https://github.com/pragmaticivan/dotfiles);
- @wkentaro's [dotfiles](https://github.com/wkentaro/dotfiles);
- @marceldias' [dotfiles](https://github.com/marceldiass/dotfiles);
- @davidkna's [dotfiles](https://github.com/davidkna/dotfiles);
- and probably [some more](https://github.com/search?q=antibody&type=Code);

### Hacking

#### Static loading

In [#67](https://github.com/caarlos0/antibody/issues/67) I was asked if there
is some sort of static loading.

Short answer: no, there isn'. But you can hack arount it.

If you want to use antibody just to download and/or update your dependencies,
you can run it like this:

```sh
$ ANTIBODY_FOLDER/bin/antibody bundle < bundles.txt | xargs -I {} echo "source {}" >> sourceables.sh
# In your zshrc (or whatever):
$ source sourceables.sh
```

With this approach you don' even need to source `antibody.zsh` if you don't
want to, and, yes, your shell will probably be even faster. It comes with
the cost of additional work though.

### Throughput Graph

[![Throughput Graph](https://graphs.waffle.io/caarlos0/antibody/throughput.svg)](https://waffle.io/caarlos0/antibody/metrics)

### Thanks

- [@pragmaticivan](https://github.com/pragmaticivan), for the logo design.

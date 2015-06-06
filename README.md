# antibody [![Build Status](https://travis-ci.org/caarlos0/antibody.svg?branch=master)](https://travis-ci.org/caarlos0/antibody) [![Coverage Status](https://coveralls.io/repos/caarlos0/antibody/badge.svg?branch=master)](https://coveralls.io/r/caarlos0/antibody?branch=master)

A faster and simpler version of antigen

### Why?

Antigen is really nice. But it is bloated and it is slow - 5+ seconds to load
on my Mac... that's way too much to wait for a prompt to load!

I'm aware that there is other attempts, like
[antigen-hs](https://github.com/Tarrasch/antigen-hs), but I'm don't want to
install a lot of stuff for this to work.

So, why Go, you might ask.

Well, the compiled Go program run anywhere and doesn't depend on any shared
libraries. I also don't need to source it as it would be necessary with
plain old shell. The amount of shell written is needed because I can't source
something from inside a Go program (or at least don't yet know how to do it).

### What works

The only two antigen commands I ever used:

- `bundle`
- `update`

You don't even need apply. Running `antibody bundle` will already apply the
bundled plugin.

### What doesn't work

- Modules without a `*.plugin.zsh` file;
- Modules that are not in GitHub (you can open a PR if you wish);
- The `theme` command (although some themes might just work with bundle);
- oh-my-zsh support: it looks very ugly to me and I won't do it;

### Usage

- Download the [last release](https://github.com/caarlos0/antibody/releases/);
- Uncompress it somewhere;
- `source antibody.zsh`.

Now, you can just `antibody bundle` stuff, e.g.,
`antibody bundle Tarrasch/zsh-autoenv`. The repository will be cloned at
`~/.antibody` and all `.zsh.plugin` files will be loaded.

If you ever need to update your bundles, just run `antibody update`.

### In the wild

- I did this mostly for myself, so, my [dotfiles](https://github.com/caarlos0/dotfiles/pull/78);

---
title: Outro
weight: 500
menu: true
---

Let's look what other commands antibody has available for us!

## Update

Antibody can update all bundles in a single pass.

Just run:

```console
$ antibody update
Updating all bundles in /Users/carlos/Library/Caches/antibody...
```

and that's it.

## Purge

You can remove a bundle completely by purging it:

```console
$ antibody purge ohmyzsh/ohmyzsh
Removing ohmyzsh/ohmyzsh...
```

## List

If you want to see what plugins you have in your home folder, you can of
course list them:

```console
$ antibody list
https://github.com/Tarrasch/zsh-bd            /Users/carlos/Library/Caches/antibody/https-COLON--SLASH--SLASH-github.com-SLASH-Tarrasch-SLASH-zsh-bd
https://github.com/caarlos0/git-add-remote    /Users/carlos/Library/Caches/antibody/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-git-add-remote
# ...
```

## Path

You can see the path being used for a cloned bundle.

```console
$ antibody path ohmyzsh/ohmyzsh
/Users/carlos/Library/Caches/antibody/https-COLON--SLASH--SLASH-github.com-SLASH-ohmyzsh-SLASH-ohmyzsh
```

This is particularly useful for projects like oh-my-zsh that rely on
storing its path in the `$ZSH` environment variable:

```console
$ ZSH=$(antibody path ohmyzsh/ohmyzsh)
```

## Home

You can also see where antibody is keeping the plugins with the home
command:

```console
$ antibody home
/Users/carlos/Library/Caches/antibody
```

Of course, you can remove the entire thing with:

```sh
rm -rf `antibody home`
```

if you decide to start fresh or to use something else.

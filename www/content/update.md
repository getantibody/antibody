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
$ antibody purge robbyrussell/oh-my-zsh
Removing robbyrussell/oh-my-zsh...
```

## List

If you want to see what plugins you have in your home folder, you can of
course list them:

```console
$ antibody list
/Users/carlos/Library/Caches/antibody/https-COLON--SLASH--SLASH-github.com-SLASH-Tarrasch-SLASH-zsh-bd
/Users/carlos/Library/Caches/antibody/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-git-add-remote
# ...
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

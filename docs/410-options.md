---
title: Options
---

There are a few options you can use that should cover most common use cases.
Let's take a look!

## Kind

The `kind` annotation can be used to determine how a bundle should be treated.

The default is `kind:zsh`, which will look for files that match these globs:

- `*.plugin.zsh`
- `*.zsh`
- `*.sh`
- `*.zsh-theme`

And `source` them.

Example:

```console
$ antibody bundle caarlos0/jvm kind:zsh
source /Users/carlos/Library/Caches/antibody/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-jvm/jvm.plugin.zsh
```

The `kind:path` mode will just put the plugin folder in your `$PATH`.

Example:

```console
$ antibody bundle caarlos0/ports kind:path
export PATH="/Users/carlos/Library/Caches/antibody/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-ports:$PATH"
```

## Branch

You can also specify a branch to download, if you don't want the `master` branch
for whatever reason.

Example:

```console
$ antibody bundle caarlos0/jvm branch:v2
source /Users/carlos/Library/Caches/antibody/https-COLON--SLASH--SLASH-github.com-SLASH-caarlos0-SLASH-jvm/jvm.plugin.zsh
```

## Sub-folders

You may specify a subfolder if the repo you are bundling contains multiple
plugins.

Example:

```console
$ antibody bundle robbyrussell/oh-my-zsh folder:plugins/aws
source /Users/carlos/Library/Caches/antibody/https-COLON--SLASH--SLASH-github.com-SLASH-robbyrussell-SLASH-oh-my-zsh/plugins/aws/aws.plugin.zsh
```

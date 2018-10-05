---
title: Usage
weight: 400
menu: true
---

There are mainly two ways of using antibody: static and dynamic.
We will also see how we can keep a plugins file.

## Plugins file

A plugin file is basically any text file that has one plugin per line.

In our examples, let's assume we have a `~/.zsh_plugins.txt` with these
contents:

```sh
caarlos0/jvm
djui/alias-tips
# comments are supported like this
caarlos0/zsh-mkc
zsh-users/zsh-completions
caarlos0/zsh-open-github-pr

# empty lines are skipped

# annotations are also allowed:
robbyrussell/oh-my-zsh path:plugins/aws

zsh-users/zsh-syntax-highlighting
zsh-users/zsh-history-substring-search
```

That being said, let's look how can we load them!

## Dynamic loading

This is the most common way. Basically, every time the a new shell starts,
antibody will apply the plugins given to it.

For this to work, antibody needs to be wrapped into your `~/.zshrc`. To do
that, run:

```sh
# ~/.zshrc
source <(antibody init)
```

And reload your current shell or open a new one.

Then, you will also need to tell antibody which plugins to bundle.
This can also be done in the `~/.zshrc` file:

```sh
# ~/.zshrc
antibody bundle < ~/.zsh_plugins.txt
```

## Static loading

This is the faster alternative. Basically, you'll run antibody only when
you change your plugins, and then you can just load the "static" plugins file.

Note that in this case, **we should not put `antibody init` on our `~/.zshrc`**.
If you did that already, remove it from your `~/.zshrc` and start a fresh
terminal session.

Assuming the same `~/.zsh_plugins.txt` as before, we can run:

```sh
antibody bundle < ~/.zsh_plugins.txt > ~/.zsh_plugins.sh
```

At any time to update our `~/.zsh_plugins.sh` file. Now, we just need to
`source` that file on `~/.zshrc`:

```sh
# ~/.zshrc
source ~/.zsh_plugins.sh
```

And that's it!

## CleanMyMac and others

If you use CleanMyMac or similar tools, make sure to set it up to ignore the
`antibody home` folder, otherwise it may delete your plugins.

You may also change Antibody's home folder, for example:

```sh
export ANTIBODY_HOME=~/Libary/antibody
```

---
title: Compatibility
---

Since antibody started as a subset clone of antigen, one might wonder
how compatible one is with another. Let's take a look.

Antibody can only `bundle` and `update` plugins. The `apply` command is not
needed because running `antibody bundle` will already download and apply the
given plugin.

The `theme` command is not implemented. You can just use `bundle` instead.

oh-my-zsh plugins are supported by using the [folder annotation](#options.sub_folders):

```sh
antibody bundle robbyrussell/oh-my-zsh folder:plugins/aws
```

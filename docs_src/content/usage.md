+++
date = "2017-02-04T20:51:38-05:00"
title = "Usage"
[menu.main]
  weight = 10
+++

Find takes an argument a filename expression and outputs a list of all files whose
names match that expression below the current working directory.

For example:

```bash
$ go-find "*.go"
```

Will find all files with the extension *.go*.

## Additional flags

```
Usage of go-find:
  -dir string
        Directory for search (default ".")
```

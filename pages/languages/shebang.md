# Shebang

## Description

A _shebang_ or _hashbang_ is a particular line that often begins scripts. The
line begins with the characters `#!`. This line defines which program should be
passed the contents of the file. For example:

```bash
#!/bin/sh

echo "Hi there!"
```

## `/usr/bin/env`

The paths in the shebang line may be different on different machines. The `env`
program exists to solve this problem and make scripts portable. This program
looks in the user's `PATH` to find the location of an application. The path
`/usr/bin/env` is commonly used for this utility.

```bash
#!/usr/bin/env sh

echo "Hi there!"
```

## Script Execution

Once a shebang line exists, the `chmod` command can be used to make a script
executable.

```bash
chmod +x script.sh
./script.sh
```

The first line adds (`+`) the executable (`x`) permission to the file.

# Bash

Bash (Bourne Again SHell) is a command language and Unix shell. Bash is a
superset of Bourne shell (`sh`); commands that work in `sh` also work in `bash`.

## When to Use Bash

The [Google Shell Style Guide](https://google.github.io/styleguide/shell.xml)
gives advice on when to use shell.

> Shell should only be used for small utilities or simple wrapper scripts.
>
> - If you're mostly calling other utilities and are doing relatively little
>   data manipulation, shell is an acceptable choice for the task.
> - If you are writing a script that is more than 100 lines long, you should
>   probably be writing it in Python instead.

## Variables

```bash
# Use uppercase for constants.
NAME="Elliot"
echo "Hi, ${NAME}"

# Use snake_case for variables.
count=0
(( count++ ))
echo ${count}
```

Prefer `${var}` over `$var` for variable expansion.

## Functions

```bash
# Use snake_case for function names.
greet() {
    # Declare function-specific variables with 'local'.
    local name="$1"
    echo "Hello, ${name}"
}

greet "Elliot"
```

### Arguments

Use _special parameters_ to access parameters given to the function.

| Expression | Description         |
| ---------- | ------------------- |
| `$#`       | Number of arguments |
| `$@`       | All arguments       |
| `$1`       | First argument      |
| `$2`       | Second argument     |

## Control Flow

```bash
if [[ expression ]]; then
    ...
else
    ...
fi
```

### Conditionals

| Expression               | Description      |
| ------------------------ | ---------------- |
| `[[ NUM -eq NUM ]]`      | Numbers equal    |
| `[[ STRING == STRING ]]` | Strings equal    |
| `[[ ! EXPR ]]`           | Not              |
| `[[ X ]] && [[ Y ]]`     | And              |
| `[[ X ]] ❘❘ [[ Y ]]`     | Or               |
| `[[ -z STRING ]]`        | Empty string     |
| `[[ -e FILE ]]`          | File exists      |
| `[[ -d FILE ]]`          | Directory exists |

## Loops

```bash
for i in {1..5}; do
    ...
done
```

## Command Substitution

Command substitution is used to insert the output of one command into a second
command.

```bash
today=$(date)
echo "${today}"
```

## Style

Start each file with a [[shebang]] and a description of the contents.

```bash
#!/bin/bash
#
# Display an example.
```

Executables (`chmod +x file`) should have no extension.

## `set`

Use `set` to configure shell. Many scripts start with `set -eou pipefail`.

| Argument      | Description                         |
| ------------- | ----------------------------------- |
| `-e`          | Stop on first non-zero exit status  |
| `-o pipefail` | Don't mask errors in pipelines      |
| `-u`          | Treat undefined variables as errors |
| `-x`          | Print commands before execution     |

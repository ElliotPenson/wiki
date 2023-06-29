# Vim

## Basic

| Command          | Description                       |
| ---------------- | --------------------------------- |
| `:e <file>`      | Open file                         |
| `:w`             | Save current buffer               |
| `:q`             | Quit                              |
| `:bn`, `:bp`     | Next, previous buffer             |
| `i`              | Enter insert mode                 |
| `v`, `V`         | Enter visual mode, line-wise      |
| `<esc>`          | Return to normal mode             |
| `<c-u>`, `<c-d>` | Move up, down 1/2 page            |
| `<c-w> <c-s>`    | Split horizontally                |
| `<c-w> <c-v>`    | Split vertically                  |
| `<c-w> <c-h>`    | Change to window left             |
| `<c-w> <c-j>`    | Change to window below            |
| `<c-w> <c-k>`    | Change to window above            |
| `<c-w> <c-l>`    | Change to window right            |
| `<c-w> q`        | Close current window              |
| `<c-w> o`        | Close all windows but the current |
| `zz`             | Center on cursor                  |

## Navigation

| Command              | Description                        |
| -------------------- | ---------------------------------- |
| `h`, `j`, `k`, `l`   | Left, down, up, right              |
| `w`, `W`             | Start of word, WORD                |
| `e`, `E`             | End of word, WORD                  |
| `b`, `B`             | Backwards by word, WORD            |
| `0`, `$`             | Start, end of line                 |
| `^`                  | First non-blank character of line  |
| `gg`, `G`            | Beginning, end of file             |
| `/<text>`            | Search forwards                    |
| `?<text>`            | Search backwards                   |
| `*`                  | Search word under cursor forwards  |
| `#`                  | Search word under cursor backwards |
| `n`, `N`             | Next, previous in search results   |
| `''`                 | Last location                      |
| `f<char>`, `F<char>` | Next, previous character           |
| `:<num>` or `<num>G` | Line number                        |

## Editing

| Command  | Description                |
| -------- | -------------------------- |
| `i`, `I` | Insert before cursor, line |
| `a`, `A` | Append after cursor, line  |
| `o`, `O` | Open line below, above     |
| `J`      | Join lines                 |
| `yy`     | Yank line                  |
| `dd`     | Delete line                |
| `D`      | Delete rest of line        |
| `x`      | Delete character           |
| `p`, `P` | Paste after, before        |
| `u`      | Undo                       |
| `.`      | Do again                   |
| `<c-a>`  | Increment number           |
| `<c-x>`  | Decrement number           |
| `>>`     | Indent line right          |
| `<<`     | Indent line left           |
| `>`      | Indent selection right     |
| `<`      | Indent selection left      |

## Macros

| Command     | Description                                  |
| ----------- | -------------------------------------------- |
| `q<letter>` | Start recording macro in register `<letter>` |
| `q`         | Stop recording macro                         |
| `@<letter>` | Execute macro in register `<letter>`         |
| `@@`        | Execute macro again                          |

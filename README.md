# fswatch

A simple file watcher for the command line.
It takes a path to watch and a command to execute
when that file is changed (write events only), with
the file name as the argument.

## Installation

```
go get github.com/ecix/fswatch
```

## Usage

```bash
# watch all files in a directory `mydir` and print their names on change
watch -dir mydir -cmd echo
```

## Contributors

Written by [Veit Heller](https://github.com/hellerve) of Port Zero on
behalf of [ecix](https://github.com/ecix).

<hr/>

Have fun!

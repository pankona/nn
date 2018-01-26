# nn

Shows "n 日ぶり n 回目"

## Install

```bash
$ go get -u github.com/pankona/nn
```

## Usage

```bash
NAME:
   nn - show nn

USAGE:
   nn [options] [arguments...]

VERSION:
   0.1

OPTIONS:
   --id value                specify id of nn (default: "default")
   -f value, --format value  specify format of nn (default: "%d 日ぶり %d 回目")
   --help, -h                show help
   --version, -v             print the version
``` 

## Usage example

### Default

```bash
$ nn
0 日ぶり 1 回目

$ nn
0 日ぶり 2 回目
```

### Specify ID

count is identified by specified "id"

```bash
$ nn -id myapp1
0 日ぶり 1 回目

$ nn -id myapp1
0 日ぶり 2 回目

$ nn -id myapp2
0 日ぶり 1 回目

$ nn -id myapp2
0 日ぶり 2 回目
```

### Change display format

display format can be changed by specifying "-f" or "--format".  
it must specify two "%d", like below.


```bash
$ nn -f "while in %d days, %d time"
while in 0 days, 4 time
```

### Reset

nn records are stored on `&HOME/.config/nn/{id}.txt`, just remove it.

## License

MIT



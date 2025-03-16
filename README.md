# MyIP

Fetch your public IP address from the command line

## Create binary (check releases for pre-built binaries)

***Requires Go version 1.23+***

```bash
cd myip
go build -o myip main.go
mv myip /your/binary/path/bin
```

**Optimized Build (stripping debug information)**

Use the -ldflags flag with the -w and -s options.
- w omits the DWARF symbol table, removing debugging information.
- s strips the symbol table and debug information.

```bash
cd myip
go build -ldflags="-w -s" -o myip main.go
mv myip /your/binary/path/bin
```

## Add to PATH

```zsh
export MYIP="$HOME/.me"
export PATH="$MYIP/bin:$PATH"
```

## Usage

```bash
Usage: myip v4 (or v6)
```

**Get IPv4 Address:**
```bash
myip v4
```

**Get IPv6 Address:**
```bash
myip v6
```


## Notes

This script uses the [ifconfig.me](https://ifconfig.me) website. Its a simple wrapper.

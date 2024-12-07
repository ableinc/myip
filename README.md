# MyIP

Fetch your public IP address from the command line

## Create binary (check releases for pre-built binaries)

***Requires Go version 1.23+***

```bash
cd myip
go build -o myip main.go
mv main.go /your/binary/path/bin
```

## Add to PATH

```zsh
export MYIP_INSTALL="$HOME/.me"
export PATH="$P_INSTALL/bin:$PATH"
```

## Usage

```bash
Usage: myip [options]
Options:
  -v int
    	Which IP version to display (default 4)
```

**Get IPv4 Address:**
```bash
myip
```

**Get IPv6 Address:**
```bash
myip -v 6
```


## Notes

This script uses the [ifconfig.me](https://ifconfig.me) website. It is a simple wrapper.

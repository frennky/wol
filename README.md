# wol

## Install

```
go install github.com/frennky/wol@latest
```
## Usage

```
Simple Wake On LAN CLI tool.

Usage:
  wol [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  wake        Wake On LAN

Flags:
  -h, --help   help for wol

Use "wol [command] --help" for more information about a command.
```

### Wake command

```
Sends magic packet to target MAC.

Usage:
  wol wake [flags]

Flags:
  -h, --help          help for wake
      --host string   Broadcast IP Address (optional) (default "255.255.255.255")
      --mac string    Target MAC Address (required)
      --port int      Broadcast Port (optional)
```

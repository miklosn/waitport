# waitport

Waits for a tcp port to be in LISTEN state. Useful for waiting for dependencies in a docker entrypoint etc.

Supports a configurable timeout (default: 30s) and exits 1 if it is exceeded. 

## Install

`curl -s i.jpillora.com/miklosn/waitport | bash`

## Example

```
waitport -t 10s 9000 || exit 1
```

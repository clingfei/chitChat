# chitChat

1. make go env, set $GOPATH and add  $GOPATH/bin to PATH env
```
mkdir go && cd go
mkdir src pkg bin
export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin
```
2. build && run
```
go build && ./chitChat
```
3. test
```
go test
go test -v -cover

```
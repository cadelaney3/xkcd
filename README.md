# xkcd

## Description

This project includes a package that allows to create an offline index of xkcd comics
and then query the index with command line arguments.

### Installation

Make sure to have Go installed on your device.

```
go get github.com/cadelaney3/xkcd
cd $GOPATH/src/github.com/xkcd
go build
./xkcd [search-terms...]
```
Or, if you run
```
go env
```
and you see: GOROOT="/usr/local/bin", you can do from any directory:
```
git clone https://github.com/cadelaney3/xkcd
cd xkcd
go build
./xkcd [search-terms...]
```
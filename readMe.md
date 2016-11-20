[![Build Status](https://travis-ci.org/andiMenge/restAPISkeleton.svg?branch=master)](https://travis-ci.org/andiMenge/restAPISkeleton)

trigger api 

```
curl http://localhost:1337/
curl http://localhost:1337/foo
```

start API

```
go build .
./restAPISkeleton -h
./restAPISkeleton
```

test API

```
curl -v 'http://localhost:1337/set?key=foo&value=bar'
curl -v 'http://localhost:1337/get?key=foo'
```

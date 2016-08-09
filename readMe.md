trigger api 

```
curl -H "Content-Type: application/json" -d '{"name":"foobar"}' http://localhost:1337/queueJob
```

start API

```
go run ./*
```

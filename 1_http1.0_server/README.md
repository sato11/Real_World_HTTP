## HTTP/0.9
Simulate HTTP/0.9 and see what server outputs to logs.
```sh
$ curl --http1.0 http://localhost:18888/greeting
$ curl --http1.0 --get --data-urlencode "search word" http://localhost:18888
```

## HTTP/1.0
```sh
$ curl --http1.0 -v http://localhost:18888:
```

## Send header
```sh
$ curl --http1.0 -H "X-Test: Hello" http://localhost:18888
```

## Redirect
```sh
$ curl -L http://localhost:18888
```

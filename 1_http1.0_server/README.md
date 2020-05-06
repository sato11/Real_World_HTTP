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

## x-www-form-urlencoded
Use `--data-urlencode` instead of `-d` to escape special characters.
```sh
$ curl --http1.0 --data-urlencode title="Head First PHP & MySQL" --data-urlencode author="Lynn Beighley, Michael Morrison" localhost:18888
```

## multipart/form-data
Use `-F` option to create `Content-Type: multipart/form-data`.

```sh
$ curl --http1.0 -F title="The Art of Community" -F author="Jono Bacon" -F attachment-file=@test.txt localhost:18888

# or specify metadata
$ curl --http1.0 -F title="The Art of Community" -F author="Jono Bacon" -F "attachment-file=@test.txt;type=text/html" localhost:18888
$ curl --http1.0 -F title="The Art of Community" -F author="Jono Bacon" -F "attachment-file=@test.txt;filename=sample.txt" localhost:18888
```

## encoding
Use `--compressed` as a shorthand for `-H "Accept-Encoding: deflate, gzip"`.
```sh
$ curl --http1.0 --compressed localhost:3000
```

## cookie
Use `-c` to save cookies returned by the target server in a file, and `-b` to use the saved cookie to make a request.
```sh
$ curl -I -c cookie.txt -b cookie.txt https://www.google.com
```

## Digest authentication
Use `--digest` to authenticate user with hash algorithm.
```sh
$ curl --http1.0 --digest -u user:pass localhost:18888/digest -v
```

## Proxy
Use `-x/--proxy` to use a server as proxy. Our server implementation, however, does not forward request not return response from `example.com`.
```sh
$ curl --http1.0 -x http://localhost:18888 -U user:pass http://example.com/helloworld
```

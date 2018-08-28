## Checking out webassembly in Go with the Go1.11 release


```
$ GOARCH=wasm GOOS=js go build -o test.wasm main.go
```

To Copy over the html & JS Support files

```
cp $(go env GOROOT)/misc/wasm/wasm_exec.{html,js} .
```
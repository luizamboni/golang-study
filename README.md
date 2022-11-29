Golang Study
===
Its a dummy project only for my personal studies.

GOROOT:

GOPATH:


# 1 - Output
Explore output methods

# 2 - Variables
Explore variables:
structs, contants, primitives etc

# 8 - Object Orientation
ps.: use commands bellow in example directory (8-oo)

up redis
```bash
sudo docker run -d -p 6379:6379 -i -t redis:3.2.5-alpine
```

need install external modules (go modules)
```bash
go get -u github.com/go-redis/redis/v7
```

to test all test
```bash
go  test -v *_test.go
```

# 11 - go modules
created with  
```
go mod init go-study.com/m 
```

To install a public module
```
go get github.com/masatana/go-textdistance
```

To add dependencies in vendor package 
```
go mod vendor
```

and run
```
go run main.go
```
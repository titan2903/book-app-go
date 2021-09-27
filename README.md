# How To Start Development
**1.** Clone this project inside your GOPATH src file 
```shell
$ git clone https://github.com/titan2903/book-app-go.git book-app
$ cd book-app
```

**2.** Install Prerequisite
1. Install go-migrate `https://github.com/golang-migrate/migrate` for running migration

**3.** Migration
1. Run below command to run migration

```shell
$ migrate -path migration -database "mysql://user:password@tcp(host:port)/dbname?query" up
```

2. To create a new migration file

```shell
$ migrate create -ext sql -dir migration -seq name
```

3. or execute command in `Makefile`
```shell
$ make -f Makefile
```

**4.** Run below command to run app server
```shell
$ go run main.go
```
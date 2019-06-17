|branch|build|test|
|------|-----|----|
|master|[![Build Status](https://travis-ci.org/eduardocerqueira/learning-go.svg?branch=master)](https://travis-ci.org/eduardocerqueira/learning-go)|[![codecov](https://codecov.io/gh/eduardocerqueira/learning-go/branch/master/graph/badge.svg)](https://codecov.io/gh/eduardocerqueira/learning-go)|
|dev|[![Build Status](https://travis-ci.org/eduardocerqueira/learning-go.svg?branch=dev)](https://travis-ci.org/eduardocerqueira/learning-go)|[![codecov](https://codecov.io/gh/eduardocerqueira/learning-go/branch/master/graph/badge.svg)](https://codecov.io/gh/eduardocerqueira/learning-go)|

# Learning GoLang

personal repo for learning and getting more familia with GoLang

#### Install

on time I am writing this, latest **go** version is 1.12.6, so:

```
# download and save locally/extract
curl https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz --output go1.12.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz

# add in your system path ( bashrc or profile )
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc

# verify
[ecerquei@desktop bin]$ go version
go version go1.12.6 linux/amd64

# get dependency
go get github.com/gorilla/mux
go get github.com/stretchr/testify/assert
```

Access official download page for latest release and more instructions if needed: https://golang.org/dl/

#### Build and Run

```
git clone XXXXX
cd learning-go/bin
# build
go build ../src/hello/hello.go
go build ../src/api/api_simple.go
go build ../src/apimux/api_mux.go

# run
./hello
./api_simple
./api_mux
```

Once running **api_mux** you can add new articles, they will be saved in memory:

```
curl -d '{"Id":"3", "Title":"Title 1", "Desc": "Article Description", "Content": "Article Content"}' -X POST http://localhost:10000/article
```

then listing by 

```
curl http://localhost:10000/articles
```

#### References and Credits

https://tutorialedge.net/golang/creating-restful-api-with-golang/
https://golang.org/
https://github.com/gorilla/mux
https://blog.golang.org/go-slices-usage-and-internals
https://golang.org/doc/effective_go.html

# Go Rest App created using Fiber hosted to Heroku.


## Installation
```
$ go get github.com/hamza-kay/GoRestApp
$ cd $GOPATH/src/github.com/hamza-kay/GoRestApp
$ go build
```

## Features

- Fiber 
- GORM
- Heroku
- sqlite
- RESTful

## API Endpoint

- https://calm-coast-05863.herokuapp.com/api/v1/book
  - ```GET```: get list of books
  - ```POST```: create a book

- https://calm-coast-05863.herokuapp.com/api/v1/book/{id}
  - ```GET```: get a specific book
  - ```DELETE```: delete a specific book


## Data Structure

    { 
        "title": "test",
        "author": "testname",
        "rating": 0 
    }











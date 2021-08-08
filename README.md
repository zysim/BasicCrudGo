# Basic CRUD in Go
This is me making a basic CRUD in Go. Literally just going through the [Go docs](https://golang.org/doc/articles/wiki).

### To Run
`go run src/server/main.go`

### What `main.go` Contains
Main entry points are the `test_<function>` functions. They are encapsulations of each section in the docs page linked above, to the best of my ability. This way, I can push aside stuff I don't currently wanna toy with at the moment.

From there, the contents of the `test_<function>` functions are what you wanna read. They're the implementations of each section of the docs. Which are, well, also from the docs.

### Notes for Myself
`loadPage` reads from the calling directory. So for example, if you're in this folder (~/Documents/../BasicCrud), and you call `go run src/server/main.go`, it will try to look for files in BasicCrud/. So your txt file's parent dir HAS to be BasicCrud.

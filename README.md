# Simple Go Server

![Cute Gopher](https://golang.org/lib/godoc/images/footer-gopher.jpg)

Simple Go Server implements a basic web server using Go to create and view files.

### Endpoints
- `/`: View all files in your folder
- `/file/<filename>`: View filename in browser
- `/edit/<filename>`: Edit file if it exists, or create new
- `/save/`: Controller to save file

### How to Run
1. Make sure to have Go installed
2. Run `go build`
3. Run the executable generated

### What I learnt
- GET / POST using Go
- Routing using Go
- Templating in Go
- File handling in Go
- Middlewares in Go
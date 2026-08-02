>go version
go version go1.26.5 windows/amd64

Go modules :
A Go Module is the standard way to organize a Go project and manage its dependencies.

>What does go.mod contain? 

Module name
Go version
Dependencies


>Difference Between go run and go build

>go run
Compiles and immediately runs the program.
## go run main.go → Suitable only if your program is contained entirely in main.go, or if you explicitly list the other files:
>go build
Compiles the program into an executable (.exe on Windows) without running it.
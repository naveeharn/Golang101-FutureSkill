# Golang101

A brief description of what this project does and who it's for

## run multi .go files
execute command in the terminal 
```bash
  go run <filename.go> [<another-filename.go>]
```
for example `go run main.go workshop.go` 

## go.mod
create go.mod file in the project at terminal for any .go file can import packages in local repository and export to github
```bash
  go mod init <module-name>
```
for example `go mod init github.com/naveeharn/golang101`

## go.sum 
create go.sum file in the project at terminal for any .go file import for packages (module name in import ()) from github 
```bash
  go mod tidy
```
for example
```go
package main

import (
	"fmt"

	"github.com/naveeharn/Golang101_FutureSkill/go20-project/movie"
)

func main() {
	movieName := movie.FindName("tt4154796")
  fmt.Println(movieName)
}
```
execute command `go mod tidy`

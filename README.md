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
create go.sum file in the project at terminal for any .go file import for packages from github through module name
### - approach 1
```zsh
  go get <module name>
```
for example code `go get github.com/naveeharn/Golang101_FutureSkill`

### - approach 2
```bash
  go mod tidy
```
for example code in main.go
```go
package main

import (
	"fmt"

	"github.com/naveeharn/Golang101_FutureSkill/go20-project/movie"	// module nmae
)

func main() {
	movieName := movie.FindName("tt4154796")
  	fmt.Println(movieName)
}
```
then execute command `go mod tidy`

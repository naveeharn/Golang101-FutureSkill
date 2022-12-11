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

## go test
1) create test file in the project for any .go file and write code fot testing in it
```bash
  <filename>_test.go
```
`P.S.` before start testing, **the project necessary has go.mod**\

2) execute any command at current directory in terminal
```bash
  go test .
```
| command | result |
|---|---|
| `go test .`  | show **FAIL** if any testcase fail|
| `go test -v .`  | show **PASS/FAIL** all testcases |
| `go test -v ./...`  | start testing from current package to below package in directorty  |

example how to test Add function in main.go
- code in `main.go`
```go
package main

import (
	"fmt"
)

func Add(a, b float64) float64 {
	return a + b
}

func main() {
  fmt.Println(Add(1.0, 2.0))

}
```
- create go module in the project
```bash
go mod <module-name>
```
example `go mod main_testing`


- code in `main_test.go`
```go
package main

import "testing"

func TestAddWhenInput1and2ShouldBe3(t *testing.T) {
	r := Add(1.0, 2.0)
	if r != 3.0 {
		t.Error("Add(1.0, 2.0) should be 3.0")
	}
}
```
- start testing with any command from above table
```bash
go test .
```
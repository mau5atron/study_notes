# study_notes

```
Go Installation instructions:

install with brew: in terminal type : brew install go

After brew is successfully installed, you need to set a path for go.
To do this, type into terminal: 

export GOROOT=/usr/local/opt/go/libexec
export GOPATH=$HOME/.go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

Once the path is configured, test go with your fast hello world program.

```

```
First create a file inside your go folder that you specified. 

In terminal type: touch hello.go

Then open the file in your text editor of choice.

```

```
Inside your hello.go file, type in: 

```

```go
// hello.go
package main
import "fmt"

func main() {
  fmt.Printf("Hello, world!")
}

```

```
To make sure everything is running correctly, run your program. 
In terminal => go run hello.go
Output => Hello, world!

```
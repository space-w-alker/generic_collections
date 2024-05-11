## generic_collections

This project implements various generic data structures in Golang using generics (introduced in Go 1.18). 

### Supported Collections

* Stack: A Last-In-First-Out (LIFO) data structure.

### Installation

This project requires Go 1.18 or later to be installed. You can download and install it from the official website: [Golang download](https://golang.org/dl/).

Once Go is installed, clone this repository and run `go mod init` to initialize the Go module:

```bash
git clone https://your-github-url/generic_collections.git
cd generic_collections
go mod init github.com/your-username/generic_collections
```

### Usage

**Stack**

```go
package main

import (
	"fmt"

	"github.com/space-w-alker/generic_collections/stack"
)

func main() {
	s := make(stack.Stack[int], 10)
	s.Push(20)
	s.Push(9)
	v, _ := s.Pop()
	fmt.Printf("%v\n", v) //9
	v, _ = s.Pop()
	fmt.Printf("%v\n", v) //20
	_, err := s.Pop()
	fmt.Printf("%v\n", err) //true
}
```

**Note:** Replace `"./stack"` and `"./set"` with the actual import path based on your project structure.

### Contributing

We welcome contributions to this project! Please feel free to fork the repository, make changes, and submit pull requests.

### License

This project is licensed under the MIT License. See the LICENSE file for details.

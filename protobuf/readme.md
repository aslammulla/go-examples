
# Protocol Buffers (Protobuf) Setup Guide for Golang

This guide provides detailed steps to install Protocol Buffers (`protoc`) on **Windows** and **Linux**, set up your environment, and use it with Go. It also includes a working Go example from this package.

---

## 📌 Prerequisites

- Go (>= 1.18) installed ([Download Go](https://go.dev/dl/))

---

## 🔧 1. Install Protobuf Compiler (`protoc`)

### 🖥️ Windows

1. Go to the [Protobuf Releases](https://github.com/protocolbuffers/protobuf/releases) page.
2. Download the latest `protoc-*-win64.zip` (or `win32` for 32-bit) under **Assets**.
3. Extract the zip file to a folder, e.g., `C:\protoc`.
4. Add the `bin` directory to your system `PATH`:
   - Open **System Properties** > **Advanced** > **Environment Variables**.
   - Under **System variables**, find and edit `Path`.
   - Add: `C:\protoc\bin`
   - Click OK to save.
5. Open a new terminal and verify:
   ```sh
   protoc --version
   ```

### 🐧 Linux (Ubuntu/Debian)

1. Download the latest release:
   ```sh
   wget https://github.com/protocolbuffers/protobuf/releases/download/v<version>/protoc-<version>-linux-x86_64.zip
   # Replace <version> with the latest, e.g., 26.1
   unzip protoc-<version>-linux-x86_64.zip -d $HOME/protoc
   ```
2. Add to your PATH (add to `~/.bashrc` or `~/.zshrc`):
   ```sh
   export PATH="$HOME/protoc/bin:$PATH"
   source ~/.bashrc
   # or source ~/.zshrc
   ```
3. Verify installation:
   ```sh
   protoc --version
   ```

---

## 🔌 2. Install Go Protobuf Plugins

Install the Go plugin and API:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go-grpc@latest
```

Ensure your Go `bin` directory is in your `PATH`:

- **Windows:** Usually `C:\Users\<YourUser>\go\bin`
- **Linux:** Usually `$HOME/go/bin`

---

## 📦 3. Generate Go Code from .proto

Assuming your proto file is at `proto/user.proto`:

```sh
protoc --go_out=. --go_opt=paths=source_relative \
	   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	   proto/user.proto
```

This will generate Go files in the appropriate package (e.g., `userpb/user.pb.go`).

---

## 🚀 4. Example: Using Protobuf in Go

See `main.go` for a working example:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/aslammulla/go-examples/protobuf/userpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := &userpb.User{
		Id:     101,
		Name:   "Aslam",
		Email:  "aslammulla.13@gmail.com",
		Skills: []string{"Go", "Python", "AWS", "Docker"},
	}

	// Serialize to Protobuf
	protoData, err := proto.Marshal(u)
	if err != nil {
		log.Fatal("Protobuf Marshal error: ", err)
	}
	fmt.Println("Protobuf serialized size:", len(protoData))

	// Deserialize back
	var u2 userpb.User
	if err := proto.Unmarshal(protoData, &u2); err != nil {
		log.Fatal("Protobuf Unmarshal error: ", err)
	}
	fmt.Println("Protobuf deserialized object:", &u2)

	// Serialize to JSON
	jsonData, err := json.Marshal(u)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}
	fmt.Println("JSON serialized size:", len(jsonData))
	fmt.Println("JSON string:", string(jsonData))

	// Deserialize back
	var u3 userpb.User
	if err := json.Unmarshal(jsonData, &u3); err != nil {
		log.Fatal("JSON Unmarshal error: ", err)
	}
	fmt.Println("JSON deserialized object:", &u3)
}
```

**Sample Output:**

```
Protobuf serialized size: 59
Protobuf deserialized object: id:101 name:"Aslam" email:"aslammulla.13@gmail.com" skills:"Go" skills:"Python" skills:"AWS" skills:"Docker"
JSON serialized size: 99
JSON string: {"id":101,"name":"Aslam","email":"aslammulla.13@gmail.com","skills":["Go","Python","AWS","Docker"]}
JSON deserialized object: id:101 name:"Aslam" email:"aslammulla.13@gmail.com" skills:"Go" skills:"Python" skills:"AWS" skills:"Docker"
```

---

## 📄 5. Example .proto File

`proto/user.proto`:

```proto
syntax = "proto3";

package user;

option go_package = "github.com/aslammulla/go-examples/protobuf/userpb";

message User {
  int32 id = 1;
  string name = 2;
  string email = 3;
  repeated string skills = 4;
}
```

---

## 🧹 Troubleshooting

- If `protoc` is not found, check your `PATH`.
- If Go code is not generated, ensure plugins are installed and `protoc-gen-go` is in your `PATH`.
- For more, see [Protocol Buffers Go Quick Start](https://developers.google.com/protocol-buffers/docs/gotutorial).

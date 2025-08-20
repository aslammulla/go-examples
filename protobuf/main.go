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

	// ---------------- Protobuf ----------------
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

	// ---------------- JSON ----------------
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

/*
OUTPUT:
$ go run main.go
Protobuf serialized size: 59
Protobuf deserialized object: id:101 name:"Aslam" email:"aslammulla.13@gmail.com" skills:"Go" skills:"Python" skills:"AWS" skills:"Docker"
JSON serialized size: 99
JSON string: {"id":101,"name":"Aslam","email":"aslammulla.13@gmail.com","skills":["Go","Python","AWS","Docker"]}
JSON deserialized object: id:101 name:"Aslam" email:"aslammulla.13@gmail.com" skills:"Go" skills:"Python" skills:"AWS" skills:"Docker"
*/

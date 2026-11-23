package main

import (
	"encoding/json"
	"fmt"
	"time"

	"buf.build/go/protovalidate"
	"github.com/SaveliiYam/otus-protobuf-advanced/api_models/gogo"
	protoc_gen_validate "github.com/SaveliiYam/otus-protobuf-advanced/api_models/protoc-gen-validate"
	protovalidatepb "github.com/SaveliiYam/otus-protobuf-advanced/api_models/protovalidate"
	v001 "github.com/SaveliiYam/otus-protobuf-advanced/api_models/v001"
	"github.com/go-playground/validator/v10"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	//classicProto()

	//protoGenValidate()

	//protoValidate()

	gogoValidate()
}

func gogoValidate() {
	updateUser := &gogo.UpdateUserRequest{
		Id:        0,
		Name:      "Jonh",
		Email:     "",
		CreatedAt: time.Now().Add(-time.Hour),
	}

	v := validator.New()

	err := v.Struct(updateUser)
	if err != nil {
		fmt.Println(err)
	}
}

func protoValidate() {
	userRequest := &protovalidatepb.CreateUserRequest2{
		Name:  "",
		Email: "test@mail.ru",
		Age:   0,
	}

	err := protovalidate.Validate(userRequest)
	if err != nil {
		fmt.Println(err)
	}
}

func protoGenValidate() {
	userRequest := &protoc_gen_validate.CreateUserRequest{
		Name:  "",
		Email: "test@mail.ru",
		Age:   0,
	}

	err := userRequest.Validate()
	if err != nil {
		fmt.Println(err)
	}
}

func classicProto() {
	user := &v001.User{
		Id:         1,
		Name:       "Jonh",
		LastName:   "Doe",
		MiddleName: nil,
		CreatedAt:  timestamppb.New(time.Now()),
	}

	fmt.Println(user)

	raw, _ := json.Marshal(user)

	fmt.Println(string(raw))

	raw, _ = proto.Marshal(user)

	fmt.Println(string(raw))
}

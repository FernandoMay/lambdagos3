package main

//go get github.com/aws/aws-lambda-go/lambda
//goos=linux go build -o main main.go
import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type InputEvent struct {
	Entidad  int    `json:"entidad"`
	Distrito int    `json:"distrito"`
	Username string `json:"username"`
	Paquete  string `json:"paquete"`
}

type Response struct {
	Entidad  int    `json:"entidad"`
	Distrito int    `json:"distrito"`
	Username string `json:"username"`
	Paquete  string `json:"paquete"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(event InputEvent) (Response, error) {
	fmt.Println(event)
	var sname string
	if event.Username == "" {
		sname = "std" + time.Now().String()

	} else {
		sname = event.Username
	}
	var spaquete string
	if event.Username == "" {
		spaquete = "pq:" + time.Now().String()

	} else {
		spaquete = event.Paquete
	}
	var nentidad int
	if event.Entidad < 0 || event.Entidad > 193 {
		nentidad = 444
	} else {
		nentidad = event.Entidad
	}
	var ndistrito int
	if event.Entidad < 0 || event.Distrito > 300 {
		ndistrito = 777
	} else {
		ndistrito = event.Entidad
	}

	return Response{
		Username: sname,
		Paquete:  spaquete,
		Entidad:  nentidad,
		Distrito: ndistrito,
	}, nil
}

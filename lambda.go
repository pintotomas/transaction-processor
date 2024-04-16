package main

import (
	"context"
	"fmt"
)

type Request struct {
	Email    string `json:"email"`
	FileName string `json:"file-name"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, request Request) (Response, error) {

	err := ProcessAndSendEmail(request.Email, request.FileName)
	if err != nil {
		return Response{Message: "Failed to send transaction summary details"}, err
	}

	message := fmt.Sprintf("Successfully sent transaction summary details to, %s!", request.Email)
	return Response{Message: message}, nil
}

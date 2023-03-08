package main

import (
    "context"
    "errors"

    "github.com/go-kit/kit/endpoint"
)

type UserService interface {
    Create(email string) (int, error)
}

type userService struct{}

func (userService) Create(email string) (int, error) {
    if email == "" {
        return 0, ErrEmpty
    }

    return 1, nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty string")

type createRequest struct {
    Email string `json:"email"`
}

type createResponse struct {
    ID  int    `json:"id"`
    Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func createEndpoint(user UserService) endpoint.Endpoint {
    return func(_ context.Context, request interface{}) (interface{}, error) {
        req := request.(createRequest)
        id, err := user.Create(req.Email)
        if err != nil {
            return createResponse{Err: err.Error()}, nil
        }
        return createResponse{ID: id}, nil
    }
}

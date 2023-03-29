package main

import (
    "context"
    "errors"

    "github.com/go-kit/kit/endpoint"

    "github.com/software-engr-full-stack/vaskafka"
)

type UserService interface {
    Create(email string) (int64, error)
}

type userService struct {
    db interface {
        Insert(string) (int64, error)
        // Get(int64) (*models.User, error)
    }
}

func (user userService) Create(email string) (int64, error) {
    if email == "" {
        return 0, ErrEmpty
    }
    id, err := user.db.Insert(email)
    if err != nil {
        return 0, err
    }

    err = vaskafka.Produce("user-created", email)
    if err != nil {
        return 0, err
    }

    return id, nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty string")

type createRequest struct {
    Email string `json:"email"`
}

type createResponse struct {
    ID  int64  `json:"id"`
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

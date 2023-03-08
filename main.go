package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"

    httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
    svc := userService{}

    createHandler := httptransport.NewServer(
        createEndpoint(svc),
        decodeCreateRequest,
        encodeResponse,
    )

    http.Handle("/user/create", createHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var request createRequest
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        return nil, err
    }
    return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}

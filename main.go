package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"

    "virtual-assistant/services/user/pkg/models"
    "virtual-assistant/services/user/pkg/models/postgres"

    httptransport "github.com/go-kit/kit/transport/http"
)

var db *sqlx.DB

const url = os.Getenv("VA_USERS_URL")

func main() {
    db = sqlx.Open("postgres", url)
    err := db.Ping()
    if err != nil {
        panic(err)
    }

    svc := userService{
        DB: db,
    }

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

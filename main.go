package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "strings"

    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/jmoiron/sqlx"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"

    "virtual-assistant/services/user/pkg/models/postgres"
)

func main() {
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    err := godotenv.Load()
    if err != nil {
        errorLog.Fatal("error loading .env file")
    }

    url := strings.TrimSpace(os.Getenv("DATABASE_URL"))
    if url == "" {
        errorLog.Fatal("DATABASE_URL env var is blank")
    }

    db, err := sqlx.Open("postgres", url)
    if err != nil {
        errorLog.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        errorLog.Fatal(err)
    }

    svc := userService{
        db: &postgres.UserModel{DB: db},
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

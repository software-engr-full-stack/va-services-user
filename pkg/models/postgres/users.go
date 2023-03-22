package postgres

type UserModel struct {
    ID    int64  `db:"id"`
    Email string `db:"email"`
}

func (user UserModel) Insert(db *sqlx.DB, email string) {
    stmt := "INSERT INTO users (email) VALUES ($1);"
    result, err := db.Exec(stmt, email)
}

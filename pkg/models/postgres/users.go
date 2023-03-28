package postgres

import (
    "github.com/jmoiron/sqlx"

    "virtual-assistant/services/user/pkg/models"
)

type UserModel struct {
    DB *sqlx.DB
}

const sqlInsert = "INSERT INTO users (email) VALUES (:email) RETURNING id"

func (user UserModel) Insert(email string) (int64, error) {
    tx := user.DB.MustBegin()

    query, args, err := tx.BindNamed(sqlInsert, &models.User{Email: email})
    if err != nil {
        rbErr := tx.Rollback()
        if err != nil {
            return 0, rbErr
        }
        return 0, err
    }

    var id struct {
        Val int64 `db:"id"`
    }

    err = sqlx.Get(tx, &id, query, args...)
    if err != nil {
        rbErr := tx.Rollback()
        if err != nil {
            return 0, rbErr
        }
        return 0, err
    }

    err = tx.Commit()
    if err != nil {
        rbErr := tx.Rollback()
        if err != nil {
            return 0, rbErr
        }
        return 0, err
    }

    return id.Val, nil
}

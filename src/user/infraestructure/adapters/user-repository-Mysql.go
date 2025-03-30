package adapters

import (
	"api-v1/src/database"
	"api-v1/src/user/domain/entities"
	"database/sql"
	"log"
)

type UserRepositoryMysql struct{
	DB *sql.DB
}

func NewUserRepositoryMysql() (*UserRepositoryMysql, error) {
	db, err := database.Connect()

	if err != nil {
		return nil, err
	}

	return &UserRepositoryMysql{DB: db}, nil
}

func (r *UserRepositoryMysql) Create(user entities.User)(entities.User, error){
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Fatal(err, 1)
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email, user.Password)
	
	if err != nil {
        return entities.User{}, err
    }

	id, err := result.LastInsertId()

	if err != nil {
        return entities.User{}, err
    }

	user.ID = int(id)
	user.Password = ""
	return user, nil
}


func (r *UserRepositoryMysql) GetByID(id int64)(entities.User, error){
	query := `SELECT id, name, email, password FROM users WHERE id =?`
    row := r.DB.QueryRow(query, id)

    var user entities.User

    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

    if err == sql.ErrNoRows {
        return entities.User{}, err
    } else if err != nil {
        return entities.User{}, err
    }

	user.Password = "" 

    return user, nil
}

func (r *UserRepositoryMysql) GetByEmail(email string)(entities.User, error){
	query := `SELECT id, name, email, password FROM users WHERE email =?`
	row := r.DB.QueryRow(query, email)

	var user entities.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return entities.User{}, err
	} else if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
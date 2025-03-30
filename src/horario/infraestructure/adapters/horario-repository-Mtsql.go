package adapters

import (
	"api-v1/src/database"
	"api-v1/src/horario/domain/entities"
	"database/sql"
)

type HorarioRepositoryMysql struct {
	DB *sql.DB
}

func NewHorarioRepositoryMysql() (*HorarioRepositoryMysql, error){
	db, err := database.Connect()

	if err != nil{
		return nil, err
	}

	return &HorarioRepositoryMysql{DB: db}, nil
}

func (r *HorarioRepositoryMysql) Create(horario entities.Horario) (entities.Horario, error){
	query := "INSERT INTO horarios (Minute, Hour) VALUES (?,?)"
	result, err := r.DB.Exec(query, horario.Minute, horario.Hour)

	if err != nil {
		return entities.Horario{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entities.Horario{}, err
	}

	horario.ID = int(id)

	return horario, nil
}

func (r *HorarioRepositoryMysql) GetAll() ([]entities.Horario, error){
	query := "SELECT ID, Minute, Hour FROM horarios"
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var horarios []entities.Horario
	for rows.Next() {
		var horario entities.Horario
		err := rows.Scan(&horario.ID, &horario.Minute, &horario.Hour)
		if err != nil {
			return nil, err
		}
		horarios = append(horarios, horario)
	}

	return horarios, nil
}

func (r *HorarioRepositoryMysql) Delete(id int64) (bool, error) {
	query := "DELETE FROM horarios WHERE ID = ?"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
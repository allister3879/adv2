package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Insert(moduleInfo ModuleInfo) error {
	stmt := `INSERT INTO module_info (module_name, module_duration, exam_type, version) VALUES ($1, $2, $3, $4)`
	_, err := m.DB.Exec(stmt, moduleInfo.ModuleName, moduleInfo.ModuleDuration, moduleInfo.ExamType, moduleInfo.Version)
	if err != nil {
		return fmt.Errorf("error inserting moduleInfo: %w", err)
	}
	return nil
}

func (m *DBModel) Retrieve(id int) (ModuleInfo, error) {
	var moduleInfo ModuleInfo
	stmt := `SELECT id, created_at, updated_at, module_name, module_duration, exam_type, version FROM module_info WHERE id = $1`
	row := m.DB.QueryRow(stmt, id)
	err := row.Scan(&moduleInfo.ID, &moduleInfo.CreatedAt, &moduleInfo.UpdatedAt, &moduleInfo.ModuleName, &moduleInfo.ModuleDuration, &moduleInfo.ExamType, &moduleInfo.Version)
	if err != nil {
		return ModuleInfo{}, fmt.Errorf("error retrieving moduleInfo: %w", err)
	}
	return moduleInfo, nil
}

func (m *DBModel) Update(moduleInfo ModuleInfo) error {
	stmt := `UPDATE module_info SET module_name=$1, module_duration=$2, exam_type=$3, version=$4, updated_at=now() WHERE id=$5`
	_, err := m.DB.Exec(stmt, moduleInfo.ModuleName, moduleInfo.ModuleDuration, moduleInfo.ExamType, moduleInfo.Version, moduleInfo.ID)
	if err != nil {
		return fmt.Errorf("error updating moduleInfo: %w", err)
	}
	return nil
}

func (m *DBModel) Delete(id int) error {
	stmt := `DELETE FROM module_info WHERE id = $1`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return fmt.Errorf("error deleting moduleInfo: %w", err)
	}
	return nil
}

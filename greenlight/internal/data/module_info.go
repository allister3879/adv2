package data

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type ModuleInfo struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	ModuleName     string    `json:"module_name"`
	ModuleDuration int64     `json:"module_duration"`
	ExamType       string    `json:"exam_type"`
	Version        int       `json:"version"`
}

type ModuleModel struct {
	DB *sql.DB
}

func (m ModuleModel) Insert(moduleInfo *ModuleInfo) error {
	query := `INSERT INTO module_info (module_name, module_duration, exam_type, version) 
	          VALUES ($1, $2, $3, $4)
			  RETURNING id, created_at`
	return m.DB.QueryRow(query, &moduleInfo.ModuleName, &moduleInfo.ModuleDuration, &moduleInfo.ExamType, &moduleInfo.Version).Scan(&moduleInfo.ID, &moduleInfo.CreatedAt)
}

func (m ModuleModel) Get(id int64) (*ModuleInfo, error) {
	query := `SELECT * FROM module_info WHERE id=$1`
	var moduleInfo ModuleInfo

	err := m.DB.QueryRow(query, id).Scan(
		&moduleInfo.ID,
		&moduleInfo.CreatedAt,
		&moduleInfo.UpdatedAt,
		&moduleInfo.ModuleName,
		&moduleInfo.ModuleDuration,
		&moduleInfo.ExamType,
		&moduleInfo.Version,
	)

	if err != nil {
		return nil, err
	}

	return &moduleInfo, nil
}

func (m ModuleModel) Update(moduleInfo *ModuleInfo) error {
	query := `UPDATE module_info 
	          SET module_name=$1, module_duration=$2, exam_type=$3, version=$4, updated_at=now() 
			  WHERE id=$5
			  RETURNING version`
	args := []interface{}{
		moduleInfo.ModuleName,
		moduleInfo.CreatedAt,
		moduleInfo.ModuleDuration,
		moduleInfo.ExamType,
	}

	return m.DB.QueryRow(query, args...).Scan(&moduleInfo.Version)
}

func (m ModuleModel) Delete(id int64) error {
	query := `DELETE FROM module_info WHERE id=$1`
	_, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

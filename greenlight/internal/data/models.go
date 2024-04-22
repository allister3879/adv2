package data

import (
	"database/sql"
)

type DBModels struct {
	Modules ModuleModel
	Module2 DepartmentModel
	Module3 UserModel
}

func NewModels(db *sql.DB) DBModels {
	return DBModels{
		Modules: ModuleModel{DB: db},
		Module2: DepartmentModel{DB: db},
		Module3: UserModel{DB: db},
	}
}

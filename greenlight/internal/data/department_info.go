package data

import "database/sql"

type DepartmentInfo struct {
	ID                 int64  `json:"id"`
	DepartmentName     string `json:"department_name"`
	StaffQuantity      int    `json:"staff_quantity"`
	DepartmentDirector string `json:"department_director"`
	ModuleID           int    `json:"module_id"`
}

type DepartmentModel struct {
	DB *sql.DB
}

func (m DepartmentModel) Insert(depInfo *DepartmentInfo) error {
	query := `INSERT INTO department_info (department_name, staff_quantity, department_director, module_id) 
	          VALUES ($1, $2, $3, $4)
			  RETURNING department_name, staff_quantity, department_director, module_id`
	return m.DB.QueryRow(query, depInfo.DepartmentName, depInfo.StaffQuantity, depInfo.DepartmentDirector, depInfo.ModuleID).Scan(&depInfo.DepartmentName, &depInfo.StaffQuantity, &depInfo.DepartmentDirector, &depInfo.ModuleID)
}

func (m DepartmentModel) Get(id int64) (*DepartmentInfo, error) {
	query := `SELECT * FROM department_info WHERE id=$1`
	var depInfo DepartmentInfo

	err := m.DB.QueryRow(query, id).Scan(
		&depInfo.ID,
		&depInfo.DepartmentName,
		&depInfo.DepartmentDirector,
		&depInfo.StaffQuantity,
		&depInfo.ModuleID,
	)

	if err != nil {
		return nil, err
	}

	return &depInfo, nil
}

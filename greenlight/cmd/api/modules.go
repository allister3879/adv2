package main

import (
	"fmt"
	"net/http"

	"greenlight.d.net/internal/data"
)

func (app *application) createModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ModuleName     string `json:"module_name"`
		ModuleDuration int64  `json:"module_duration"`
		ExamType       string `json:"exam_type"`
		Version        int    `json:"version"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	moduleINfo := &data.ModuleInfo{
		ModuleName:     input.ModuleName,
		ModuleDuration: input.ModuleDuration,
		ExamType:       input.ExamType,
		Version:        input.Version,
	}

	err = app.models.Modules.Insert(moduleINfo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("v1/module-info/%d", moduleINfo.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"moduleInfo": moduleINfo}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	module, err := app.models.Modules.Get(id)

	err = app.writeJSON(w, http.StatusOK, envelope{"module": module}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Modules.Delete(id)

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) editModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	module, err := app.models.Modules.Get(id)

	var input struct {
		ModuleName     string `json:"module_name"`
		ModuleDuration int64  `json:"module_duration"`
		ExamType       string `json:"exam_type"`
		Version        int    `json:"version"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	module.ModuleName = input.ModuleName
	module.ModuleDuration = input.ModuleDuration
	module.ExamType = input.ExamType
	module.Version = input.Version

	err = app.models.Modules.Update(module)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"module": module}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) createDepInfoHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		DepartmentName     string `json:"department_name"`
		StaffQuantity      int    `json:"staff_quantity"`
		DepartmentDirector string `json:"department_director"`
		ModuleID           int    `json:"module_id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	depInfo := &data.DepartmentInfo{
		DepartmentName:     input.DepartmentName,
		StaffQuantity:      input.StaffQuantity,
		DepartmentDirector: input.DepartmentDirector,
		ModuleID:           input.ModuleID,
	}

	err = app.models.Module2.Insert(depInfo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("v1/department-info/%d", depInfo.ID))
	err = app.writeJSON(w, http.StatusCreated, envelope{"moduleInfo": depInfo}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getDepInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	depInfo, err := app.models.Module2.Get(id)

	err = app.writeJSON(w, http.StatusOK, envelope{"depInfo": depInfo}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

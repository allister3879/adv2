package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func createModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	var moduleInfo ModuleInfo
	err := json.NewDecoder(r.Body).Decode(&moduleInfo)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	id, err := dbModel.Insert("INSERT INTO module_info (created_at, updated_at, module_name, module_duration, exam_type, version) VALUES (NOW(), NOW(), $1, $2, $3, $4) RETURNING id",
		moduleInfo.ModuleName, moduleInfo.ModuleDuration, moduleInfo.ExamType, moduleInfo.Version)
	if err != nil {
		http.Error(w, "Failed to insert moduleInfo into database", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, map[string]int{"id": id}, http.StatusCreated)
}

func getModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	row := dbModel.Retrieve("SELECT id, created_at, updated_at, module_name, module_duration, exam_type, version FROM module_info WHERE id = $1", id)
}

func editModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updatedModuleInfo ModuleInfo
	err := json.NewDecoder(r.Body).Decode(&updatedModuleInfo)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	_, err = dbModel.Update("UPDATE module_info SET module_name = $1, module_duration = $2, exam_type = $3, version = $4 WHERE id = $5",
		updatedModuleInfo.ModuleName, updatedModuleInfo.ModuleDuration, updatedModuleInfo.ExamType, updatedModuleInfo.Version, id)
	if err != nil {
		http.Error(w, "Failed to update moduleInfo in database", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, map[string]string{"message": "ModuleInfo updated successfully"}, http.StatusOK)
}

func deleteModuleInfoHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := dbModel.Delete("DELETE FROM module_info WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Failed to delete moduleInfo from database", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, map[string]string{"message": "ModuleInfo deleted successfully"}, http.StatusOK)
}

func routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/module-info", createModuleInfoHandler)

	r.Get("/module-info/{id}", getModuleInfoHandler)

	r.Put("/module-info/{id}", editModuleInfoHandler)
	r.Patch("/module-info/{id}", editModuleInfoHandler)

	r.Delete("/module-info/{id}", deleteModuleInfoHandler)

	return r
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

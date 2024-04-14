package main

import "time"

type ModuleInfo struct {
	ID             int       `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	ModuleName     string    `json:"moduleName"`
	ModuleDuration int       `json:"moduleDuration"`
	ExamType       string    `json:"examType"`
	Version        string    `json:"version"`
}

package model

import "gorm.io/gorm"

type UnrealRunner struct {
	gorm.Model
	Name     string
	ExecPath string
	Params   StringSlice
	Pid      int
	Status   int
}

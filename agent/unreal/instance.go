package unreal

import "gorm.io/gorm"

type Instance struct {
	gorm.Model
	Name     string
	ExecPath string
	Params   []string
	Pid      int
	Status   int
}

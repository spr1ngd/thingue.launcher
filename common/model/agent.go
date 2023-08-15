package model

type Agent struct {
	Instances []Instance `gorm:"foreignKey:InstanceID"`
}

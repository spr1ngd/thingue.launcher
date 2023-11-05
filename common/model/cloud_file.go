package model

type CloudFile struct {
	Id       uint   `json:"id" gorm:"primarykey"`
	FileName string `json:"fileName"`
	Hash     string `json:"hash"`
	Res      string `json:"res"`
}

package models

type FileModel struct {
	Filename string `json:"filename"`
	File     byte   `json:"file"`
}

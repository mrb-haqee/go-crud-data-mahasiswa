package model

import "gorm.io/gorm"

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type DataMahasiswa struct {
	gorm.Model
	Mahasiswa Mahasiswa `gorm:"embedded"`
}

type Mahasiswa struct {
	Nama         string 
	NIM          string `gorm:"unique"`
	Gender       string `gorm:"type:varchar(2)"`
	TempatLahir  string
	TanggalLahir string
	Prodi        string
	NoHP         string
	Alamat       string
}

type Message struct {
	Validation any
	Pesan      string
}

type Send struct {
	Nama         string `validate:"required" label:"Nama Lengkap"`
	NIM          string `validate:"required"` 
	Gender       string `validate:"required"` 
	TempatLahir  string `validate:"required" label:"Tempat Lahir"`
	TanggalLahir string `validate:"required" label:"Tanggal lahir"`
	Prodi        string `validate:"required"`
	NoHP         string `validate:"required" label:"Nomer HP"`
	Alamat       string `validate:"required"`
}

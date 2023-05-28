package main

import (

	"net/http"

	"github.com/mrb-haqee/go-crud-data-mahasiswa/controller"
	"log"

	// "github.com/mrb-haqee/go-crud-data-mahasiswa/model"
	// "github.com/mrb-haqee/go-crud-data-mahasiswa/service"
)

func main() {

	// add := service.NewDB()
	// err:=add.Update(1,model.Mahasiswa{
	// 	Nama:         "9kkknd",
	// 	NIM:          "90gcdrd",
	// 	Gender:       "ll",
	// 	TempatLahir:  "negare",
	// 	TanggalLahir: "20-01-2000",
	// 	Prodi:        "matematika",
	// 	NoHP:         "09891232",
	// 	Alamat:       "adwadwada",
	// })
	// if err == nil {
	// 	log.Print("berhasil di ubah")
	// }

	http.HandleFunc("/", controller.HomePage)
	http.HandleFunc("/mahasiswa", controller.HomePage)
	http.HandleFunc("/mahasiswa/home", controller.HomePage)
	http.HandleFunc("/mahasiswa/add", controller.Add)
	http.HandleFunc("/mahasiswa/update", controller.Update)
	http.HandleFunc("/mahasiswa/delete", controller.Delete)

	log.Println("Server Berjalan pada: http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

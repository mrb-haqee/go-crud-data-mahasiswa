package controller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/mrb-haqee/go-crud-data-mahasiswa/lib"
	"github.com/mrb-haqee/go-crud-data-mahasiswa/model"
	"github.com/mrb-haqee/go-crud-data-mahasiswa/service"
)

var DBMahasiswa = service.NewDB()
var validasi = lib.NewVaditaion()

func HomePage(w http.ResponseWriter, _ *http.Request) {

	var GetData = service.NewDB()
	data, _ := GetData.FindAll()

	for i, dm := range data {
		if dm.Mahasiswa.Gender == "lk" {
			data[i].Mahasiswa.Gender = "Laki-laki"
		} else if dm.Mahasiswa.Gender == "pr" {
			data[i].Mahasiswa.Gender = "Perempuan"
		}
		tgl, _ := time.Parse("2006-01-02", dm.Mahasiswa.TanggalLahir)
		data[i].Mahasiswa.TanggalLahir = tgl.Format("02-01-2006")
	}

	impo := map[string]interface{}{"data": data}

	temp, _ := template.ParseFiles("view/home.html")
	temp.Execute(w, impo)
}

func Add(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("view/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		var mahasiswa model.Send
		mahasiswa.Nama = r.FormValue("name")
		mahasiswa.NIM = r.FormValue("nim")
		mahasiswa.Gender = r.FormValue("gender")
		mahasiswa.TempatLahir = r.FormValue("tampat_lahir")
		mahasiswa.TanggalLahir = r.FormValue("tanggal_lahir")
		mahasiswa.Prodi = r.FormValue("prodi")
		mahasiswa.NoHP = r.FormValue("no_hp")
		mahasiswa.Alamat = r.FormValue("alamat")

		mess := make(map[string]any)

		vErr := validasi.Struct(mahasiswa)
		if vErr != nil {
			mess["validation"] = vErr
		} else {
			DBMahasiswa.Add(mahasiswa)
			mess["pesan"] = "Data Mahasiswa berhasi di input"
		}

		temp, _ := template.ParseFiles("view/add.html")
		temp.Execute(w, mess)
	}

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		getId := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(getId)

		var mahasiswa model.DataMahasiswa
		DBMahasiswa.FindId(id, &mahasiswa)

		impo := map[string]interface{}{"data": mahasiswa}

		temp, _ := template.ParseFiles("view/update.html")
		temp.Execute(w, impo)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		getId := r.FormValue("id")
		id, _ := strconv.Atoi(getId)

		var mahasiswa model.Mahasiswa
		mahasiswa.Nama = r.FormValue("name")
		mahasiswa.NIM = r.FormValue("nim")
		mahasiswa.Gender = r.FormValue("gender")
		mahasiswa.TempatLahir = r.FormValue("tampat_lahir")
		mahasiswa.TanggalLahir = r.FormValue("tanggal_lahir")
		mahasiswa.Prodi = r.FormValue("prodi")
		mahasiswa.NoHP = r.FormValue("no_hp")
		mahasiswa.Alamat = r.FormValue("alamat")

		DBMahasiswa.Update(id, mahasiswa)
		var mess model.Message
		mess.Pesan = "Data Mahasiswa berhasi di Update"

		temp, _ := template.ParseFiles("view/update.html")
		temp.Execute(w, mess)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	getId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(getId)
	DBMahasiswa.Delete(id)

	http.Redirect(w, r, "/mahasiswa", http.StatusSeeOther)
}

package data

import (
	"github.com/irvandandung/goAPI/config"
	"github.com/irvandandung/goAPI/pkg/data/local"
	"log"
	"strconv"
)

func InsertDataUser(data map[string]string) (string, error){
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryInsert(db, "user", data)
	return response, err
}

func InsertDataBuku(data map[string]string) (string, error){
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryInsert(db, "buku", data)
	return response, err
}

func UpdateDataUser(data map[string]string, wheredata map[string]string) (string, error){
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryUpdate(db, "user", data, wheredata)
	return response, err
}

func UpdateDataBuku(data map[string]string, wheredata map[string]string) (string, error){
	db := config.ConnectDB()
	defer db.Close()
	response, err := local.QueryUpdate(db, "buku", data, wheredata)
	return response, err
}

func GetAllDataUsers() ([]Users){
	var user Users
	var list_user []Users
	wheredata := map[string]string{}
	db := config.ConnectDB()
	defer db.Close()
	var fields = []string{ "username", "password" }
	rows := local.QuerySelect(db, "user", fields, wheredata)
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password)
		if(err != nil){
			log.Fatal(err.Error())
		}
		list_user = append(list_user, user)
	}

	return list_user
}

func GetDataUser(username string, password string) (bool, Users){
	var user Users
	data := false
	wheredata := map[string]string{"username =":"'"+username+"' AND", "password =":"'"+config.GetMD5Hash(password)+"'"}
	db := config.ConnectDB()
	defer db.Close()
	var fields = []string{"id", "username", "password"}
	rows := local.QuerySelect(db, "user", fields, wheredata)
	for rows.Next(){
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if(err != nil){
			log.Fatal(err.Error())
		}
	}
	if(user.Username != ""){
		data = true
	}

	return data, user
}

func GetAllDataBuku() ([]Buku){
	var buku Buku
	var list_buku []Buku
	wheredata := map[string]string{}
	db := config.ConnectDB()
	defer db.Close()
	var fields = []string{ "id", "judul", "keterangan", "pencipta", "tahun" }
	rows := local.QuerySelect(db, "buku", fields, wheredata)
	for rows.Next() {
		err := rows.Scan(&buku.Id, &buku.Judul, &buku.Keterangan, &buku.Pencipta, &buku.Tahun)
		if(err != nil){
			log.Fatal(err.Error())
		}
		list_buku = append(list_buku, buku)
	}

	return list_buku
}

func GetDataBukuById(id int) (Buku) {
	var buku Buku
	idString := strconv.Itoa(id)
	wheredata := map[string]string{ "id =":idString}
	db := config.ConnectDB()
	defer db.Close()
	var fields = []string{ "id", "judul", "keterangan", "pencipta", "tahun" }
	rows := local.QuerySelect(db, "buku", fields, wheredata)
	for rows.Next() {
		err := rows.Scan(&buku.Id, &buku.Judul, &buku.Keterangan, &buku.Pencipta, &buku.Tahun)
		if(err != nil){
			log.Fatal(err.Error())
		}
	}

	return buku
}
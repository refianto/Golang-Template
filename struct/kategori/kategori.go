package dbstruct

import "gopkg.in/mgo.v2/bson"

// GetKategori is base data get
type GetKategori struct {
	ID      bson.ObjectId `json:"id" bson:"id"`
	Nama    string        `json:"nama" bson:"nama"`
	Foto    string        `json:"foto" bson:"foto"`
	IsAktif int           `json:"isAktif" bson:"isAktif"`
}

// PostKategori is general insert data setup
type PostKategori struct {
	Nama    string `json:"nama" bson:"nama" binding:"required"`
	Foto    string `json:"foto" bson:"foto" binding:"required"`
	IsAktif int    `json:"isAktif" bson:"isAktif" default:"1" default1:"1"`
}

// PutKategori is general update data setup
type PutKategori struct {
	Nama    string `json:"nama,omitempty"  bson:"nama,omitempty" `
	Foto    string `json:"foto,omitempty"  bson:"foto,omitempty"`
	IsAktif int    `json:"isAktif,omitempty"  bson:"isAktif,omitempty"`
}

// IDKategori is general id data
type IDKategori struct {
	ID bson.ObjectId `json:"_id" bson:"_id"`
}

// GetDataStruct is general struct
// type GetDataStruct []GetKategori

// SetDataStruct is general struct
// type SetDataStruct SetKategori

// // GetData for get all data
// func GetData(dataColllection string) (Data, error) {
// 	db := conn.GetDB()
// 	data := Data{}
// 	err := db.C(dataColllection).Find(bson.M{}).All(&data)
// 	return data, err
// }

// // GetDataOne for get one data by id
// func GetDataOne(id bson.ObjectId, dataColllection string) (Data, error) {
// 	db := conn.GetDB()
// 	data := Data{}
// 	err := db.C(dataColllection).Find(bson.M{"_id": &id}).One(&data)
// 	return data, err
// }

// // GerDataPagination for get data limit 10
// func GetRowCount(dataColllection string, page int) (Data, error) {
// 	db := conn.GetDB()
// 	data := Data{}
// 	err := db.C(dataColllection).ount(bson.M{}).All(&data)
// 	return data, err
// }

// // InsertData for insert general data
// func InsertData() {

// }

// // UpdateData for gneral update data
// func UpdateData() {

// }

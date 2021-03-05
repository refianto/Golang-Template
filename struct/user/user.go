package user

import "gopkg.in/mgo.v2/bson"

// GetStruct is base data get
type GetStruct struct {
	ID       bson.ObjectId `json:"id" bson:"id"`
	Nama     string        `json:"nama" bson:"nama"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
	Alamat   string        `json:"alamat" bson:"alamat"`
	Phone    string        `json:"phone" bson:"phone"`
	Role     []int         `json:"role" bson:"role"`
	IsAktif  int           `json:"isAktif" bson:"isAktif"`
}

// PostStruct is general insert data setup
type PostStruct struct {
	Nama     string `json:"nama" bson:"nama" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Alamat   string `json:"alamat,omitempty" bson:"alamat,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Role     []int  `json:"role" bson:"role" binding:"required"`
	IsAktif  int    `json:"isAktif" bson:"isAktif"`
}

// PutStruct is general update data setup
type PutStruct struct {
	Nama     string `json:"nama,omitempty" bson:"nama,omitempty" `
	Email    string `json:"email,omitempty" bson:"email,omitempty" valiedate:"email"`
	Password string `json:"password,omitempty" bson:"password,omitempty" `
	Alamat   string `json:"alamat,omitempty" bson:"alamat,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	Role     []int  `json:"role,omitempty" bson:"role,omitempty" `
	IsAktif  int    `json:"isAktif,omitempty" bson:"isAktif,omitempty"`
}

// IDStruct is general id data
type IDStruct struct {
	ID bson.ObjectId `json:"_id" bson:"_id"`
}

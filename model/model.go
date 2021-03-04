package model

import (
	"strconv"

	"github.com/bisaai/digileaps/conn"

	"gopkg.in/mgo.v2/bson"
)

// GetData for get all data
func GetData(SortBy, Page string, SelectField, param bson.M, Data interface{}, dataColllection string) (int, interface{}, error) {
	db := conn.GetDB()
	var data []interface{}
	data = append(data, Data)
	page, _ := strconv.Atoi(Page)
	if page <= 0 {
		page = 1
	}
	sortBy := SortBy
	if sortBy == "" {
		sortBy = "-_id"
	}
	err := db.C(dataColllection).Find(param).Select(SelectField).Sort(sortBy).Skip((page - 1) * 10).Limit(10).All(&data)
	countData, _ := db.C(dataColllection).Find(param).Count()
	return countData, &data, err
}

func GetDataLimit(Lim int, Page string, param bson.M, Data interface{}, dataColllection string) (int, interface{}, error) {
	db := conn.GetDB()
	var data []interface{}
	data = append(data, Data)
	page, _ := strconv.Atoi(Page)
	if page <= 0 {
		page = 1
	}
	err := db.C(dataColllection).Find(param).Skip((page - 1) * 10).Limit(Lim).All(&data)
	countData, _ := db.C(dataColllection).Find(param).Count()
	return countData, &data, err
}

// InsertData for insert general data
func InsertData(Data interface{}, dataCollection string) (interface{}, error) {
	db := conn.GetDB()
	data := Data
	err := db.C(dataCollection).Insert(&data)
	// Set(data, "default0")
	return data, err
}

// InsertDataMany for insert general data
func InsertDataMany(Data interface{}, dataCollection string) (interface{}, error) {
	db := conn.GetDB()
	var data []interface{}
	data = append(data, Data)
	err := db.C(dataCollection).Insert(&data)
	return data, err
}

// UpdateData for gneral update data
func UpdateData(SelectData, Data interface{}, dataCollection string) (interface{}, error) {
	db := conn.GetDB()
	data := Data
	err := db.C(dataCollection).Update(&SelectData, bson.M{"$set": &data})
	return data, err
}

// // SaveFile is general save file
// func SaveFile(Location, fileName, file string) (string, error) {
// 	// return
// }

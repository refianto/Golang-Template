package kategori

import (
	"log"

	"github.com/bisaai/digileaps/conn"
	model "github.com/bisaai/digileaps/model"
	dbstruct "github.com/bisaai/digileaps/struct/kategori"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

var (
	errNotExist        = errors.New("Data are not exist")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidParam    = errors.New("Invalid parameter")
	errInvalidInt      = errors.New("Invalid integer")
	errInvalidBody     = errors.New("Invalid request body")
	errInsertionFailed = errors.New("Error in the data insertion")
	errUpdationFailed  = errors.New("Error in the data updation")
)

const dataCollection = "kategori"

var db = conn.GetDB()

// FilterParamenter to filter param get
type FilterParamenter struct {
	ID     bson.ObjectId `form:"id"`
	Page   string        `form:"page"`
	SortBy string        `form:"sortBy"`
}

// GetKategori to get data
func GetKategori(c *gin.Context) {

	structData := make([]dbstruct.GetKategori, 0)
	quary := bson.M{}
	var ft FilterParamenter

	err := c.ShouldBind(&ft)
	if err != nil {
		log.Println(errInvalidParam.Error())
	}

	/*
		- page -> pagination default 1
		- sortBY -> sort data berdasarkan filed (- untuk desc) ex sortBy : -name
		- selectField -> memilih field yang akan di tampilkan sesuai denga struct
	*/
	page := ft.Page
	sortBy := ft.SortBy
	selectField := bson.M{"nama": 1}

	rowCount, data, err := model.GetData(sortBy, page, selectField, quary, structData, dataCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"offset": 10, "page": 1, "row_count": rowCount, "data": &data})
}

// AddKategori to add data
func AddKategori(c *gin.Context) {
	structData := dbstruct.PostKategori{}
	err := c.BindJSON(&structData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}
	if structData.IsAktif == 0 {
		structData.IsAktif = 1
	}
	data, err := model.InsertData(structData, dataCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": &data})

}

// UpdateKategori to update data
func UpdateKategori(c *gin.Context) {
	structData := dbstruct.PutKategori{}
	err := c.BindJSON(&structData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	selectData := dbstruct.IDKategori{}
	selectData.ID = bson.ObjectIdHex(c.Query("id"))

	data, err := model.UpdateData(selectData, structData, dataCollection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": &data})
}

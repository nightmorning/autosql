package sql

import (
	"autosql/common"
	"log"
	"spider-book/db"
	"strings"
)

type Book struct {
	BookId int64 `gorm:"primary_key;AUTO_INCREMENT" json:"book_id"`
	MemberId int64 `json:"member_id"`
	CateId int `json:"cate_id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Image string `json:"image"`
	IsEnd int `json:"is_end"`
	Words int `json:"words"`
	Views int `json:"views"`
	Recommend int `json:"recommend"`
	Preface string `json:"preface"`
	Tag string `json:"tag"`
	IsDel int `json:"is_del"`
	IsVip int `json:"is_vip"`
	Sort int `json:"sort"`
	CreatedAt int64 `json:"created_at"`
}

func GetBookById(id int) (Book, bool) {
	var book Book
	ok := db.SqlDB.Where("book_id = ?", id).First(&book).RecordNotFound()

	return book, ok
}

func GetBookByOne(condition []string) (Book, bool) {
	var book Book
	ok := db.SqlDB.Where(condition).First(&book).RecordNotFound()

	return book, ok
}

func GetBookList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	books := make([]*Book, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.SqlDB.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&books).
		Error

	return books, err
}

func GetBookCount(condition []string) int {
	var book Book
	count := 0
	db.SqlDB.Where(condition).Find(&book).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetBookList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetBookCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteBookByIds(ids string) (error) {
	var book Book
	err := db.SqlDB.Where("id IN  ?", ids).Delete(&book).Error

	return err
}

func DeleteBookById(id int64) error {
	var book Book
	book.BookId = id
	err := db.SqlDB.Delete(&book).Error

	return err
}

func CreateBook(book Book) (Book, error) {
	err := db.SqlDB.Create(&book).Error

	return book, err
}

func UpdateBook(condition []string, params map[string]interface{}) error {
	var book Book
	err := db.SqlDB.Model(&book).Where(condition).Updates(params).Error

	return err
}

func SaveBook(book Book) error {
	err := db.SqlDB.Save(&book).Error

	return err
}

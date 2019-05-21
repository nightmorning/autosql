package model

import (
	"autosql/common"
	"strings"
)

type Article struct {
 ArticleId int `gorm:"primary_key;AUTO_INCREMENT" json:" article_id"`
CateId int `json:"cate_id"`
Title int `json:"title"`
Key int `json:"key"`
Content int `json:"content"`
Author int `json:"author"`
Source int `json:"source"`
Abstract int `json:"abstract"`
View int `json:"view"`
Comment int `json:"comment"`
Status int `json:"status"`
Top int `json:"top"`
Recommend int `json:"recommend"`
CreatedAt int `json:"created_at"`
UpdatedAt int `json:"updated_at"`
}
func GetArticleById(id int) (Article, bool) {
	var article Article
	ok := db.Where(" article_id = ?", id).First(&article).RecordNotFound()

	return article, ok
}

func GetArticleByOne(condition []string) (Article, bool) {
	var article Article
	ok := db.Where(condition).First(&article).RecordNotFound()

	return article, ok
}

func GetArticleList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	articles := make([]*Article, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&articles).
		Error

	return articles, err
}

func GetArticleCount(condition []string) int {
	var article Article
	count := 0
	db.Where(condition).Find(&article).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetArticleList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetArticleCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteArticleByIds(ids string) (error) {
	var article Article
	err := db.Where("id IN  ?", ids).Delete(&article).Error

	return err
}

func DeleteArticleById(id int64) error {
	var article Article
	article.ArticleId = id
	err := db.Delete(&article).Error

	return err
}

func CreateArticle(article Article) (Article, error) {
	err := db.Create(&article).Error

	return article, err
}

func UpdateArticle(condition []string, params map[string]interface{}) error {
	var article Article
	err := db.Model(&article).Where(condition).Updates(params).Error

	return err
}

func SaveArticle(article Article) error {
	err := db.Save(&article).Error

	return err
}
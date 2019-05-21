package model

type ArticleComment struct {
CommentId int `gorm:"primary_key;AUTO_INCREMENT" json:"comment_id"`
ForCommentId int `json:"for_comment_id"`
TownId int `json:"town_id"`
MemberId int `json:"member_id"`
ArticleId int `json:"article_id"`
Image int `json:"image"`
Content int `json:"content"`
IsDel int `json:"is_del"`
CreatedAt int `json:"created_at"`
}
func GetArticleCommentById(id int) (ArticleComment, bool) {
	var article_comment ArticleComment
	ok := db.First(&article_comment, id).RecordNotFound()

	return article_comment, ok
}

func GetArticleCommentByOne(condition map[string]interface{}) (ArticleComment, bool) {
	var article_comment ArticleComment
	ok := db.Where(condition).First(&article_comment).RecordNotFound()

	return article_comment, ok
}

func GetArticleCommentList(condition map[string]interface{}, page int, pageSize int, orderBy []string) (interface{}, error) {
	article_comments := make([]*ArticleComment, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&article_comments).
		Error

	return article_comments, err
}

func GetArticleCommentCount(condition map[string]interface{}) int {
	var article_comment ArticleComment
	count := 0
	db.Where(condition).Find(&article_comment).Count(&count)

	return count
}

func GetPageUtil(condition map[string]interface{}, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetArticleCommentList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetArticleCommentCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteArticleCommentByIds(ids string) (error) {
	var article_comment ArticleComment
	err := db.Where("id IN  ?", ids).Delete(&article_comment).Error

	return err
}

func DeleteArticleCommentById(id int64) error {
	var article_comment ArticleComment
	article_comment.ArticleCommentId = id
	err := db.Delete(&article_comment).Error

	return err
}

func CreateArticleComment(article_comment ArticleComment) (ArticleComment, error) {
	err := db.Create(&article_comment).Error

	return article_comment, err
}

func UpdateArticleComment(condition map[string]interface{}, params map[string]interface{}) error {
	var article_comment ArticleComment
	err := db.Model(&article_comment).Where(condition).Updates(params).Error

	return err
}

func SaveArticleComment(article_comment ArticleComment) error {
	err := db.Save(&article_comment).Error

	return err
}
package model

type TopicPost struct {
PostId int `gorm:"primary_key;AUTO_INCREMENT" json:"post_id"`
TopicId int `json:"topic_id"`
MemberId int `json:"member_id"`
Content int `json:"content"`
ReplyId int `json:"reply_id"`
IsDel int `json:"is_del"`
CreatedAt int `json:"created_at"`
}
func GetTopicPostById(id int) (TopicPost, bool) {
	var topic_post TopicPost
	ok := db.First(&topic_post, id).RecordNotFound()

	return topic_post, ok
}

func GetTopicPostByOne(condition map[string]interface{}) (TopicPost, bool) {
	var topic_post TopicPost
	ok := db.Where(condition).First(&topic_post).RecordNotFound()

	return topic_post, ok
}

func GetTopicPostList(condition map[string]interface{}, page int, pageSize int, orderBy []string) (interface{}, error) {
	topic_posts := make([]*TopicPost, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&topic_posts).
		Error

	return topic_posts, err
}

func GetTopicPostCount(condition map[string]interface{}) int {
	var topic_post TopicPost
	count := 0
	db.Where(condition).Find(&topic_post).Count(&count)

	return count
}

func GetPageUtil(condition map[string]interface{}, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetTopicPostList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetTopicPostCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteTopicPostByIds(ids string) (error) {
	var topic_post TopicPost
	err := db.Where("id IN  ?", ids).Delete(&topic_post).Error

	return err
}

func DeleteTopicPostById(id int64) error {
	var topic_post TopicPost
	topic_post.TopicPostId = id
	err := db.Delete(&topic_post).Error

	return err
}

func CreateTopicPost(topic_post TopicPost) (TopicPost, error) {
	err := db.Create(&topic_post).Error

	return topic_post, err
}

func UpdateTopicPost(condition map[string]interface{}, params map[string]interface{}) error {
	var topic_post TopicPost
	err := db.Model(&topic_post).Where(condition).Updates(params).Error

	return err
}

func SaveTopicPost(topic_post TopicPost) error {
	err := db.Save(&topic_post).Error

	return err
}
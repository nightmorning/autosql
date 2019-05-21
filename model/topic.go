package model

type Topic struct {
TopicId int `gorm:"primary_key;AUTO_INCREMENT" json:"topic_id"`
MemberId int `json:"member_id"`
Content int `json:"content"`
Images int `json:"images"`
Type int `json:"type"`
Comment int `json:"comment"`
IsDel int `json:"is_del"`
CreatedAt int `json:"created_at"`
}
func GetTopicById(id int) (Topic, bool) {
	var topic Topic
	ok := db.Where("topic_id = ?", id).First(&topic).RecordNotFound()

	return topic, ok
}

func GetTopicByOne(condition []string) (Topic, bool) {
	var topic Topic
	ok := db.Where(condition).First(&topic).RecordNotFound()

	return topic, ok
}

func GetTopicList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	topics := make([]*Topic, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&topics).
		Error

	return topics, err
}

func GetTopicCount(condition []string) int {
	var topic Topic
	count := 0
	db.Where(condition).Find(&topic).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetTopicList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetTopicCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteTopicByIds(ids string) (error) {
	var topic Topic
	err := db.Where("id IN  ?", ids).Delete(&topic).Error

	return err
}

func DeleteTopicById(id int64) error {
	var topic Topic
	topic.TopicId = id
	err := db.Delete(&topic).Error

	return err
}

func CreateTopic(topic Topic) (Topic, error) {
	err := db.Create(&topic).Error

	return topic, err
}

func UpdateTopic(condition []string, params map[string]interface{}) error {
	var topic Topic
	err := db.Model(&topic).Where(condition).Updates(params).Error

	return err
}

func SaveTopic(topic Topic) error {
	err := db.Save(&topic).Error

	return err
}
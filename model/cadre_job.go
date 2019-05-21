package model

type CadreJob struct {
JobId int `gorm:"primary_key;AUTO_INCREMENT" json:"job_id"`
TownId int `json:"town_id"`
Name int `json:"name"`
Sort int `json:"sort"`
CreatedAt int `json:"created_at"`
}
func GetCadreJobById(id int) (CadreJob, bool) {
	var cadre_job CadreJob
	ok := db.First(&cadre_job, id).RecordNotFound()

	return cadre_job, ok
}

func GetCadreJobByOne(condition map[string]interface{}) (CadreJob, bool) {
	var cadre_job CadreJob
	ok := db.Where(condition).First(&cadre_job).RecordNotFound()

	return cadre_job, ok
}

func GetCadreJobList(condition map[string]interface{}, page int, pageSize int, orderBy []string) (interface{}, error) {
	cadre_jobs := make([]*CadreJob, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&cadre_jobs).
		Error

	return cadre_jobs, err
}

func GetCadreJobCount(condition map[string]interface{}) int {
	var cadre_job CadreJob
	count := 0
	db.Where(condition).Find(&cadre_job).Count(&count)

	return count
}

func GetPageUtil(condition map[string]interface{}, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetCadreJobList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetCadreJobCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteCadreJobByIds(ids string) (error) {
	var cadre_job CadreJob
	err := db.Where("id IN  ?", ids).Delete(&cadre_job).Error

	return err
}

func DeleteCadreJobById(id int64) error {
	var cadre_job CadreJob
	cadre_job.CadreJobId = id
	err := db.Delete(&cadre_job).Error

	return err
}

func CreateCadreJob(cadre_job CadreJob) (CadreJob, error) {
	err := db.Create(&cadre_job).Error

	return cadre_job, err
}

func UpdateCadreJob(condition map[string]interface{}, params map[string]interface{}) error {
	var cadre_job CadreJob
	err := db.Model(&cadre_job).Where(condition).Updates(params).Error

	return err
}

func SaveCadreJob(cadre_job CadreJob) error {
	err := db.Save(&cadre_job).Error

	return err
}
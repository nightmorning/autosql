package model

type CadreServer struct {
ServerId int `gorm:"primary_key;AUTO_INCREMENT" json:"server_id"`
TownId int `json:"town_id"`
Area int `json:"area"`
Sort int `json:"sort"`
CreatedAt int `json:"created_at"`
}
func GetCadreServerById(id int) (CadreServer, bool) {
	var cadre_server CadreServer
	ok := db.Where("server_id = ?", id).First(&cadre_server).RecordNotFound()

	return cadre_server, ok
}

func GetCadreServerByOne(condition []string) (CadreServer, bool) {
	var cadre_server CadreServer
	ok := db.Where(condition).First(&cadre_server).RecordNotFound()

	return cadre_server, ok
}

func GetCadreServerList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	cadre_servers := make([]*CadreServer, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&cadre_servers).
		Error

	return cadre_servers, err
}

func GetCadreServerCount(condition []string) int {
	var cadre_server CadreServer
	count := 0
	db.Where(condition).Find(&cadre_server).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetCadreServerList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetCadreServerCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteCadreServerByIds(ids string) (error) {
	var cadre_server CadreServer
	err := db.Where("id IN  ?", ids).Delete(&cadre_server).Error

	return err
}

func DeleteCadreServerById(id int64) error {
	var cadre_server CadreServer
	cadre_server.CadreServerId = id
	err := db.Delete(&cadre_server).Error

	return err
}

func CreateCadreServer(cadre_server CadreServer) (CadreServer, error) {
	err := db.Create(&cadre_server).Error

	return cadre_server, err
}

func UpdateCadreServer(condition []string, params map[string]interface{}) error {
	var cadre_server CadreServer
	err := db.Model(&cadre_server).Where(condition).Updates(params).Error

	return err
}

func SaveCadreServer(cadre_server CadreServer) error {
	err := db.Save(&cadre_server).Error

	return err
}
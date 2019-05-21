package model

type Station struct {
StationId int `gorm:"primary_key;json:"station_id"`
TownId int `json:"town_id"`
Name int `json:"name"`
Info int `json:"info"`
Address int `json:"address"`
CreatedAt int `json:"created_at"`
}
func GetStationById(id int) (Station, bool) {
	var station Station
	ok := db.Where("station_id = ?", id).First(&station).RecordNotFound()

	return station, ok
}

func GetStationByOne(condition []string) (Station, bool) {
	var station Station
	ok := db.Where(condition).First(&station).RecordNotFound()

	return station, ok
}

func GetStationList(condition []string, page int, pageSize int, orderBy []string) (interface{}, error) {
	stations := make([]*Station, 0)
	orderBys := strings.Join(orderBy, ",");

	offset := (page - 1)*pageSize
	err := db.
		Where(condition).
		Order(orderBys).
		Offset(offset).
		Limit(pageSize).
		Find(&stations).
		Error

	return stations, err
}

func GetStationCount(condition []string) int {
	var station Station
	count := 0
	db.Where(condition).Find(&station).Count(&count)

	return count
}

func GetPageUtil(condition []string, page int, pageSize int, orderBy []string) common.Page {
	list,err := GetStationList(condition, page, pageSize, orderBy)
	if err != nil {
		log.Fatal(err)
	}

	count := GetStationCount(condition)

	return common.PageUtil(count, page, pageSize, list)
}

func DeleteStationByIds(ids string) (error) {
	var station Station
	err := db.Where("id IN  ?", ids).Delete(&station).Error

	return err
}

func DeleteStationById(id int64) error {
	var station Station
	station.StationId = id
	err := db.Delete(&station).Error

	return err
}

func CreateStation(station Station) (Station, error) {
	err := db.Create(&station).Error

	return station, err
}

func UpdateStation(condition []string, params map[string]interface{}) error {
	var station Station
	err := db.Model(&station).Where(condition).Updates(params).Error

	return err
}

func SaveStation(station Station) error {
	err := db.Save(&station).Error

	return err
}
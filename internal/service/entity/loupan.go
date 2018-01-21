package entity

type Loupan struct {
	Id string `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Url string `json:"url" bson:"url"`
	Lng float64 `json:"lng" bson:"lng"`
	Lat float64 `json:"lat" bson:"lat"`
	StartDate int64 `json:"start_date" bson:"start_date"`
	StopDate int64 `json:"stop_date" bson:"stop_date"`
	FoundTime int64 `json:"found_time" bson:"found_time"`
}
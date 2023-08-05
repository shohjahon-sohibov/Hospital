package models

type Hospital struct {
	ID         string `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	Address    string `json:"address" bson:"address"`
	ImageUrl   string `json:"image_url" bson:"image_url"`
	CallCenter string `json:"call_center" bson:"call_center"`
	CreatedAt  string `json:"created_at" bson:"created_at"`
	UpdatedAt  string `json:"updated_at" bson:"updated_at"`
	DeletedAt  string `json:"deleted_at" bson:"deleted_at"`
}

type CreateHospital struct {
	Name       string `json:"name" bson:"name"`
	Address    string `json:"address" bson:"address"`
	ImageUrl   string `json:"image_url" bson:"image_url"`
	CallCenter string `json:"call_center" bson:"call_center"`
}

type UpdateHospital struct {
	ID         string `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	Address    string `json:"address" bson:"address"`
	ImageUrl   string `json:"image_url" bson:"image_url"`
	CallCenter string `json:"call_center" bson:"call_center"`
}

type GetListHospitalReq struct {
	Limit    int32  `json:"limit" bson:"limit"`
	Offset   int32  `json:"offset" bson:"offset"`
}

type GetListHospitalRes struct {
	Hospitals []*Hospital `json:"hospital" bson:"hospital"`
	Count     int32       `json:"count" bson:"count"`
}

package models

type Service struct {
	ID        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	ClinicId string `json:"clinic_id" bson:"clinic_id"`
	Price     string `json:"price" bson:"price"`
	CreatedAt string `json:"created_at" bson:"created_at"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
}

type CreateService struct {
	Name     string `json:"name" bson:"name"`
	ClinicId string `json:"clinic_id" bson:"clinic_id"`
	Price    string `json:"price" bson:"price"`
}

type UpdateService struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	ClinicId string `json:"clinic_id" bson:"clinic_id"`
	Price    string `json:"price" bson:"price"`
}

type GetListServiceReq struct {
	Limit  int32  `json:"limit" bson:"limit"`
	Offset int32  `json:"offset" bson:"offset"`
	ClinicId string `json:"clinic_id" bson:"clinic_id"`
}

type GetListServiceRes struct {
	Count    int32      `json:"count" bson:"count"`
	Services []*Service `json:"service" bson:"service"`
}

package models

type Diagnosis struct {
	ID          string           `json:"id" bson:"id"`
	UserId      string           `json:"user_id" bson:"user_id"`
	DoctorId    string           `json:"doctor_id" bson:"doctor_id"`
	Date        string           `json:"date" bson:"date"`
	Duration    string           `json:"duration" bson:"duration"`
	Description *DescriptionList `json:"description" bson:"description"`
	CreatedAt   string           `json:"created_at" bson:"created_at"`
	UpdatedAt   string           `json:"updated_at" bson:"updated_at"`
	DeletedAt   string           `json:"deleted_at" bson:"deleted_at"`
}

type CreateDiagnosis struct {
	UserId      string           `json:"user_id" bson:"user_id"`
	DoctorId    string           `json:"doctor_id" bson:"doctor_id"`
	Date        string           `json:"date" bson:"date"`
	Duration    string           `json:"duration" bson:"duration"`
	Description *DescriptionList `json:"description" bson:"description"`
	CreatedAt   string           `json:"created_at" bson:"created_at"`
	UpdatedAt   string           `json:"updated_at" bson:"updated_at"`
	DeletedAt   string           `json:"deleted_at" bson:"deleted_at"`
}

type UpdateDiagnosis struct {
	ID          string           `json:"id" bson:"id"`
	DoctorId    string           `json:"doctor_id" bson:"doctor_id"`
	Date        string           `json:"date" bson:"date"`
	Duration    string           `json:"duration" bson:"duration"`
	Description *DescriptionList `json:"description" bson:"description"`
}

type GetListDiagnosisReq struct {
	Limit    int32  `json:"limit" bson:"limit"`
	Offset   int32  `json:"offset" bson:"offset"`
	UserId   string `json:"user_id" bson:"user_id"`
	DoctorId string `json:"doctor_id" bson:"doctor_id"`
	Date     string `json:"date" bson:"date"`
}

type GetListDiagnosisRes struct {
	Diagnoses []*Diagnosis `json:"diagnoses" bson:"diagnoses"`
	Count     int32        `json:"count" bson:"count"`
}

type DescriptionList struct {
	Names []string `json:"names" bson:"names"`
}

type Id struct {
	ID string `json:"id" bson:"id"`
}

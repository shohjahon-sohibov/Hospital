package models

type Queue struct {
	ID            string `json:"id" bson:"id"`
	HospitalId    string `json:"hospital_id" bson:"hospital_id"`
	ServiceId     string `json:"service_id" bson:"service_id"`
	DoctorId      string `json:"doctor_id" bson:"doctor_id"`
	UserId        string `json:"user_id" bson:"user_id"`
	Floor         int32  `json:"floor" bson:"floor"`
	RoomNumber    int32  `json:"room_number" bson:"room_number"`
	QueueQuantity int32  `json:"queue_quantity" bson:"queue_quantity"`
	Status        string `json:"status" bson:"status"`
	CreatedAt     string `json:"created_at" bson:"created_at"`
	UpdatedAt     string `json:"updated_at" bson:"updated_at"`
	IsExpired     bool   `json:"is_expired" bson:"is_expired"`
}

type CreateQueue struct {
	HospitalId    string `json:"hospital_id" bson:"hospital_id"`
	ServiceId     string `json:"service_id" bson:"service_id"`
	DoctorId      string `json:"doctor_id" bson:"doctor_id"`
	UserId        string `json:"user_id" bson:"user_id"`
	UserName      string `json:"user_name" bson:"user_name"`
	Floor         int32  `json:"floor" bson:"floor"`
	RoomNumber    int32  `json:"room_number" bson:"room_number"`
	QueueQuantity int32  `json:"queue_quantity" bson:"queue_quantity"`
	Status        string `json:"status" bson:"status"`
}

type UpdateQueue struct {
	ID            string `json:"id" bson:"id"`
	HospitalId    string `json:"hospital_id" bson:"hospital_id"`
	ServiceId     string `json:"service_id" bson:"service_id"`
	DoctorId      string `json:"doctor_id" bson:"doctor_id"`
	Floor         int32  `json:"floor" bson:"floor"`
	RoomNumber    int32  `json:"room_number" bson:"room_number"`
	QueueQuantity int32  `json:"queue_quantity" bson:"queue_quantity"`
}

type GetListQueueReq struct {
	Limit    int32  `json:"limit" bson:"limit"`
	Offset   int32  `json:"offset" bson:"offset"`
	DoctorId string `json:"doctor_id" bson:"doctor_id"`
	Date     string `json:"date" bson:"date"`
}

type GetListQueueRes struct {
	Count  int32     `json:"count" bson:"count"`
	Queues []*Ticket `json:"ticket" bson:"ticket"`
}

type Ticket struct {
	QueueId       string `json:"queue_id" bson:"queue_id"`
	Hospital      string `json:"hospital" bson:"hospital"`
	Service       string `json:"service" bson:"service"`
	Doctor        string `json:"doctor" bson:"doctor"`
	DoctorId      string `json:"doctor_id" bson:"doctor_id"`
	User          string `json:"user" bson:"user"`
	UserPhone     string `json:"user_phone" bson:"user_phone"`
	Floor         int32  `json:"floor" bson:"floor"`
	RoomNumber    int32  `json:"room_number" bson:"room_number"`
	QueueQuantity int32  `json:"queue_quantity" bson:"queue_quantity"`
	Status        string `json:"status" bson:"status"`
	CreatedAt     string `json:"created_at" bson:"created_at"`
	IsExpired     bool   `json:"is_expired" bson:"is_expired"`
}

type ChangeStatusQueue struct {
	ID     string `json:"id" bson:"id"`
	Status string `json:"status" bson:"status"`
}

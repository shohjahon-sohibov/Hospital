package models

type User struct {
	ID           string `json:"id" bson:"id"`
	FirstName    string `json:"first_name" bson:"first_name"`
	LastName     string `json:"last_name" bson:"last_name"`
	Age          int32  `json:"age" bson:"age"`
	Username     string `json:"username" bson:"username"`
	PhoneNumber  string `json:"phone_number" bson:"phone_number"`
	Gmail        string `json:"gmail" bson:"gmail"`
	Password     string `json:"password" bson:"password"`
	Role         string `json:"role" bson:"role"`
	Gender       string `json:"gender" bson:"gender"`
	Token        string `json:"token" bson:"token"`
	ImageUrl     string `json:"image_url" bson:"image_url"`
	Speciality   string `json:"speciality" bson:"speciality"`
	Clinic       string `json:"clinic" bson:"clinic"`
	WorkTime     *Time  `json:"work_time" bson:"work_time"`
	WorkDay      *Days  `json:"work_day" bson:"work_day"`
	Floor        int32  `json:"floor" bson:"floor"`
	RoomNumber   int32  `json:"room_number" bson:"room_number"`
	Experience   string `json:"experience" bson:"experience"`
	ClientNumber int32 `json:"client_number" bson:"client_number"`
	CreatedAt    string `json:"created_at" bson:"created_at"`
	UpdatedAt    string `json:"updated_at" bson:"updated_at"`
	DeletedAt    string `json:"deleted_at" bson:"deleted_at"`
}

type CreateUser struct {
	ID          string `json:"id" bson:"id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Age         int32  `json:"age" bson:"age"`
	Username    string `json:"username" bson:"username"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Gmail       string `json:"gmail" bson:"gmail"`
	Password    string `json:"password" bson:"password"`
	Role        string `json:"role" bson:"role"`
	Gender      string `json:"gender" bson:"gender"`
	Token       string `json:"token" bson:"token"`
	ImageUrl    string `json:"image_url" bson:"image_url"`
	Speciality  string `json:"speciality" bson:"speciality"`
	ServiceId   string `json:"service_id" bson:"service_id"`
	Clinic      string `json:"clinic" bson:"clinic"`
	ClinicId    string `json:"clinic_id" bson:"clinic_id"`
	WorkTime    *Time  `json:"work_time" bson:"work_time"`
	WorkDay     *Days  `json:"work_day" bson:"work_day"`
	Floor       int32  `json:"floor" bson:"floor"`
	RoomNumber  int32  `json:"room_number" bson:"room_number"`
	Experience  string `json:"experience" bson:"experience"`
}

type UpdateUser struct {
	ID          string `json:"id" bson:"id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	Age         int32  `json:"age" bson:"age"`
	Username    string `json:"username" bson:"username"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Gmail       string `json:"gmail" bson:"gmail"`
	Password    string `json:"password" bson:"password"`
	Role        string `json:"role" bson:"role"`
	Gender      string `json:"gender" bson:"gender"`
	Token       string `json:"token" bson:"token"`
	ImageUrl    string `json:"image_url" bson:"image_url"`
	Speciality  string `json:"speciality" bson:"speciality"`
	ServiceId   string `json:"service_id" bson:"service_id"`
	Clinic      string `json:"clinic" bson:"clinic"`
	ClinicId    string `json:"clinic_id" bson:"clinic_id"`
	WorkTime    *Time  `json:"work_time" bson:"work_time"`
	WorkDay     *Days  `json:"work_day" bson:"work_day"`
	Floor       int32  `json:"floor" bson:"floor"`
	RoomNumber  int32  `json:"room_number" bson:"room_number"`
	Experience  string `json:"experience" bson:"experience"`
}

type GetUsersListReq struct {
	Limit       int32  `json:"limit" bson:"limit"`
	Offset      int32  `json:"offset" bson:"offset"`
	Role        string `json:"role" bson:"role"`
	ClinicId    string `json:"clinic_id" bson:"clinic_id"`
	ServiceId   string `json:"service_id" bson:"service_id"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type GetUsersListRes struct {
	Count int32   `json:"count" bson:"count"`
	Users []*User `json:"users" bson:"users"`
}

type ClientID struct {
	ID string `json:"id" bson:"id"`
}

type Message struct {
	Message string `json:"message" bson:"message"`
	Success bool   `json:"success" bson:"success"`
}

type Days struct {
	Monday    bool `json:"monday" bson:"monday"`
	Tuesday   bool `json:"tuesday" bson:"tuesday"`
	Wednesday bool `json:"wednesday" bson:"wednesday"`
	Thursday  bool `json:"thursday" bson:"thursday"`
	Friday    bool `json:"friday" bson:"friday"`
	Saturday  bool `json:"saturday" bson:"saturday"`
	Sunday    bool `json:"sunday" bson:"sunday"`
}

type Time struct {
	Day   bool `json:"day" bson:"day"`
	Night bool `json:"night" bson:"night"`
}

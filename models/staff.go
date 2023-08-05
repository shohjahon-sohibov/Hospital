package models

type Staffs struct {
	ID         string `json:"id" bson:"id"`
	UserId     string `json:"user_id" bson:"user_id"`
	ImageUrl   string `json:"image_url" bson:"image_url"`
	Speciality string `json:"speciality" bson:"speciality"`
	Clinic     string `json:"clinic" bson:"clinic"`
	WorkTime   *Time  `json:"work_time" bson:"work_time"`
	WorkDay    *Days  `json:"work_day" bson:"work_day"`
	Floor      int32  `json:"floor" bson:"floor"`
	RoomNumber int32  `json:"room_number" bson:"room_number"`
	Experience string `json:"experience" bson:"experience"`
	CreatedAt  string `json:"created_at" bson:"created_at"`
	UpdatedAt  string `json:"updated_at" bson:"updated_at"`
	DeletedAt  string `json:"deleted_at" bson:"deleted_at"`
}

type CreateStaff struct {
	ID         string `json:"id" bson:"id"`
	ImageUrl   string `json:"image_url" bson:"image_url"`
	Speciality string `json:"speciality" bson:"speciality"`
	Clinic     string `json:"clinic" bson:"clinic"`
	Worktime   *Time  `json:"work_time" bson:"work_time"`
	WorkDay    *Days  `json:"work_day" bson:"work_day"`
	Floor      string `json:"floor" bson:"floor"`
	RoomNumber string `json:"room_number" bson:"room_number"`
	Experience string `json:"experience" bson:"experience"`
}

type UpdateStaff struct {
	ID         string `json:"id" bson:"id"`
	ImageUrl   string `json:"image_url" bson:"image_url"`
	Speciality string `json:"speciality" bson:"speciality"`
	Clinic     string `json:"clinic" bson:"clinic"`
	WorkTime   *Time  `json:"work_time" bson:"work_time"`
	WorkDay    *Days  `json:"work_day" bson:"work_day"`
	Floor      string `json:"floor" bson:"floor"`
	RoomNumber string `json:"room_number" bson:"room_number"`
	Experience string `json:"experience" bson:"experience"`
}

type GetStaffsListReq struct {
	Limit      int64  `json:"limit" bson:"limit"`
	Offset     int64  `json:"offset" bson:"offset"`
	Floor      string `json:"floor" bson:"floor"`
	Speciality string `json:"speciality" bson:"speciality"`
	Clinic     string `json:"clinic" bson:"clinic"`
	Experience string `json:"experience" bson:"experience"`
	WorkTime   *Time  `json:"work_time" bson:"work_time"`
	WorkDay    *Days  `json:"work_day" bson:"work_day"`
}

type GetStaffsListRes struct {
	Count  int64     `json:"count" bson:"count"`
	Staffs []*Staffs `json:"doctors" bson:"doctors"`
}

type StudentID struct {
	ID string `json:"id" bson:"id"`
}

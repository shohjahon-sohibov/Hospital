package mongodb

import (
	"context"
	"errors"
	"fmt"
	"freelance/clinic_queue/models"
	"freelance/clinic_queue/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	usersCollection  *mongo.Collection
	staffsCollection *mongo.Collection
	queueCollection  *mongo.Collection
}

func NewUserRepo(db *mongo.Database) storage.UserI {
	return &userRepo{
		usersCollection:  db.Collection("users"),
		staffsCollection: db.Collection("staffs"),
		queueCollection:  db.Collection("queues"),
	}
}

func (u userRepo) CreateUser(ctx context.Context, req *models.CreateUser) (resp *models.Id, err error) {
	id := primitive.NewObjectID()
	con := bson.M{
		"id":           id.Hex(),
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"age":          req.Age,
		"username":     req.Username,
		"phone_number": req.PhoneNumber,
		"gmail":        req.Gmail,
		"password":     req.Password,
		"role":         req.Role,
		"created_at":   time.Now().Format("02.01.2006 15:04:05"),
	}
	_, err = u.usersCollection.InsertOne(ctx, con)
	if err != nil {
		return nil, errors.New("Error while insert user")
	}

	if req.Clinic != "" {
		con := bson.M{
			"user_id":     id.Hex(),
			"image_url":   req.ImageUrl,
			"speciality":  req.Speciality,
			"service_id":  req.ServiceId,
			"clinic":      req.Clinic,
			"clinic_id":   req.ClinicId,
			"work_time":   req.WorkTime,
			"work_day":    req.WorkDay,
			"floor":       req.Floor,
			"room_number": req.RoomNumber,
			"experience":  req.Experience,
			"created_at":  time.Now().Format("02.01.2006 15:04:05"),
		}
		_, err := u.staffsCollection.InsertOne(ctx, con)
		if err != nil {
			return nil, errors.New("Error while insert staff err")
		}

	}
	return &models.Id{
		ID: id.Hex(),
	}, nil
}

func (u userRepo) UpdateUser(ctx context.Context, req *models.UpdateUser) (res *models.Message, err error) {
	if req.ID == "" {
		return nil, fmt.Errorf("Id is not provided, error")
	}
	updateReq := bson.M{
		"$set": bson.M{
			"first_name":   req.FirstName,
			"last_name":    req.LastName,
			"age":          req.Age,
			"username":     req.Username,
			"phone_number": req.PhoneNumber,
			"gmail":        req.Gmail,
			"password":     req.Password,
			"role":         req.Role,
			"updated_at":   time.Now().Format("02.01.2006 15:04:05"),
		},
	}
	resp, err := u.usersCollection.UpdateOne(ctx, bson.M{"id": req.ID, "deleted_at": nil}, updateReq)
	if resp.MatchedCount == 0 {
		fmt.Println("Document not found for ID:", req.ID)
		return nil, fmt.Errorf("Document not found for ID: %s", req.ID)
	} else if err != nil {
		fmt.Println("Error while updating doagnosis:", err.Error())
		return nil, err
	}

	updateStaffReq := bson.M{
		"$set": bson.M{
			"gender":      req.Gender,
			"image_url":   req.ImageUrl,
			"speciality":  req.Speciality,
			"service_id":  req.ServiceId,
			"clinic":      req.Clinic,
			"clinic_id":   req.ClinicId,
			"floor":       req.Floor,
			"room_number": req.RoomNumber,
			"experience":  req.Experience,
			"updated_at":  time.Now().Format("02.01.2006 15:04:05"),
		},
	}

	resp, err = u.staffsCollection.UpdateOne(ctx, bson.M{"user_id": req.ID, "deleted_at": nil}, updateStaffReq)
	if resp.MatchedCount == 0 {
		fmt.Println("Document not found for ID:", req.ID)
		return nil, fmt.Errorf("Document not found for ID: %s", req.ID)
	} else if err != nil {
		fmt.Println("Error while updating doagnosis:", err.Error())
		return nil, err
	}

	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}

func (u userRepo) GetSingleUser(ctx context.Context, id string) (res *models.User, err error) {
	res = &models.User{}
	if id == "" {
		return nil, fmt.Errorf("Id is not provided")
	}
	filter := bson.M{}
	filter["id"] = id
	filter["deleted_at"] = nil

	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "staffs",
				"localField":   "id",
				"foreignField": "user_id",
				"as":           "staffs",
			},
		},
		{
			"$unwind": "$staffs",
		},
		{
			"$project": bson.M{
				"id":           1,
				"first_name":   1,
				"last_name":    1,
				"age":          1,
				"username":     1,
				"phone_number": 1,
				"gmail":        1,
				"role":         1,
				"gender":       1,
				"token":        1,
				"image_url":    "$staffs.image_url",
				"speciality":   "$staffs.speciality",
				"clinic":       "$staffs.clinic",
				"work_time":    "$staffs.work_time",
				"work_day":     "$staffs.work_day",
				"floor":        "$staffs.floor",
				"room_number":  "$staffs.room_number",
				"experience":   "$staffs.experience",
				"created_at":   1,
				"updated_at":   1,
				"deleted_at":   1,
			},
		},
		{
			"$match": filter,
		},
	}

	cursor, err := u.usersCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("aggregate single user error: %s", err.Error())
	}

	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		if err := cursor.Decode(res); err != nil {
			fmt.Println("Error decoding user:", err)
		}
	} else {
		return nil, fmt.Errorf("user not found")
	}
	filterClient := bson.M{}
	filterClient["doctor_id"] = id
	// filterClient["created_at"] = bson.M{
	// 	"$gte": time.Now().Add(-24 * time.Hour),
	// 	"$lte": time.Now(),
	// }
	count, err := u.queueCollection.CountDocuments(ctx, filterClient)
	if err != nil {
		return nil, fmt.Errorf("Error insert hospital, err: %v", err)

	}
	fmt.Println(count, 1)
	if count > 0 {
		res.ClientNumber = int32(count)
	}
	return res, nil
}

func (u userRepo) GetListUser(ctx context.Context, req *models.GetUsersListReq) (res *models.GetUsersListRes, err error) {
	res = &models.GetUsersListRes{}

	filter := bson.M{}
	if req.Role != "" {
		filter["role"] = req.Role
	}
	if req.ServiceId != "" {
		filter["service_id"] = req.ServiceId
		fmt.Println(1, req.ServiceId)
	}
	if req.ClinicId != "" {
		filter["clinic_id"] = req.ClinicId
		fmt.Println(2, req.ClinicId)
	}
	filter["deleted_at"] = nil

	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "staffs",
				"localField":   "id",
				"foreignField": "user_id",
				"as":           "staffs",
			},
		},
		{
			"$unwind": "$staffs",
		},
		{
			"$project": bson.M{
				"id":           1,
				"first_name":   1,
				"last_name":    1,
				"age":          1,
				"username":     1,
				"phone_number": 1,
				"gmail":        1,
				"role":         1,
				"gender":       1,
				"token":        1,
				"image_url":    "$staffs.image_url",
				"speciality":   "$staffs.speciality",
				"service_id":   "$staffs.service_id",
				"clinic":       "$staffs.clinic",
				"clinic_id":    "$staffs.clinic_id",
				"work_time":    "$staffs.work_time",
				"work_day":     "$staffs.work_day",
				"floor":        "$staffs.floor",
				"room_number":  "$staffs.room_number",
				"experience":   "$staffs.experience",
				"created_at":   1,
				"updated_at":   1,
			},
		},
		{
			"$match": filter,
		},
		{
			"$sort": bson.M{"created_at": -1},
		},
		{
			"$skip": req.Offset,
		},
		{
			"$limit": req.Limit,
		},
	}

	cursor, err := u.usersCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("get list users error: %s", err.Error())
	}

	defer cursor.Close(ctx)

	// Extract the result
	if err := cursor.All(ctx, &res.Users); err != nil {
		fmt.Println("Error while getting all, err: ", err)
	}

	pipelineCount := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "staffs",
				"localField":   "id",
				"foreignField": "user_id",
				"as":           "staffs",
			},
		},
		{
			"$unwind": "$staffs",
		},
		{
			"$project": bson.M{
				"id":           1,
				"first_name":   1,
				"last_name":    1,
				"age":          1,
				"username":     1,
				"phone_number": 1,
				"gmail":        1,
				"role":         1,
				"gender":       1,
				"token":        1,
				"image_url":    "$staffs.image_url",
				"speciality":   "$staffs.speciality",
				"service_id":   "$staffs.service_id",
				"clinic":       "$staffs.clinic",
				"clinic_id":    "$staffs.clinic_id",
				"work_time":    "$staffs.work_time",
				"work_day":     "$staffs.work_day",
				"floor":        "$staffs.floor",
				"room_number":  "$staffs.room_number",
				"experience":   "$staffs.experience",
				"created_at":   1,
				"updated_at":   1,
			},
		},
		{
			"$match": filter,
		},
		{
			"$sort": bson.M{"created_at": -1},
		},
		{
			"$skip": req.Offset,
		},
		{
			"$limit": req.Limit,
		},
		{
			"$count": "count",
		},
	}

	rows, err := u.usersCollection.Aggregate(ctx, pipelineCount)
	if err != nil {
		return nil, fmt.Errorf("count users error: %s", err.Error())
	}

	defer cursor.Close(ctx)

	// Extract the result
	var countResult []bson.M
	if err := rows.All(ctx, &countResult); err != nil {
		fmt.Println("Error while counting all, err: ", err)
	}

	if len(countResult) > 0 {
		res.Count = countResult[0]["count"].(int32)

	}
	return res, nil
}

func (u userRepo) DeleteUser(ctx context.Context, id string) (res *models.Message, err error) {
	fmt.Println(1)
	_, err = u.usersCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now().Format("02.01.2006 15:04:05")}})
	if err != nil {
		fmt.Println("Error while deleting user:", err)
		return nil, err
	}
	fmt.Println(2)

	_, err = u.staffsCollection.UpdateOne(ctx, bson.M{"user_id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now().Format("02.01.2006 15:04:05")}})
	if err != nil {
		fmt.Println("Error while deleting staff:", err)
		return nil, err
	}
	fmt.Println(3)

	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}

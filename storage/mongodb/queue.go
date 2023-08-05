package mongodb

import (
	"context"
	"fmt"
	"freelance/clinic_queue/models"
	"freelance/clinic_queue/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type queueRepo struct {
	queueCollection     *mongo.Collection
	usersCollection     *mongo.Collection
	hospitalCollection  *mongo.Collection
	staffsCollection    *mongo.Collection
	serviceCollection   *mongo.Collection
	diagnosisCollection *mongo.Collection
}

func NewQueueRepo(db *mongo.Database) storage.QueueI {
	return &queueRepo{
		queueCollection:     db.Collection("queues"),
		usersCollection:     db.Collection("users"),
		hospitalCollection:  db.Collection("hospitals"),
		staffsCollection:    db.Collection("staffs"),
		serviceCollection:   db.Collection("services"),
		diagnosisCollection: db.Collection("diagnoses"),
	}
}

func (q queueRepo) CreateQueue(ctx context.Context, req *models.CreateQueue) (res *models.Ticket, err error) {
	var (
		resp     = &models.Queue{}
		user     = &models.User{}
		doctor   = &models.User{}
		hospital = &models.Hospital{}
		service  = &models.Service{}
		ticket   = &models.Ticket{}
	)
	fmt.Println(resp)

	err = q.usersCollection.FindOne(ctx, bson.M{"id": req.UserId}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("get one user when create queue error: %s", err.Error())
	}

	err = q.usersCollection.FindOne(ctx, bson.M{"id": req.DoctorId}).Decode(&doctor)
	if err != nil {
		return nil, fmt.Errorf("get one user when create queue error: %s", err.Error())
	}

	err = q.hospitalCollection.FindOne(ctx, bson.M{"id": req.HospitalId}).Decode(&hospital)
	if err != nil {
		return nil, fmt.Errorf("get one hospital when create queue error: %s", err.Error())
	}

	err = q.serviceCollection.FindOne(ctx, bson.M{"id": req.ServiceId}).Decode(&service)
	if err != nil {
		return nil, fmt.Errorf("get one hospital when create queue error: %s", err.Error())
	}

	id := primitive.NewObjectID()

	con := bson.M{
		"id":          id.Hex(),
		"hospital_id": req.HospitalId,
		"service_id":  req.ServiceId,
		"doctor_id":   req.DoctorId,
		"user_id":     req.UserId,
		"floor":       req.Floor,
		"room_number": req.RoomNumber,
		"status":      "waiting",
		"created_at":  time.Now().Format("02.01.2006 15:04:05"),
	}

	_, err = q.queueCollection.InsertOne(ctx, con)
	if err != nil {
		return nil, fmt.Errorf("Error insert hospital, err: %v", err)

	}

	count, err := q.diagnosisCollection.CountDocuments(ctx, bson.M{ "user_id": req.UserId })
	if err != nil {
		return nil, fmt.Errorf("Error insert hospital, err: %v", err)

	}

	fmt.Println(count, 1)

	ticket = &models.Ticket{
		QueueId:       id.Hex(),
		Hospital:      hospital.Name,
		Service:       service.Name,
		Doctor:        doctor.FirstName + " " + doctor.LastName,
		User:          user.Username,
		UserPhone:     user.PhoneNumber,
		Floor:         req.Floor,
		RoomNumber:    req.RoomNumber,
		QueueQuantity: int32(count),
		Status:        "waiting",
		CreatedAt:     time.Now().Format("02.01.2006 15:04:05"),
	}

	return ticket, nil
}

func (q queueRepo) ChangeStatusQueue(ctx context.Context, req *models.ChangeStatusQueue) (res *models.Message, err error) {
	if req.ID == "" {
		return nil, fmt.Errorf("Id is not provided, error")
	}
	updateReq := bson.M{
		"$set": bson.M{
			"status":     "done",
			"updated_at": time.Now().Format("02.01.2006 15:04:05"),
		},
	}
	resp, err := q.queueCollection.UpdateOne(ctx, bson.M{"id": req.ID}, updateReq)
	if resp.MatchedCount == 0 {
		fmt.Println("Document not found for ID:", req.ID)
		return nil, fmt.Errorf("Document not found for ID: %s", req.ID)
	} else if err != nil {
		fmt.Println("Error while updating hospital:", err.Error())
		return nil, err
	}
	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}

func (q queueRepo) GetListQueue(ctx context.Context, req *models.GetListQueueReq) (res *models.GetListQueueRes, err error) {
	res = &models.GetListQueueRes{}

	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "hospitals",
				"localField":   "hospital_id",
				"foreignField": "id",
				"as":           "hospital",
			},
		},
		{
			"$unwind": "$hospital",
		},
		{
			"$lookup": bson.M{
				"from":         "services",
				"localField":   "service_id",
				"foreignField": "id",
				"as":           "service",
			},
		},
		{
			"$unwind": "$service",
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "user_id",
				"foreignField": "id",
				"as":           "user",
			},
		},
		{
			"$unwind": "$user",
		},
		{
			"$project": bson.M{
				"doctor_id":      1,
				"hospital":       "$hospital.name",
				"service":        "$service.name",
				"doctor":         "$user.username",
				"user":           "$user.username",
				"user_phone":     "$user.phone_number",
				"floor":          1,
				"room_number":    1,
				"queue_quantity": 1,
				"status":         1,
				"created_at":     1,
			},
		},
		{
			"$match": bson.M{
				"deleted_at": nil,
			},
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

	cursor, err := q.queueCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("get list hospitals error: %s", err.Error())
	}

	defer cursor.Close(ctx)

	// Extract the result
	if err := cursor.All(ctx, &res.Queues); err != nil {
		fmt.Println("Error while getting all, err: ", err)
	}

	pipelineCount := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "hospitals",
				"localField":   "hospital_id",
				"foreignField": "id",
				"as":           "hospital",
			},
		},
		{
			"$unwind": "$hospital",
		},
		{
			"$lookup": bson.M{
				"from":         "services",
				"localField":   "service_id",
				"foreignField": "id",
				"as":           "service",
			},
		},
		{
			"$unwind": "$service",
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "user_id",
				"foreignField": "id",
				"as":           "user",
			},
		},
		{
			"$unwind": "$user",
		},
		{
			"$project": bson.M{
				"doctor_id":      1,
				"hospital":       "$hospital.name",
				"service":        "$service.name",
				"doctor":         "$user.username",
				"user":           "$user.username",
				"user_phone":     "$user.phone_number",
				"floor":          1,
				"room_number":    1,
				"queue_quantity": 1,
				"status":         1,
				"created_at":     1,
			},
		},
		{
			"$match": bson.M{
				"deleted_at": nil,
			},
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

	rows, err := q.queueCollection.Aggregate(ctx, pipelineCount)
	if err != nil {
		return nil, fmt.Errorf("count queues error: %s", err.Error())
	}

	defer cursor.Close(ctx)

	// Extract the result
	var countResult []bson.M
	if err := rows.All(ctx, &countResult); err != nil {
		fmt.Println("Error while getting all, err: ", err)
	}

	res.Count = countResult[0]["count"].(int32)
	return res, nil
}

func (q queueRepo) DeleteQueue(ctx context.Context, id string) (res *models.Message, err error) {
	_, err = q.queueCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now().Format("02.01.2006 15:04:05")}})
	if err != nil {
		fmt.Println("Error while deleting hospital:", err)
		return nil, err
	}
	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}

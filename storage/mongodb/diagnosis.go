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

type diagnosisRepo struct {
	diagnosisCollection *mongo.Collection
}

func NewDiagnosisRepo(db *mongo.Database) storage.DiagnosisI {
	return *&diagnosisRepo{
		diagnosisCollection: db.Collection("diagnoses"),
	}
}

func (d diagnosisRepo) CreateDiagnosis(ctx context.Context, req *models.CreateDiagnosis) (res *models.Message, err error) {
	Id := primitive.NewObjectID()

	con := bson.M{
		"id":          Id.Hex(),
		"user_id":     req.UserId,
		"doctor_id":   req.DoctorId,
		"date":        req.Date,
		"duration":    req.Duration,
		"description": req.Description,
		"created_at":  time.Now().Format("02.01.2006 15:04"),
		"deleted_at":  nil,
	}

	_, err = d.diagnosisCollection.InsertOne(ctx, con)
	if err != nil {
		return nil, fmt.Errorf("Error insert diagnosis, err: %v", err)
	}

	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}

func (d diagnosisRepo) GetSingleDiagnosis(ctx context.Context, id string) (res *models.Diagnosis, err error) {
	res = &models.Diagnosis{}
	err = d.diagnosisCollection.FindOne(ctx, bson.M{"id": &id, "deleted_at": nil}).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("Error get single diagnosis, err: %v", err)
	}

	return res, nil
}

func (d diagnosisRepo) GetListDiagnosis(ctx context.Context, req *models.GetListDiagnosisReq) (res *models.GetListDiagnosisRes, err error) {
	res = &models.GetListDiagnosisRes{}
	filter := bson.M{}

	if req.Date != "" {
		filter["date"] = req.Date
	}
	if req.DoctorId != "" {
		filter["doctor_id"] = req.DoctorId
	}
	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	filter["deleted_at"] = nil

	pipeline := []bson.M{
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

	cursor, err := d.diagnosisCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("get list diagnosis error: %s", err.Error())
	}

	defer cursor.Close(ctx)

	// Extract the result
	if err := cursor.All(ctx, &res.Diagnoses); err != nil {
		fmt.Println("Error while getting all, err: ", err)
	}

	pipelineCount := []bson.M{
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

	rows, err := d.diagnosisCollection.Aggregate(ctx, pipelineCount)
	if err != nil {
		return nil, fmt.Errorf("count Documents error: %s", err.Error())
	}

	defer cursor.Close(ctx)

	// Extract the result
	var countResult []bson.M
	if err := rows.All(ctx, &countResult); err != nil {
		fmt.Println("Error while getting all, err: ", err)
	}

	if len(countResult) > 0 {
		res.Count = countResult[0]["count"].(int32)

	}
	return res, nil
}

func (d diagnosisRepo) UpdateDiagnosis(ctx context.Context, req *models.UpdateDiagnosis) (res *models.Message, err error) {
	if req.ID == "" {
		return nil, fmt.Errorf("Id is not provided, error")
	}
	updateReq := bson.M{
		"$set": bson.M{
			"doctor_id":   req.DoctorId,
			"date":        req.Date,
			"duration":    req.Duration,
			"description": req.Description,
			"updated_at":  time.Now().Format("02.01.2006 15:04"),
		},
	}
	resp, err := d.diagnosisCollection.UpdateOne(ctx, bson.M{"id": req.ID, "deleted_at": nil}, updateReq)
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

func (d diagnosisRepo) DeleteDiagnosis(ctx context.Context, id string) (res *models.Message, err error) {
	_, err = d.diagnosisCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now().Format("02.01.2006 15:04")}})
	if err != nil {
		fmt.Println("Error while deleting diagnosis:", err)
		return nil, err
	}
	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}
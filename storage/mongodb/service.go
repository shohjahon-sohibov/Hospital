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

type serviceRepo struct {
	serviceCollection *mongo.Collection
}

func NewServiceRepo(db *mongo.Database) storage.ServiceI {
	return &serviceRepo{
		serviceCollection: db.Collection("services"),
	}
}

func (s serviceRepo) CreateService(ctx context.Context, req *models.CreateService) (res *models.Message, err error) {
	id := primitive.NewObjectID()

	con := bson.M{
		"id":         id.Hex(),
		"name":       req.Name,
		"clinic_id":  req.ClinicId,
		"price":      req.Price,
		"created_at": time.Now().Format("02.01.2006 15:04:05"),
	}

	_, err = s.serviceCollection.InsertOne(ctx, con)
	if err != nil {
		return nil, fmt.Errorf("Error insert hospital, err: %v", err)

	}

	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}

func (s serviceRepo) GetListService(ctx context.Context, req *models.GetListServiceReq) (res *models.GetListServiceRes, err error) {
	res = &models.GetListServiceRes{}
	filter := bson.M{}
	
	if req.ClinicId != "" {
		filter["clinic_id"] = req.ClinicId
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

	cursor, err := s.serviceCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("get list diagnosis error: %s", err.Error())
	}

	defer cursor.Close(ctx)

	// Extract the result
	if err := cursor.All(ctx, &res.Services); err != nil {
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

	rows, err := s.serviceCollection.Aggregate(ctx, pipelineCount)
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

func (s serviceRepo) UpdateService(ctx context.Context, req *models.UpdateService) (res *models.Message, err error) {
	if req.ID == "" {
		return nil, fmt.Errorf("Id is not provided, error")
	}
	updateReq := bson.M{
		"$set": bson.M{
			"name":       req.Name,
			"clinic_id":  req.ClinicId,
			"price":      req.Price,
			"updated_at": time.Now().Format("02.01.2006 15:04"),
		},
	}
	resp, err := s.serviceCollection.UpdateOne(ctx, bson.M{"id": req.ID, "deleted_at": nil}, updateReq)
	if resp.MatchedCount == 0 {
		fmt.Println("Document not found for ID:", req.ID)
		return nil, fmt.Errorf("Document not found for ID: %s", req.ID)
	} else if err != nil {
		fmt.Println("Error while updating service:", err.Error())
		return nil, err
	}

	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}

func (s serviceRepo) DeleteService(ctx context.Context, id string) (res *models.Message, err error) {
	_, err = s.serviceCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now().Format("02.01.2006 15:04")}})
	if err != nil {
		fmt.Println("Error while deleting diagnosis:", err)
		return nil, err
	}
	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}

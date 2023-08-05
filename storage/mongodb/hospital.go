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

type hospitalRepo struct {
	hospitalCollection *mongo.Collection
}

func NewHospitalRepo(db *mongo.Database) storage.HospitalI {
	return &hospitalRepo{
		hospitalCollection: db.Collection("hospitals"),
	}
}

func (h hospitalRepo) CreateHospital(ctx context.Context, req *models.CreateHospital) (res *models.Message, err error) {
	id := primitive.NewObjectID()

	con := bson.M{
		"id":          id.Hex(),
		"name":        req.Name,
		"address":     req.Address,
		"image_url":   req.ImageUrl,
		"call_center": req.CallCenter,
		"created_at":  time.Now().Format("02.01.2006 15:04:05"),
	}

	_, err = h.hospitalCollection.InsertOne(ctx, con)
	if err != nil {
		return nil, fmt.Errorf("Error insert hospital, err: %v", err)

	}

	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}

func (h hospitalRepo) UpdateHospital(ctx context.Context, req *models.UpdateHospital) (res *models.Message, err error) {
	if req.ID == "" {
		return nil, fmt.Errorf("Id is not provided, error")
	}
	updateReq := bson.M{
		"$set": bson.M{
			"name":        req.Name,
			"address":     req.Address,
			"image_url":   req.ImageUrl,
			"call_center": req.CallCenter,
			"updated_at":  time.Now().Format("02.01.2006 15:04:05"),
		},
	}
	resp, err := h.hospitalCollection.UpdateOne(ctx, bson.M{"id": req.ID}, updateReq)
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

func (h hospitalRepo) GetSingleHospital(ctx context.Context, id string) (res *models.Hospital, err error) {
	res = &models.Hospital{}
	err = h.hospitalCollection.FindOne(ctx, bson.M{"id": &id, "deleted_at": nil}).Decode(res)
	if err != nil {
		return nil, fmt.Errorf("Error get single diagnosis, err: %v", err)
	}

	return res, nil
}

func (h hospitalRepo) GetListHospital(ctx context.Context, req *models.GetListHospitalReq) (res *models.GetListHospitalRes, err error) {
	res = &models.GetListHospitalRes{}

	pipeline := []bson.M{
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

	cursor, err := h.hospitalCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("get list hospitals error: %s", err.Error())
	}

	defer cursor.Close(ctx)

	// Extract the result
	if err := cursor.All(ctx, &res.Hospitals); err != nil {
		fmt.Println("Error while getting all, err: ", err)
	}
	
	pipelineCount := []bson.M{
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

	rows, err := h.hospitalCollection.Aggregate(ctx, pipelineCount)
	if err != nil {
		return nil, fmt.Errorf("count hospitals error: %s", err.Error())
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

func (h hospitalRepo) DeleteHospital(ctx context.Context, id string) (res *models.Message, err error) {
	_, err = h.hospitalCollection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now().Format("02.01.2006 15:04:05")}})
	if err != nil {
		fmt.Println("Error while deleting hospital:", err)
		return nil, err
	}
	return &models.Message{
		Message: "success",
		Success: true,
	}, nil
}
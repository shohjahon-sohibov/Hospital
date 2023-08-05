package mongodb

import (
	"context"
	"errors"
	"fmt"
	"freelance/clinic_queue/models"
	"freelance/clinic_queue/pkg/helper"
	"freelance/clinic_queue/storage"
	"time"

	"github.com/saidamir98/udevs_pkg/security"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authRepo struct {
	usersCollection   *mongo.Collection
	staffsCollection *mongo.Collection
}

func NewAuthRepo(db *mongo.Database) storage.AuthI {
	return &authRepo{
		usersCollection:   db.Collection("users"),
		staffsCollection: db.Collection("staffs"),
	}
}

func (a authRepo) SignUp(ctx context.Context, req *models.SignUp) (res *models.Token, err error) {
	// check is user exist
	var (
		resp models.User
	)
	err = a.usersCollection.FindOne(ctx, bson.M{"username": req.Username}).Decode(&resp)
	fmt.Println(resp.ID, 1)
	if resp.ID != "" {
		fmt.Println("User exist")
		return nil, fmt.Errorf("User exist")
	}

	if len(req.Password) < 6 {
		err := fmt.Errorf("password must not be less than 6 characters")
		return nil, err
	}

	// create user if not exists
	id := primitive.NewObjectID()

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := helper.GenerateToken(req.Username)
	if err != nil {
		fmt.Println("Token Generate err: ", err)
		return nil, err
	}

	req.Password = hashedPassword
	con := bson.M{
		"id":           id.Hex(),
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"age":          req.Age,
		"username":     req.Username,
		"phone_number": req.PhoneNumber,
		"gmail":        req.Gmail,
		"password":     req.Password,
		"token":        token,
		"role":         "user",
		"created_at":   time.Now().Format("02.01.2006 15:04"),
	}

	_, err = a.usersCollection.InsertOne(ctx, con)
	if err != nil {
		fmt.Println("Client insert err: ", err)
		return nil, err
	}

	return &models.Token{
		Token:   token,
		Success: true,
	}, nil
}

func (a authRepo) Login(ctx context.Context, req *models.Login) (res *models.Token, err error) {
	if req.Username != "" && req.Password != "" {
		res = &models.Token{}
		user := &models.User{}
		err := a.usersCollection.FindOne(ctx, bson.M{"username": req.Username}).Decode(&user)
		if err != nil {
			return nil, fmt.Errorf("Username is invalid")
		}

		match, err := security.ComparePassword(user.Password, req.Password)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if !match {
			fmt.Println(111)
			return nil, errors.New("username or password is wrong")
		}

		if req.Token != "" {
			_, err = a.usersCollection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": bson.M{"token": req.Token}})
			if err != nil {
				fmt.Println(err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		} else {
			token, err := helper.GenerateToken(req.Username)
			if err != nil {
				fmt.Println("Token Generate err: ", err)
				return nil, err
			}

			_, err = a.usersCollection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": bson.M{"token": token}})
			if err != nil {
				fmt.Println(err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			res.Token = token
		}
		res.Success = true
		return res, nil
	} else {
		err := fmt.Errorf("password or username is not prodvided")
		return nil, err
	}
}

// func (s studentRepo) Single(ctx context.Context, in string) (*models.Student, error) {
// 	filter := bson.M{}
// 	if in != "" {
// 		filter["id"] = in
// 	}

// 	pipeline := []bson.M{
// 		{
// 			"$lookup": bson.M{
// 				"from":         "branches",
// 				"localField":   "branch_id",
// 				"foreignField": "id",
// 				"as":           "branch",
// 			},
// 		},
// 		{
// 			"$unwind": "$branch",
// 		},
// 		{
// 			"$lookup": bson.M{
// 				"from":         "groups",
// 				"localField":   "group_id",
// 				"foreignField": "id",
// 				"as":           "group",
// 			},
// 		},
// 		{
// 			"$unwind": "$group",
// 		},
// 		{
// 			"$match": filter,
// 		},
// 		{
// 			"$project": bson.M{
// 				"_id":          0,
// 				"id":           1,
// 				"student_id":   1,
// 				"first_name":   1,
// 				"last_name":    1,
// 				"phone_number": 1,
// 				"teacher":      1,
// 				"coordinator":  1,
// 				"branchName":   "$branch.name",
// 				"groupName":    "$group.name",
// 				"created_at":   1,
// 				"graduated_at": 1,
// 			},
// 		},
// 	}

// 	cursor, err := s.studentCollection.Aggregate(ctx, pipeline)
// 	if err == mongo.ErrNoDocuments {
// 		return &models.Student{}, nil
// 	} else if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	var results []models.Student
// 	for cursor.Next(ctx) {
// 		var student models.Student
// 		err = cursor.Decode(&student)
// 		if err != nil {
// 			return nil, err
// 		}
// 		results = append(results, student)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	resp := &models.Student{}
// 	if len(results) > 0 {
// 		resp = &results[0]
// 	}

// 	return resp, nil
// }

// func (s studentRepo) List(ctx context.Context, req *models.GetStudentListRequest) (res *models.GetStudentListResponse, err error) {
// 	var filter bson.M

// 	if req.Limit == 0 {
// 		req.Limit = 10
// 	}

// 	if req.BranchID != "" {
// 		filter = bson.M{"branch_id": req.BranchID}
// 	}

// 	if req.GroupID != "" {
// 		filter = bson.M{"group_id": req.GroupID}
// 	}

// 	opts := options.Find().SetSort(bson.M{
// 		"created_at": -1,
// 	},
// 	).SetSkip(req.Offset).SetLimit(req.Limit)

// 	count, err := s.studentCollection.CountDocuments(context.Background(), filter)
// 	if err != nil {
// 		fmt.Printf("count Documents error: %v", err.Error())
// 		return nil, err
// 	}

// 	rows, err := s.studentCollection.Find(
// 		context.Background(),
// 		filter,
// 		opts,
// 	)
// 	var students = []*models.Student{}

// 	if err != nil {
// 		fmt.Printf("Find Student List err: %v", err.Error())
// 		return nil, err
// 	}

// 	if err = rows.All(context.Background(), &students); err != nil {
// 		fmt.Printf("get all student error: %s", err.Error())
// 		return nil, err
// 	}

// 	return &models.GetStudentListResponse{
// 		Count:    count,
// 		Students: students,
// 	}, nil
// }

// func (s studentRepo) Update(ctx context.Context, req *models.StudentUpdate) (res *models.Response, err error) {
// 	updateReq := bson.M{
// 		"$set": bson.M{
// 			"student_id":   req.StudentID,
// 			"first_name":   req.FirstName,
// 			"last_name":    req.LastName,
// 			"phone_number": req.PhoneNumber,
// 			"teacher":      req.Teacher,
// 			"coordinator":  req.Coordinator,
// 			"branch_id":    req.BranchID,
// 			"group_id":     req.GroupID,
// 			"updated_at":   time.Now().Format("02.01.2006 15:04"),
// 			"graduated_at": req.GraduatedAt,
// 		},
// 	}
// 	_, err = s.studentCollection.UpdateOne(ctx, bson.M{"id": req.ID}, updateReq)
// 	if err != nil {
// 		fmt.Printf("Error while updating student: %v", err.Error())
// 		return nil, err
// 	}

// 	return &models.Response{
// 		Text: "Succesfully Updated !",
// 	}, nil
// }

// func (s studentRepo) Delete(ctx context.Context, id string) (res *models.Response, err error) {
// 	_, err = s.studentCollection.DeleteOne(ctx, bson.M{"id": id})
// 	if err != nil {
// 		fmt.Printf("Delete Student err: %v", err)
// 		return nil, err
// 	}

// 	return &models.Response{
// 		Text: "Successfully Deleted !",
// 	}, nil
// }

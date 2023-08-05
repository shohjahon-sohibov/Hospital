package storage

import (
	"context"
	"freelance/clinic_queue/models"
)

type StorageI interface {
	Auth() AuthI
	User() UserI
	Diagnosis() DiagnosisI
	Hospital() HospitalI
	Service() ServiceI
	Queue() QueueI
}

type AuthI interface {
	SignUp(ctx context.Context, req *models.SignUp) (res *models.Token, err error)
	Login(ctx context.Context, req *models.Login) (res *models.Token, err error)
}

type UserI interface {
	CreateUser(ctx context.Context, req *models.CreateUser) (*models.Id, error)
	GetSingleUser(ctx context.Context, id string) (res *models.User, err error)
	GetListUser(ctx context.Context, req *models.GetUsersListReq) (res *models.GetUsersListRes, err error)
	UpdateUser(ctx context.Context, req *models.UpdateUser) (res *models.Message, err error)
	DeleteUser(ctx context.Context, id string) (res *models.Message, err error)
}

type DiagnosisI interface {
	CreateDiagnosis(ctx context.Context, req *models.CreateDiagnosis) (res *models.Message, err error)
	GetSingleDiagnosis(ctx context.Context, id string) (res *models.Diagnosis, err error)
	GetListDiagnosis(ctx context.Context, req *models.GetListDiagnosisReq) (res *models.GetListDiagnosisRes, err error)
	UpdateDiagnosis(ctx context.Context, req *models.UpdateDiagnosis) (res *models.Message, err error)
	DeleteDiagnosis(ctx context.Context, id string) (res *models.Message, err error)
}

type HospitalI interface {
	CreateHospital(ctx context.Context, req *models.CreateHospital) (res *models.Message, err error)
	GetSingleHospital(ctx context.Context, id string) (res *models.Hospital, err error)
	GetListHospital(ctx context.Context, req *models.GetListHospitalReq) (res *models.GetListHospitalRes, err error)
	UpdateHospital(ctx context.Context, req *models.UpdateHospital) (res *models.Message, err error)
	DeleteHospital(ctx context.Context, id string) (res *models.Message, err error)
}

type ServiceI interface {
	CreateService(ctx context.Context, req *models.CreateService) (res *models.Message, err error)
	GetListService(ctx context.Context, req *models.GetListServiceReq) (res *models.GetListServiceRes, err error)
	UpdateService(ctx context.Context, req *models.UpdateService) (res *models.Message, err error)
	DeleteService(ctx context.Context, id string) (res *models.Message, err error)
}

type QueueI interface {
	CreateQueue(ctx context.Context, req *models.CreateQueue) (res *models.Ticket, err error)
	GetListQueue(ctx context.Context, req *models.GetListQueueReq) (res *models.GetListQueueRes, err error)
	ChangeStatusQueue(ctx context.Context, req *models.ChangeStatusQueue) (res *models.Message, err error)
	DeleteQueue(ctx context.Context, id string) (res *models.Message, err error)
}
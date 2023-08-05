package mongodb

import (
	"freelance/clinic_queue/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type storagePg struct {
	authRepo      storage.AuthI
	userRepo      storage.UserI
	diagnosisRepo storage.DiagnosisI
	hospitalRepo  storage.HospitalI
	serviceRepo storage.ServiceI
	queueRepo storage.QueueI
}

func NewStoragePg(db *mongo.Database) storage.StorageI {
	return &storagePg{
		authRepo:      NewAuthRepo(db),
		userRepo:      NewUserRepo(db),
		diagnosisRepo: NewDiagnosisRepo(db),
		hospitalRepo:  NewHospitalRepo(db),
		serviceRepo: NewServiceRepo(db),
		queueRepo: NewQueueRepo(db),
	}
}

func (s *storagePg) Auth() storage.AuthI {
	return s.authRepo
}

func (s *storagePg) User() storage.UserI {
	return s.userRepo
}

func (s *storagePg) Diagnosis() storage.DiagnosisI {
	return s.diagnosisRepo
}

func (s *storagePg) Hospital() storage.HospitalI {
	return s.hospitalRepo
}

func (s *storagePg) Service() storage.ServiceI {
	return s.serviceRepo
}

func (s *storagePg) Queue() storage.QueueI {
	return s.queueRepo
}
package status

import "Assignment3/app/model"

type StatusUsecase interface {
	UpdateStatus() (model.StatusBahaya, error)
}

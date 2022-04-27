package grpc

import (
	"calendar/internal/models"
	"calendar/internal/repository"
	"calendar/pkg/api"
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type SHandler struct {
	api.UnimplementedGHandlersServer
	Storage repository.Db
}

func (s SHandler) UpdateUser(c context.Context, gr *api.GRequest) (*api.GResponse, error) {

	var user models.User

	req := []byte(gr.GetX())
	err := json.Unmarshal(req, &user)
	if err != nil {
		log.Println("JSON parse: ", err)
	}

	err = s.Storage.Db.UpdateUserTimezoneById(string(user.Id), user.Timezone)
	if err != nil {
		log.Println("SQL: ", err)
	}
	resp := api.GResponse{
		Result: "Successfully saved",
	}
	return &resp, status.Errorf(codes.OK, "ok")
}

func (s SHandler) GetEvents(c context.Context, gr *api.GRequest) (*api.GResponse, error) {
	var user models.User

	req := []byte(gr.GetX())
	err := json.Unmarshal(req, &user)
	if err != nil {
		log.Println("JSON parse: ", err)
	}

	user, err = s.Storage.Db.GetUserById(string(user.Id))
	if err != nil {
		log.Println("SQL user: ", err)
	}

	events, err := s.Storage.Db.GetEvents()
	if err != nil {
		log.Println("SQL event: ", err)
	}

	var mSlice = make([]models.EventResult, 0)
	for _, event := range events {
		e := event.MakeApiData(user)
		mSlice = append(mSlice, e)
	}

	js, err := json.Marshal(mSlice)
	resp := api.GResponse{
		Result: string(js),
	}
	return &resp, status.Errorf(codes.OK, "ok")
}

func (s SHandler) CreateEvent(c context.Context, gr *api.GRequest) (*api.GResponse, error) {
	var user models.User

	req := []byte(gr.GetX())
	err := json.Unmarshal(req, &user)
	if err != nil {
		log.Println("JSON parse: ", err)
	}

	user, err = s.Storage.Db.GetUserById(string(user.Id))
	if err != nil {
		log.Println("SQL user: ", err)
	}

	events, err := s.Storage.Db.GetEvents()
	if err != nil {
		log.Println("SQL event: ", err)
	}

	var mSlice = make([]models.EventResult, 0)
	for _, event := range events {
		e := event.MakeApiData(user)
		mSlice = append(mSlice, e)
	}

	js, err := json.Marshal(mSlice)
	resp := api.GResponse{
		Result: string(js),
	}
	return &resp, status.Errorf(codes.OK, "ok")
}

func (s SHandler) GetEventById(c context.Context, gr *api.GRequest) (*api.GResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventById not implemented")
}

func (s SHandler) UpdateEvent(c context.Context, gr *api.GRequest) (*api.GResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}

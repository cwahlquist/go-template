package handler

import (
	"context"
	"log"

	pb "go-template/api/go"
	m "go-template/model"
	s "go-template/service"
)

type Handler struct {
	service s.Service
}

func NewHandler(service *s.Service) *Handler {
	return &Handler{
		service: *service,
	}
}

func (h *Handler) Create(ctx context.Context, req *pb.ServiceCommand) (*pb.ServiceCommand, error) {
	log.Println("Create: ", *req)

	model := protoToModel(req)

	todo, err := h.service.Create(model)
	if err != nil {
		return nil, err
	}

	return modelToProto(todo), nil
}

func (h *Handler) Read(ctx context.Context, req *pb.Id) (*pb.ServiceCommand, error) {
	log.Println("Read todo: ", *req)

	model, err := h.service.Read(req.Id)
	if err != nil {
		return nil, err
	}

	return modelToProto(model), nil
}

// Private conversion methods
func protoToModel(pb *pb.ServiceCommand) *m.ServiceCommand {
	return &m.Todo{
		Id:        pb.Id,
		Name:      pb.Name,
		Completed: pb.Completed,
	}
}

func modelToProto(model *m.ServiceCommand) *pb.ServiceCommand {
	return &pb.Todo{
		Id:        model.Id,
		Name:      model.Name,
		Completed: model.Completed,
	}
}

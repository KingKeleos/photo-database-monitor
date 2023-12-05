package server

import (
	"context"
	"fmt"

	"github.com/kingkeleos/database-monitor/graphite"
	DatabaseMonitor "github.com/kingkeleos/database-monitor/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	DatabaseMonitor.UnimplementedDatabaseMonitorServer
}

func createMap(name string, content float64) map[string]float64 {
	return map[string]float64{
		name: content,
	}
}

func (h Handler) UpdateProjects(ctx context.Context, request *DatabaseMonitor.UpdateCountReqest) (*emptypb.Empty, error) {
	count := float64(*request.CurrentCount)

	metric := createMap("project_count", count)

	err := graphite.Client.SendData(metric)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h Handler) UpdatePeople(ctx context.Context, request *DatabaseMonitor.UpdateCountReqest) (*emptypb.Empty, error) {
	count := float64(*request.CurrentCount)

	metric := createMap("people_count", count)

	err := graphite.Client.SendData(metric)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h Handler) UpdateParticipantsToProject(ctx context.Context, request *DatabaseMonitor.CountToProjectRequest) (*emptypb.Empty, error) {
	count := float64(*request.Count)

	label := fmt.Sprintf("%s_participants", *request.Name)
	metric := createMap(label, count)

	err := graphite.Client.SendData(metric)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h Handler) UpdatePostsToProject(ctx context.Context, request *DatabaseMonitor.CountToProjectRequest) (*emptypb.Empty, error) {
	count := float64(*request.Count)

	label := fmt.Sprintf("%s_posts", *request.Name)
	metric := createMap(label, count)

	err := graphite.Client.SendData(metric)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h Handler) UpdateSocials(ctx context.Context, request *DatabaseMonitor.UpdateCountReqest) (*emptypb.Empty, error) {
	count := float64(*request.CurrentCount)

	metric := createMap("socials_count", count)

	err := graphite.Client.SendData(metric)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h Handler) NewProjects(ctx context.Context, request *DatabaseMonitor.UpdateCountReqest) (*emptypb.Empty, error) {
	count := float64(*request.CurrentCount)

	metric := createMap("new_projects", count)

	err := graphite.Client.SendData(metric)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h Handler) ActiveProjects(ctx context.Context, request *DatabaseMonitor.UpdateCountReqest) (*emptypb.Empty, error) {
	count := float64(*request.CurrentCount)

	metric := createMap("active_projects", count)

	err := graphite.Client.SendData(metric)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h Handler) FinishedProjects(ctx context.Context, request *DatabaseMonitor.UpdateCountReqest) (*emptypb.Empty, error) {
	count := float64(*request.CurrentCount)

	metric := createMap("finished_projects", count)

	err := graphite.Client.SendData(metric)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

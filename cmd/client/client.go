package client

import (
	"context"

	DatabaseMonitor "github.com/KingKeleos/photo-database-monitor/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MonitorClient struct {
	Client DatabaseMonitor.DatabaseMonitorClient
}

func NewClient(address string) (MonitorClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return MonitorClient{}, err
	}

	client := DatabaseMonitor.NewDatabaseMonitorClient(conn)

	return MonitorClient{Client: client}, nil
}

func (client *MonitorClient) UpdateProjects(in uint32) error {
	request := DatabaseMonitor.UpdateCountReqest{
		CurrentCount: &in,
	}

	_, err := DatabaseMonitor.DatabaseMonitorClient.UpdateProjects(client.Client, context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}

func (client *MonitorClient) UpdatePeople(in uint32) error {
	request := DatabaseMonitor.UpdateCountReqest{
		CurrentCount: &in,
	}

	_, err := DatabaseMonitor.DatabaseMonitorClient.UpdatePeople(client.Client, context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}

func (client *MonitorClient) UpdateParticipantsToProject(projectName string, in uint32, id string) error {
	request := DatabaseMonitor.CountToIDRequest{
		ID:    &id,
		Name:  &projectName,
		Count: &in,
	}

	_, err := DatabaseMonitor.DatabaseMonitorClient.UpdateParticipantsToProject(client.Client, context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}

func (client *MonitorClient) UpdatePostsToProject(projectName string, in uint32, id string) error {
	request := DatabaseMonitor.CountToIDRequest{
		ID:    &id,
		Name:  &projectName,
		Count: &in,
	}

	_, err := DatabaseMonitor.DatabaseMonitorClient.UpdatePostsToProject(client.Client, context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}

func (client *MonitorClient) UpdateSocials(projectName string, in uint32, id string) error {
	request := DatabaseMonitor.CountToIDRequest{
		ID:    &id,
		Name:  &projectName,
		Count: &in,
	}

	_, err := DatabaseMonitor.DatabaseMonitorClient.UpdatePostsToProject(client.Client, context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}

func (client *MonitorClient) NewProjects(in uint32) error {
	request := DatabaseMonitor.UpdateCountReqest{
		CurrentCount: &in,
	}

	_, err := DatabaseMonitor.DatabaseMonitorClient.NewProjects(client.Client, context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}

func (client *MonitorClient) ActiveProjects(in uint32) error {
	request := DatabaseMonitor.UpdateCountReqest{
		CurrentCount: &in,
	}

	_, err := DatabaseMonitor.DatabaseMonitorClient.ActiveProjects(client.Client, context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}

func (client *MonitorClient) FinishedProjects(in uint32) error {
	request := DatabaseMonitor.UpdateCountReqest{
		CurrentCount: &in,
	}

	_, err := DatabaseMonitor.DatabaseMonitorClient.FinishedProjects(client.Client, context.Background(), &request)
	if err != nil {
		return err
	}
	return nil
}

package grpc

import (
	"context"

	pb "github.com/LuhTonkaYeat/GoHomeworks/hw2-4/services/gateway/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SubscribeClient struct {
	client pb.SubscribeServiceClient
	conn   *grpc.ClientConn
}

func NewSubscribeClient(addr string) (*SubscribeClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewSubscribeServiceClient(conn)

	return &SubscribeClient{
		client: client,
		conn:   conn,
	}, nil
}

func (c *SubscribeClient) Close() error {
	return c.conn.Close()
}

func (c *SubscribeClient) CreateSubscription(ctx context.Context, owner, repo, userID string) error {
	req := &pb.CreateSubscriptionRequest{
		Owner:  owner,
		Repo:   repo,
		UserId: userID,
	}

	_, err := c.client.CreateSubscription(ctx, req)
	return err
}

func (c *SubscribeClient) DeleteSubscription(ctx context.Context, owner, repo, userID string) error {
	req := &pb.DeleteSubscriptionRequest{
		Owner:  owner,
		Repo:   repo,
		UserId: userID,
	}

	_, err := c.client.DeleteSubscription(ctx, req)
	return err
}

func (c *SubscribeClient) GetSubscriptions(ctx context.Context, userID string) ([]*pb.Repository, error) {
	req := &pb.Empty{}

	resp, err := c.client.GetSubscriptions(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Subscriptions, nil
}

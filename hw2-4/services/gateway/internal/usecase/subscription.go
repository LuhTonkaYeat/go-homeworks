package usecase

import (
	"context"

	pb "github.com/LuhTonkaYeat/GoHomeworks/hw2-4/services/gateway/api/proto"
	"github.com/LuhTonkaYeat/GoHomeworks/hw2-4/services/gateway/internal/adapter/grpc"
)

type SubscriptionUseCase interface {
	CreateSubscription(ctx context.Context, owner, repo, userID string) error
	DeleteSubscription(ctx context.Context, owner, repo, userID string) error
	GetSubscriptions(ctx context.Context, userID string) ([]*pb.Repository, error)
}

type subscriptionUseCase struct {
	subscribeClient *grpc.SubscribeClient
}

func NewSubscriptionUseCase(subscribeClient *grpc.SubscribeClient) SubscriptionUseCase {
	return &subscriptionUseCase{
		subscribeClient: subscribeClient,
	}
}

func (uc *subscriptionUseCase) CreateSubscription(ctx context.Context, owner, repo, userID string) error {
	return uc.subscribeClient.CreateSubscription(ctx, owner, repo, userID)
}

func (uc *subscriptionUseCase) DeleteSubscription(ctx context.Context, owner, repo, userID string) error {
	return uc.subscribeClient.DeleteSubscription(ctx, owner, repo, userID)
}

func (uc *subscriptionUseCase) GetSubscriptions(ctx context.Context, userID string) ([]*pb.Repository, error) {
	return uc.subscribeClient.GetSubscriptions(ctx, userID)
}

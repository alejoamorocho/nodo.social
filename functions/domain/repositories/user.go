package repositories

import (
    "context"
    "github.com/nodo-social/functions/domain/models"
)

type UserRepository interface {
    Create(ctx context.Context, user *models.User) error
    Update(ctx context.Context, user *models.User) error
    Get(ctx context.Context, id string) (*models.User, error)
    Delete(ctx context.Context, id string) error
    FollowNode(ctx context.Context, userID, nodeID string) error
    UnfollowNode(ctx context.Context, userID, nodeID string) error
    FollowUser(ctx context.Context, followerID, followedID string) error
    UnfollowUser(ctx context.Context, followerID, followedID string) error
    GetFollowers(ctx context.Context, userID string) ([]string, error)
    GetFollowing(ctx context.Context, userID string) ([]string, error)
    AddAchievement(ctx context.Context, userID string, achievement models.Achievement) error
    GetAchievements(ctx context.Context, userID string) ([]models.Achievement, error)
}

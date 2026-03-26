package services

import (
    "context"

    "github.com/MiharB-E/InvCasa/internal/models"
    "github.com/MiharB-E/InvCasa/internal/repositories"
)

type GroupService struct {
    repos *repositories.Repositories
}

func (s *GroupService) Create(ctx context.Context, name string) (models.Group, error) {
    return s.repos.Groups.Create(ctx, models.Group{
        Name:       name,
        InviteCode: generateInviteCode(),
    })
}

func (s *GroupService) Join(ctx context.Context, userID int64, inviteCode string) (models.Group, error) {
    group, err := s.repos.Groups.GetByInviteCode(ctx, inviteCode)
    if err != nil {
        return models.Group{}, err
    }
    if err := s.repos.Users.UpdateGroup(ctx, userID, group.ID); err != nil {
        return models.Group{}, err
    }
    return group, nil
}

func (s *GroupService) Get(ctx context.Context, id int64) (models.Group, error) {
    return s.repos.Groups.GetByID(ctx, id)
}
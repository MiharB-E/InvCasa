package services

import (
    "context"
    "errors"
    "time"

    "github.com/MiharB-E/InvCasa/internal/models"
    "github.com/MiharB-E/InvCasa/internal/repositories"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    repos     *repositories.Repositories
    jwtSecret string
}

type RegisterInput struct {
    Email       string
    Password    string
    Name        string
    InviteCode  string
    GroupName   string
}

func (s *AuthService) Register(ctx context.Context, input RegisterInput) (models.User, string, error) {
    var group models.Group
    var err error

    if input.InviteCode != "" {
        group, err = s.repos.Groups.GetByInviteCode(ctx, input.InviteCode)
        if err != nil {
            return models.User{}, "", errors.New("invite code inválido")
        }
    } else {
        name := input.GroupName
        if name == "" {
            name = input.Name + "'s group"
        }
        group, err = s.repos.Groups.Create(ctx, models.Group{
            Name:       name,
            InviteCode: generateInviteCode(),
        })
        if err != nil {
            return models.User{}, "", err
        }
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        return models.User{}, "", err
    }

    user, err := s.repos.Users.Create(ctx, models.User{
        Email:   input.Email,
        Password: string(hash),
        Name:    input.Name,
        GroupID: group.ID,
    })
    if err != nil {
        return models.User{}, "", err
    }

    token, err := s.issueToken(user.ID, user.GroupID)
    if err != nil {
        return models.User{}, "", err
    }

    return user, token, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (models.User, string, error) {
    user, err := s.repos.Users.GetByEmail(ctx, email)
    if err != nil {
        return models.User{}, "", errors.New("credenciales inválidas")
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return models.User{}, "", errors.New("credenciales inválidas")
    }
    token, err := s.issueToken(user.ID, user.GroupID)
    if err != nil {
        return models.User{}, "", err
    }
    return user, token, nil
}

func (s *AuthService) issueToken(userID, groupID int64) (string, error) {
    claims := jwt.MapClaims{
        "user_id":  userID,
        "group_id": groupID,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtSecret))
}

func generateInviteCode() string {
    const letters = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
    b := make([]byte, 6)
    now := time.Now().UnixNano()
    for i := range b {
        now = (now*1664525 + 1013904223) % (1 << 31)
        b[i] = letters[now%int64(len(letters))]
    }
    return string(b)
}
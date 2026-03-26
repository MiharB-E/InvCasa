package middleware

import (
    "context"
    "errors"
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const (
    ctxUserID  ctxKey = "user_id"
    ctxGroupID ctxKey = "group_id"
)

func JWTAuth(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            auth := r.Header.Get("Authorization")
            if auth == "" {
                http.Error(w, "missing token", http.StatusUnauthorized)
                return
            }

            parts := strings.SplitN(auth, " ", 2)
            if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
                http.Error(w, "invalid token", http.StatusUnauthorized)
                return
            }

            token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
                if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, errors.New("invalid signing method")
                }
                return []byte(secret), nil
            })
            if err != nil || !token.Valid {
                http.Error(w, "invalid token", http.StatusUnauthorized)
                return
            }

            claims, ok := token.Claims.(jwt.MapClaims)
            if !ok {
                http.Error(w, "invalid token", http.StatusUnauthorized)
                return
            }

            userID, ok1 := claims["user_id"].(float64)
            groupID, ok2 := claims["group_id"].(float64)
            if !ok1 || !ok2 {
                http.Error(w, "invalid token", http.StatusUnauthorized)
                return
            }

            ctx := context.WithValue(r.Context(), ctxUserID, int64(userID))
            ctx = context.WithValue(ctx, ctxGroupID, int64(groupID))
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func GetUserID(ctx context.Context) int64 {
    if v, ok := ctx.Value(ctxUserID).(int64); ok {
        return v
    }
    return 0
}

func GetGroupID(ctx context.Context) int64 {
    if v, ok := ctx.Value(ctxGroupID).(int64); ok {
        return v
    }
    return 0
}
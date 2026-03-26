package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/MiharB-E/InvCasa/internal/config"
    "github.com/MiharB-E/InvCasa/internal/database"
    "github.com/MiharB-E/InvCasa/internal/handlers"
    "github.com/MiharB-E/InvCasa/internal/repositories"
    "github.com/MiharB-E/InvCasa/internal/services"
)

func main() {
    cfg := config.Load()

    ctx := context.Background()
    pool, err := database.NewPool(ctx, cfg.DBURL)
    if err != nil {
        log.Fatalf("db connection failed: %v", err)
    }
    defer pool.Close()

    repos := repositories.New(pool)
    srvs := services.New(repos, cfg.JWTSecret)

    router := handlers.NewRouter(srvs, cfg.JWTSecret)

    srv := &http.Server{
        Addr:              ":" + cfg.Port,
        Handler:           router,
        ReadHeaderTimeout: 5 * time.Second,
    }

    go func() {
        log.Printf("server listening on :%s", cfg.Port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %v", err)
        }
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Printf("shutdown error: %v", err)
    }
}
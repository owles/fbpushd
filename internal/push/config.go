package push

import (
	"context"
	"log/slog"
	"path/filepath"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

func SetupFirebase() (*firebase.App, context.Context, error) {
	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./keys.json")
	if err != nil {
		slog.Error("Firebase initialization", "message", err.Error())
		return nil, ctx, err
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		slog.Error("Firebase initialization", "message", err.Error())
		return nil, ctx, err
	}

	slog.Info("Firebase successfully initialized")
	return app, ctx, nil
}

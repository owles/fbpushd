package push

import (
	"context"

	"firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"log/slog"
)

// wrappedContext for message logging
func wrappedContext(token, title, body string) context.Context {
	ctx := context.WithValue(context.Background(), "method", "SendNotification")
	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "title", title)
	ctx = context.WithValue(ctx, "body", body)

	return ctx
}

// SendNotification sends a push notification
func SendNotification(token string, title, body string, data map[string]string, app *firebase.App) bool {
	ctx := wrappedContext(token, title, body)
	slog.InfoContext(ctx, "Start")

	client, err := app.Messaging(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error", "action", "app.Messaging", "message", err.Error())
		return false
	}

	message := &messaging.Message{
		Data: data,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: token,
	}

	_, err = client.Send(ctx, message)
	if err != nil {
		slog.ErrorContext(ctx, "Error", "action", "client.Send", "message", err.Error())
		return false
	}

	slog.InfoContext(ctx, "Sent")

	return true
}

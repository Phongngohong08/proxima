package account

import (
	"context"

	"proxima/app/user/internal/model/entity"
	"proxima/app/user/internal/service/email"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func Register(ctx context.Context) (id int, err error) {
	// Your existing registration logic
	id = 1

	// Create email service client
	emailService := email.New()

	// Send welcome email to user
	err = emailService.SendEmail(ctx,
		"ngohongphong18@gmail.com",
		"Welcome to Proxima",
		"Thank you for registering with Proxima!")

	if err != nil {
		// Log the error but don't fail the registration
		g.Log().Error(ctx, "Failed to send welcome email:", err)
	}

	return id, nil
}

func Login(ctx context.Context) (token string, err error) {
	return "I am token", nil
}

// Info get user info
func Info(ctx context.Context, token string) (user *entity.Users, err error) {
	user = &entity.Users{
		Id:        1,
		Username:  "oldme",
		Password:  "123456",
		Email:     "tyyn1022@gmail.com",
		CreatedAt: gtime.New("2024-12-05 22:00:00"),
		UpdatedAt: gtime.New("2024-12-05 22:00:00"),
	}

	// Example: Notify about account access
	emailService := email.New()
	err = emailService.SendEmail(ctx,
		user.Email,
		"Account Access Notification",
		"Your account was accessed on "+gtime.Now().String())

	if err != nil {
		g.Log().Warning(ctx, "Failed to send account access notification:", err)
		// Continue despite email error
	}

	return user, nil
}

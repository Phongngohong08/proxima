package email

import (
	"context"
	"fmt"

	v1 "proxima/app/email/api/email/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Service provides email functionality through the email service
type Service struct {
	emailClient v1.EmailClient
}

// New creates a new email service client
func New() *Service {
	// Simplest possible connection - no TLS, no timeouts, no blocking
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(),
	}

	conn, err := grpc.Dial("localhost:32003", opts...)

	if err != nil {
		fmt.Printf("Failed to connect to email service: %v\n", err)
		return &Service{emailClient: nil}
	}

	return &Service{
		emailClient: v1.NewEmailClient(conn),
	}
}

// SendEmail sends an email using the email service
func (s *Service) SendEmail(ctx context.Context, to, subject, body string) error {
	// Check if client was successfully created
	if s.emailClient == nil {
		return fmt.Errorf("email client not initialized")
	}

	res, err := s.emailClient.SendEmail(ctx, &v1.SendEmailReq{
		To:      to,
		Subject: subject,
		Body:    body,
	})

	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return err
	}

	if !res.Success {
		return fmt.Errorf("email service returned failure: %s", res.Message)
	}

	return nil
}

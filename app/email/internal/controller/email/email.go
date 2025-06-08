package email

import (
	"context"
	"fmt"
	"net/smtp"
	v1 "proxima/app/email/api/email/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedEmailServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterEmailServer(s.Server, &Controller{})
}

func (*Controller) SendEmail(ctx context.Context, req *v1.SendEmailReq) (res *v1.SendEmailRes, err error) {
	// Thông tin email người gửi (hardcoded)
	from := "phongngohong18@gmail.com"
	password := "your-email-password" // Thay thế bằng mật khẩu email của bạn

	// Cấu hình SMTP server Gmail
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Tạo auth
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Tạo nội dung email với header đầy đủ
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", req.To, req.Subject, req.Body)

	// Gửi email đơn giản sử dụng smtp.SendMail (tự động xử lý TLS)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{req.To}, []byte(message))

	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return &v1.SendEmailRes{
			Success: false,
			Message: fmt.Sprintf("Failed to send email: %v", err),
		}, nil
	}

	fmt.Printf("Email sent successfully to: %s\n", req.To)
	return &v1.SendEmailRes{
		Success: true,
		Message: "Email sent successfully!",
	}, nil
}

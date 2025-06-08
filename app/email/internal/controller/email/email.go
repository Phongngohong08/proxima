package email

import (
	"context"
	"crypto/tls"
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
	password := "your_password_or_app_specific_password"

	// Thông tin email người nhận (hardcoded)
	to := []string{
		"ngohongphong18@gmail.com",
	}

	// Cấu hình SMTP server
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Nội dung email
	message := []byte(req.Subject + "\n" + req.Body)

	// Xác thực với SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Gửi email (có thể cần tls.Config nếu dùng port 465)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	// Kết nối và gửi email
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		fmt.Println("Error connecting to SMTP server:", err)
		return
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		fmt.Println("Error creating SMTP client:", err)
		return
	}

	// Xác thực
	if err = client.Auth(auth); err != nil {
		fmt.Println("Error authenticating:", err)
		return
	}

	// Thiết lập người gửi và người nhận
	if err = client.Mail(from); err != nil {
		fmt.Println("Error setting sender:", err)
		return
	}

	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			fmt.Println("Error setting recipient:", err)
			return
		}
	}

	// Gửi nội dung email
	w, err := client.Data()
	if err != nil {
		fmt.Println("Error preparing to send data:", err)
		return
	}

	_, err = w.Write(message)
	if err != nil {
		fmt.Println("Error writing message:", err)
		return
	}

	err = w.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}

	// Đóng kết nối
	client.Quit()

	return &v1.SendEmailRes{
		Success: true,
		Message: "Email sent successfully!",
	}, nil
}

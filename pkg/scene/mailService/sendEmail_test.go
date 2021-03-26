package mailService

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	to := []string{"2483777000@qq.com"}
	email := new(SendEmail)
	email.SendEmail(to, "mySubject", "This is body!")
}

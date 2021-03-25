package mailService

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net"
	"net/smtp"
	"strings"

	"github.com/chanprogo/somemodule/internal/smsrpcsvrpkg/constant"
)

type SendEmail struct {
}

func (s *SendEmail) SendEmail(toAddr []string, subject, body string) (int, error) {

	auth := smtp.PlainAuth("", constant.USER_NAME, constant.PASSWORD, constant.HOST)

	nickname := constant.NICKNAME
	user := constant.USER_NAME

	content_type := "Content-Type: text/plain; charset=UTF-8"

	subject = fmt.Sprintf("=?UTF-8?B?%s?=", base64.StdEncoding.EncodeToString([]byte(subject)))

	// content_type := "Content-Type: text/plain; charset=UTF-8"

	msg := []byte("To: " + strings.Join(toAddr, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)

	err := SendMailUsingTLS(constant.HOST+constant.PORT, auth, user, toAddr, msg)

	fmt.Println("2 email 发送中.....")

	if err != nil {
		fmt.Printf("3 send mail error: %v \n\n", err)
		return 1, err
	} else {
		fmt.Println("send email success!")
		return 0, nil
	}

}

func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		// log.Panicln("Dialing Error:", err)
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr) // 分解主机端口字符串
	return smtp.NewClient(conn, host)
}

// 参考 net/smtp 的 func SendMail()
// 使用 net.dial 连接 tls(ssl) 端口时,smtp.NewClient() 会卡住且不提示 err
// len(to)>1 时,to[1] 开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
	fmt.Println(addr)
	c, err := dial(addr)
	if err != nil {
		// log.Logger.Error("Create smpt client error:", err)
		fmt.Printf("1 Create smpt client error: %v \n", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				// log.Logger.Error("Error during AUTH", err)
				fmt.Printf("Error during AUTH: %v \n", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

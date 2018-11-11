package util

import "gopkg.in/gomail.v2"

func SendEmail(subject string, to []string, body string) (err error) {
	email := gomail.NewMessage()
	email.SetHeaders(map[string][]string{
		"From":    {email.FormatAddress("450512013@qq.com", "monitor")},
		"To":      to,
		"Subject": {subject},
	})
	email.SetBody("text/html", body)

	dialer := gomail.NewDialer("smtp.qq.com", 587, "450512013@qq.com", "hdaqvifydcfabjeb")
	return dialer.DialAndSend(email)
}

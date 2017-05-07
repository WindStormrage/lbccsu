package service

import ("gomail"
	"math/rand"
)

func SendEmail(email string,verify string) (error) {

	m:=gomail.NewMessage()
	m.SetAddressHeader("From","269812428@qq.com","长沙网络中心实验室")
	m.SetHeader("To",m.FormatAddress(email,"xxx"),
	)
	m.SetHeader("Subject","验证码")
	m.SetBody("text/html",verify)

	d:=gomail.NewPlainDialer("smtp.qq.com",465,"269812428@qq.com","uyniahqxjajpbjac")
	if err:=d.DialAndSend(m);err!=nil{
		panic(err)
		return err

	}
	return nil
}

func RandInt64(min, max int64) int64 {
	if min >= max || min==0 || max==0{
		return max
	}
	return rand.Int63n(max-min)+min
}
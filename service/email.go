package service

import ("gomail"
	"math/rand"
	"github.com/astaxie/beego"
)

func SendEmail(email string,verify string)  {
	beego.Debug("this is send email")
	m:=gomail.NewMessage()
	m.SetAddressHeader("From","269812428@qq.com","长沙网络中心实验室")
	m.SetHeader("To",m.FormatAddress("2455425131@qq.com","xxx"),
	)
	m.SetHeader("Subject","验证码")
	m.SetBody("text/html",verify)

	d:=gomail.NewPlainDialer("smtp.qq.com",465,"269812428@qq.com","xhx971129")
	if err:=d.DialAndSend(m);err!=nil{
		panic(err)
	}
}

func RandInt64(min, max int64) int64 {
	if min >= max || min==0 || max==0{
		return max
	}
	return rand.Int63n(max-min)+min
}
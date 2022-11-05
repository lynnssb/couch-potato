/**
 * @author:       wangxuebing
 * @fileName:     emailutil.go
 * @date:         2022/11/1 11:48
 * @description:
 */

package emailutil

import (
	"couch-potato/common/utils/characterutil"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"
)

type EmailTypeArg string

const (
	PLAINTEXT EmailTypeArg = "PLAINTEXT" //纯文字
	HTML      EmailTypeArg = "HTML"      //HTML
)

// EmailReqParams 邮件发送参数
type EmailReqParams struct {
	Subject      string       //邮件主题
	TypeArg      EmailTypeArg //邮件类型(纯文字|HTML)
	Content      string       //邮件内容(纯文字|HTML)
	Files        []string     //附件
	ReceiveEmail []string     //接收者邮箱号码
	CcEmail      []string     //抄送邮箱号码
	BccEmail     []string     //秘密抄送邮箱号码
	EmailAddr    string       //邮箱服务器
	Port         string       //邮箱服务器端口
	Host         string       //邮箱服务器地址
	UserName     string       //用户名<发送者用户邮箱地址>
	Password     string       //授权码
}

func SendEmail(in EmailReqParams) error {
	pool, err := email.NewPool(characterutil.StitchingBufStr(in.EmailAddr, in.Port), 5,
		smtp.PlainAuth("", in.UserName, in.Password, in.Host))
	if err != nil {
		return err
	}

	em := &email.Email{
		From:    fmt.Sprintf("[couch-potato]<%s>", in.UserName), //发送者
		To:      in.ReceiveEmail,                                //接收者
		Subject: in.Subject,                                     //主题
	}

	if in.TypeArg == PLAINTEXT {
		em.Text = []byte(in.Content)
	} else if in.TypeArg == HTML {
		em.HTML = []byte(in.Content)
	}

	if len(in.Files) > 0 {
		for _, file := range in.Files {
			em.AttachFile(file)
		}
	}

	err = pool.Send(em, 5*time.Second)
	if err != nil {
		return err
	}
	return nil
}

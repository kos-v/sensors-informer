package channel

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
	"net/smtp"
	"strings"
)

type SmtpChannel struct {
	Opts             config.SmtpChannelOpts
	MessageFormatter message.Formatter
}

func (ch *SmtpChannel) GetName() string {
	return "smtp"
}

func (ch *SmtpChannel) IsEnable() bool {
	return ch.Opts.Enable
}

func (ch *SmtpChannel) Send(r report.Report) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", ch.Opts.From.Name, ch.Opts.From.Email)
	e.To = ch.Opts.To
	e.Subject = ch.MessageFormatter.FormatTitle(&r)
	e.HTML = []byte(ch.formatBody(r))

	hp := fmt.Sprintf("%s:%d", ch.Opts.Host, ch.Opts.Port)
	auth := smtp.PlainAuth("", ch.Opts.Username, ch.Opts.Password, ch.Opts.Host)
	tlsConf := &tls.Config{
		InsecureSkipVerify: ch.Opts.TLS.InsecureSkipVerify,
		ServerName:         ch.Opts.TLS.ServerName,
	}

	switch ch.Opts.Encryption {
	case "tls":
		return e.SendWithTLS(hp, auth, tlsConf)
	case "starttls":
		return e.SendWithStartTLS(hp, auth, tlsConf)
	}

	return e.Send(hp, auth)
}

func (ch *SmtpChannel) formatBody(r report.Report) string {
	return strings.Join(ch.MessageFormatter.FormatBodyRows(&r), "<br>")
}

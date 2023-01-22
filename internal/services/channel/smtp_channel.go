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

type SmtpChannelService struct {
	opts             config.SmtpChannelOpts
	messageFormatter message.Formatter
}

func (ch *SmtpChannelService) GetName() string {
	return "smtp"
}

func (ch *SmtpChannelService) IsEnable() bool {
	return ch.opts.Enable
}

func (ch *SmtpChannelService) Send(r report.Report) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", ch.opts.From.Name, ch.opts.From.Email)
	e.To = ch.opts.To
	e.Subject = ch.messageFormatter.FormatTitle(&r)
	e.HTML = []byte(ch.formatBody(r))

	hp := fmt.Sprintf("%s:%d", ch.opts.Host, ch.opts.Port)
	auth := smtp.PlainAuth("", ch.opts.Username, ch.opts.Password, ch.opts.Host)
	tlsConf := &tls.Config{
		InsecureSkipVerify: ch.opts.TLS.InsecureSkipVerify,
		ServerName:         ch.opts.TLS.ServerName,
	}

	switch ch.opts.Encryption {
	case "tls":
		return e.SendWithTLS(hp, auth, tlsConf)
	case "starttls":
		return e.SendWithStartTLS(hp, auth, tlsConf)
	}

	return e.Send(hp, auth)
}

func (ch *SmtpChannelService) formatBody(r report.Report) string {
	return strings.Join(ch.messageFormatter.FormatBodyRows(&r), "<br>")
}

func NewSmtpChannelService(
	opts config.SmtpChannelOpts,
	msgFormatter message.Formatter,
) *SmtpChannelService {
	return &SmtpChannelService{opts: opts, messageFormatter: msgFormatter}
}

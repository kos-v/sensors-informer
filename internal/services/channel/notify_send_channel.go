package channel

import (
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type NotifySendChannelService struct {
	opts             config.NotifySendChannelOpts
	messageFormatter message.Formatter
}

func (ch *NotifySendChannelService) GetName() string {
	return "notifySend"
}

func (ch *NotifySendChannelService) IsEnable() bool {
	return ch.opts.Enable
}

func (ch *NotifySendChannelService) Send(r report.Report) error {
	args := []string{
		"-u",
		string(ch.opts.Urgency),
		"-t",
		strconv.Itoa(ch.opts.ExpireTime),
		"-a",
		"sensors-informer",
	}

	if ch.opts.Hint != "" {
		args = append(args, "-h", ch.opts.Hint)
	}

	args = append(args, ch.formatTitle(&r), ch.formatBody(&r))

	os.Setenv("DISPLAY", ":0.0")
	cmd := exec.Command(ch.opts.Command, args...)
	_, err := cmd.Output()

	return err
}

func (ch *NotifySendChannelService) formatTitle(r *report.Report) string {
	title := ch.messageFormatter.FormatTitle(r)
	if title == "" {
		return "Critical temperature readings"
	}
	return title
}

func (ch *NotifySendChannelService) formatBody(r *report.Report) string {
	return strings.Join(ch.messageFormatter.FormatBodyRows(r), "\n")
}

func NewNotifySendChannelService(
	opts config.NotifySendChannelOpts,
	msgFormatter message.Formatter,
) *NotifySendChannelService {
	return &NotifySendChannelService{opts: opts, messageFormatter: msgFormatter}
}

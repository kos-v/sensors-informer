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

type NotifySendChannel struct {
	Opts             config.NotifySendChannelOpts
	MessageFormatter message.Formatter
}

func (ch *NotifySendChannel) IsEnable() bool {
	return ch.Opts.Enable
}

func (ch *NotifySendChannel) Send(r report.Report) error {
	args := []string{
		"-u",
		string(ch.Opts.Urgency),
		"-t",
		strconv.Itoa(ch.Opts.ExpireTime),
		"-a",
		"sensors-informer",
	}

	if ch.Opts.Hint != "" {
		args = append(args, "-h", ch.Opts.Hint)
	}

	args = append(args, ch.formatTitle(&r), ch.formatBody(&r))

	os.Setenv("DISPLAY", ":0.0")
	cmd := exec.Command(ch.Opts.Command, args...)
	_, err := cmd.Output()

	return err
}

func (ch *NotifySendChannel) formatTitle(r *report.Report) string {
	title := ch.MessageFormatter.FormatTitle(r)
	if title == "" {
		return "Critical temperature readings"
	}
	return title
}

func (ch *NotifySendChannel) formatBody(r *report.Report) string {
	return strings.Join(ch.MessageFormatter.FormatBodyRows(r), "\n")
}

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
	Config           config.Config
	MessageFormatter message.Formatter
}

func (ch *NotifySendChannel) IsEnable() bool {
	return ch.Config.Channels.NotifySend.Enable
}

func (ch *NotifySendChannel) Send(r report.Report) error {
	args := []string{
		"-u",
		"critical",
		"-t",
		strconv.Itoa(ch.Config.Channels.NotifySend.ExpireTime),
		"-a",
		"sensors-informer",
	}

	if ch.Config.Channels.NotifySend.Hint != "" {
		args = append(args, "-h", ch.Config.Channels.NotifySend.Hint)
	}

	args = append(args, ch.MessageFormatter.FormatTitle(&r), ch.formatBody(&r))

	os.Setenv("DISPLAY", ":0.0")
	cmd := exec.Command(ch.Config.Channels.NotifySend.Command, args...)
	_, err := cmd.Output()

	return err
}

func (ch *NotifySendChannel) formatBody(r *report.Report) string {
	return strings.Join(ch.MessageFormatter.FormatBodyRows(r), "\n")
}

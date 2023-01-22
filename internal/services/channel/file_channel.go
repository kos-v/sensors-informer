package channel

import (
	"fmt"
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/fs"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
	"os"
	"strings"
)

type FileChannelService struct {
	opts             config.FileChannelOpts
	messageFormatter message.Formatter
}

func (ch *FileChannelService) GetName() string {
	return "file"
}

func (ch *FileChannelService) IsEnable() bool {
	return ch.opts.Enable
}

func (ch *FileChannelService) Send(r report.Report) error {
	path := ch.opts.Path
	if path == "" {
		return fmt.Errorf("path is not specified in the config file")
	}

	err := fs.CreateNotExistsFilepath(path)
	if err != nil {
		return err
	}

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, info.Mode().Perm())
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(ch.format(r)); err != nil {
		return err
	}

	return nil
}

func (ch *FileChannelService) format(r report.Report) string {
	msg := ""
	if t := ch.messageFormatter.FormatTitle(&r); t != "" {
		msg += t + "\n"
	}

	msg += strings.Join(ch.messageFormatter.FormatBodyRows(&r), "\n")
	msg += "\n\n"

	return msg
}

func NewFileChannelService(
	opts config.FileChannelOpts,
	msgFormatter message.Formatter,
) *FileChannelService {
	return &FileChannelService{opts: opts, messageFormatter: msgFormatter}
}

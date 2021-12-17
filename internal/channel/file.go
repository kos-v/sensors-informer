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

type FileChannel struct {
	Opts             config.FileChannelOpts
	MessageFormatter message.Formatter
}

func (ch *FileChannel) IsEnable() bool {
	return ch.Opts.Enable
}

func (ch *FileChannel) Send(r report.Report) error {
	path := ch.Opts.Path
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

func (ch *FileChannel) format(r report.Report) string {
	msg := ""
	if t := ch.MessageFormatter.FormatTitle(&r); t != "" {
		msg += t + "\n"
	}

	msg += strings.Join(ch.MessageFormatter.FormatBodyRows(&r), "\n")
	msg += "\n\n"

	return msg
}

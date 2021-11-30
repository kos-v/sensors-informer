package channel

import (
	"fmt"
	"github.com/kos-v/sensors-informer/internal/config"
	"github.com/kos-v/sensors-informer/internal/fs"
	"github.com/kos-v/sensors-informer/internal/message"
	"github.com/kos-v/sensors-informer/internal/report"
	"os"
)

type FileChannel struct {
	Config           config.Config
	MessageFormatter message.Formatter
}

func (ch *FileChannel) IsEnable() bool {
	return ch.Config.Channels.File.Enable
}

func (ch *FileChannel) Send(r report.Report) error {
	path := ch.Config.Channels.File.Path
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
	if t := ch.MessageFormatter.FormatTitle(&r, true); t != "" {
		msg += t + ":\n"
	}

	for _, v := range ch.MessageFormatter.FormatBodyRows(&r) {
		msg += v + "\n"
	}

	msg += "\n"

	return msg
}

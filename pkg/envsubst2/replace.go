package envsubst2

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"regexp"
)

type Flag struct {
	Input       string `flag:"input" usage:"input file"`
	Output      string `flag:"output" usage:"output file"`
	ForceUpdate bool   `flag:"force-update" usage:"replace the placeholder, even if the environment value is empty"`
}

var patt = regexp.MustCompile(`\${([a-z0-9A-Z_]+)}`)

func Replace(ctx context.Context, flag *Flag) {
	// open input file
	in, err := os.Open(flag.Input)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	// open output file
	out, err := os.OpenFile(flag.Output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// readlines
	br := bufio.NewReader(in)
	for {
		line, _, err := br.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}

		matches := patt.FindAll(line, -1)
		for i := range matches {
			key := string(matches[i]) // key=${PORT}
			val := value(key)

			// force update
			if flag.ForceUpdate {
				line = bytes.ReplaceAll(line, []byte(key), []byte(val))
				continue
			}

			if val != "" {
				line = bytes.ReplaceAll(line, []byte(key), []byte(val))
			}
		}

		_, _ = out.Write(line)
		_, _ = out.Write([]byte("\n"))
	}
}

func value(key string) string {
	return os.ExpandEnv(key)
}

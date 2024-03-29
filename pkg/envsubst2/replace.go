package envsubst2

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"
)

type Flag struct {
	Input             string `flag:"input" usage:"input file"`
	Output            string `flag:"output" usage:"output file, os.Stdout if empty."`
	ForceReplace      bool   `flag:"force-replace" usage:"replace all the placeholders, even if their value is empty"`
	DefaultValuesFile string `flag:"default-values-file" usage:"specify the default values in YAML"`
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
	out := outWriter(flag.Output)
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
			val, exist := value(key)

			// force update
			if flag.ForceReplace {
				line = bytes.ReplaceAll(line, []byte(key), []byte(val))
				continue
			}

			// replace if the variable is set. regardless of the value is empty or not.
			if exist {
				line = bytes.ReplaceAll(line, []byte(key), []byte(val))
			}
		}

		_, _ = out.Write(line)
		_, _ = out.Write([]byte("\n"))
	}
}

func value(key string) (string, bool) {
	key = strings.Trim(key, "${}")

	val, exist := os.LookupEnv(key)

	if !exist {
		val, exist = getDefaultValue(key)
	}

	return val, exist
}

func outWriter(filename string) io.WriteCloser {
	if filename == "" {
		return os.Stdout
	}

	out, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return out
}

// Binary dltl fetches .tl schema from remote repo.
package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gotd/tl"
)

func main() {
	var (
		name   = flag.String("f", "api.tl", "file name to download; api.tl or mtproto.tl")
		base   = flag.String("base", "https://raw.githubusercontent.com/telegramdesktop/tdesktop", "base url")
		branch = flag.String("branch", "dev", "branch to use")
		dir    = flag.String("dir", "Telegram/Resources/tl", "directory of schemas")
		out    = flag.String("o", "", "output file name (blank to stdout)")
	)
	flag.Parse()

	u, err := url.Parse(*base)
	if err != nil {
		panic(err)
	}

	u.Path = path.Join(u.Path, *branch, *dir, *name)

	res, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer func() { _ = res.Body.Close() }()
	if res.StatusCode/100 != 2 {
		panic(fmt.Sprintf("status code %d", res.StatusCode))
	}

	// Parsing in-place.
	h := sha256.New()
	s, err := tl.Parse(io.TeeReader(res.Body, h))
	if err != nil {
		panic(err)
	}

	var outWriter io.Writer = os.Stdout
	if *out != "" {
		w, err := os.Create(*out)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := w.Close(); err != nil {
				panic(err)
			}
		}()
		outWriter = w
	}

	// Writing header to avoid manual edit.
	var b strings.Builder
	b.WriteString("// Code generated by ./cmd/dltl, DO NOT EDIT.\n")
	b.WriteString("//\n")
	b.WriteString(fmt.Sprintf("// Source: %s\n", u))
	if s.Layer > 0 {
		b.WriteString(fmt.Sprintf("// Layer:  %d\n", s.Layer))
	}
	b.WriteString(fmt.Sprintf("// SHA256: %x\n", h.Sum(nil)))
	b.WriteRune('\n')
	if _, err := io.WriteString(outWriter, b.String()); err != nil {
		panic(err)
	}

	if _, err := s.WriteTo(outWriter); err != nil {
		panic(err)
	}
}

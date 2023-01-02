package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PrismAIO/td/bin"
	"github.com/PrismAIO/td/internal/mt"
	"github.com/PrismAIO/td/internal/proto/codec"
	"github.com/PrismAIO/td/tg"
)

func Test_readAndPrint(t *testing.T) {
	c := codec.Intermediate{}

	input := &bytes.Buffer{}
	buf := &bin.Buffer{}

	objects := []bin.Object{
		&mt.RPCResult{},
		&mt.RPCError{},
		&tg.CodeSettings{},
	}
	for _, o := range objects {
		buf.Reset()
		require.NoError(t, o.Encode(buf))
		require.NoError(t, c.Write(input, buf))
	}

	output := &bytes.Buffer{}
	require.NoError(t, NewPrinter(input, formats("go"), c).Print(output))
	out := output.String()
	require.Contains(t, out, "RPCResult")
	require.Contains(t, out, "RPCError")
	require.Contains(t, out, "CodeSettings")
}

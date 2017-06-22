package cat

import (
	"context"
	"io"
	"os/exec"
)

var _ io.WriteCloser = new(Cat)

type Cat struct {
	in  io.WriteCloser
	cmd *exec.Cmd
}

func NewCat(ctx context.Context, w io.Writer) *Cat {
	c := new(Cat)
	/*#nosec*/ c.cmd = exec.CommandContext(ctx, "/bin/cat")
	c.cmd.Stdout = w
	var err error
	c.in, err = c.cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	err = c.cmd.Start()
	if err != nil {
		panic(err)
	}
	return c
}

func (c *Cat) Write(p []byte) (n int, err error) {
	return c.in.Write(p)
}

func (c *Cat) Close() error {
	errStdinClose := c.in.Close()
	errWait := c.cmd.Wait()
	if errStdinClose != nil {
		return errStdinClose
	}
	return errWait
}

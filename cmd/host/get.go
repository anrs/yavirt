package host

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/projecteru2/yavirt/cmd/run"
	"github.com/projecteru2/yavirt/internal/errors"
	"github.com/projecteru2/yavirt/internal/models"
)

func get(c *cli.Context, runtime run.Runtime) error {
	hn := c.Args().First()
	if len(hn) < 1 {
		return errors.New("host name is required")
	}

	h, err := model.LoadHost(hn)
	if err != nil {
		return errors.Trace(err)

	}

	fmt.Println(h)

	return nil
}

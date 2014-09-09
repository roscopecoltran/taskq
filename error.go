package taskq

import (
	"github.com/juju/errgo"
)

var (
	ErrNoTasks = errgo.New("No tasks")

	Mask = errgo.MaskFunc(IsErrNoTasks)
)

func IsErrNoTasks(err error) bool {
	return errgo.Cause(err) == ErrNoTasks
}

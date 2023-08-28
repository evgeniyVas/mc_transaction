package api

import (
	"context"
	"fmt"
	"github.com/mc_transaction/internal/models"
	"io"

	"github.com/pkg/errors"
)

var errAPIUnknown = errors.New("Unknown service error")
var errAPIInvalidPasswordFmt = errors.New("Invalid password format")

func writeErr(ctx context.Context, rw io.Writer, err error) {
	errM := &models.Error{Message: err.Error()}
	b, _ := errM.MarshalBinary()
	_, err = rw.Write(b)
	if err != nil {
		fmt.Errorf("write: err = %s", err)
		//logger.Errorf(ctx, "write: err = %s", err.Error())
	}
}

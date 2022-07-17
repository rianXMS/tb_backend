package features

import (
	"github.com/google/uuid"
	"tb_backend/util"
)

func write() {
	id, err := uuid.NewUUID()
	util.ErrorHandler(err)

}

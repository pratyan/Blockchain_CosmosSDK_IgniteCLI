package keeper

import (
	"github.com/pratyan/mars/x/mars/types"
)

var _ types.QueryServer = Keeper{}

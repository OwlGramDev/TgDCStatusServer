package antiflood

import "DcCheckerStatus/antiflood/types"

func Client() *Context {
	return &Context{
		ListFloodData: make(map[string]types.Data),
	}
}

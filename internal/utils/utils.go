package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type ParamsObject struct {
	Key    string
	Value  string
	Exists bool
}

// GetURLParams is used to extract params value from the URL
// takes the gin Comtext and a slice of keys for searching
// returns slice of PraamsObject
func GetURLParams(ctx *gin.Context, keys []string) (*[]ParamsObject, error) {
	if len(keys) == 0 {
		return nil, errors.New("keys array ie empty")
	}
	var params []ParamsObject
	for _, key := range keys {
		value, e := ctx.GetQuery(key)
		params = append(params, ParamsObject{Key: key, Value: value, Exists: e})
	}
	return &params, nil
}

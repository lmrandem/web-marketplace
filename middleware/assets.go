package middleware

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func Assets() gin.HandlerFunc {

	// Read files from dist folder
	dir, _ := ioutil.ReadDir("./public/dist")
	assetMap := make(map[string]string)
	for _, f := range dir {
		filename := f.Name()
		parts := strings.Split(filename, ".")
		name := strings.ReplaceAll(parts[0], "-", "_")
		ext := strings.ReplaceAll(parts[len(parts)-1], "-", "_")
		assetMap[fmt.Sprintf("%s_%s", name, ext)] = filename
	}

	return func(ctx *gin.Context) {
		// Add dist filenames to ctx
		values := ctx.GetStringMap("values")
		for k, v := range assetMap {
			values[k] = v
		}
		ctx.Set("values", values)
		ctx.Next()
	}
}

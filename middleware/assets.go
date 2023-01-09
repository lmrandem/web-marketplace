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
		assetMap[fmt.Sprintf("%s_%s", parts[0], parts[len(parts)-1])] = filename
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

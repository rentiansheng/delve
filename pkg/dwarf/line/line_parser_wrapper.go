package line

import (
	"os"
	"strings"
)

/***************************
    @author: tiansheng.ren
    @date: 2025/2/2
    @desc:

***************************/

var (
	goPath    = ""
	goPathTmp = ""
)

func init() {
	goPath = os.Getenv("go_path")
	goPathTmp = os.Getenv("go_path_tmp")
}

func wrapperBinaryCodeFile(f string) string {
	if len(goPathTmp) > 0 {
		f = strings.Replace(f, goPathTmp, goPath, 1)
	}

	return f
}

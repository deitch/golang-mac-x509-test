package init

import (
	"os"
)


var _ = setenv()

func setenv() bool {
	os.Setenv("GODEBUG", "x509roots=1")
	return true
}

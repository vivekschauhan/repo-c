package pkgc

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vivekschauhan/repo-b/pkg/pkgb"
)

func FuncC() {
	pkgb.FuncB()
	fmt.Printf("func c: %s\n", uuid.NewString())
}

// +build tools

// This file exists to manage the binaries relevant to this source tree via go modules.

package moderncrud

import (
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/99designs/gqlgen"
	_ "github.com/a8m/enter"
	_ "github.com/benbjohnson/litestream/cmd/litestream"
	_ "honnef.co/go/tools/cmd/staticcheck"
)

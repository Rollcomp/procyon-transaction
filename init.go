package tx

import core "github.com/procyon-projects/procyon-core"

func init() {
	core.Register(NewSimpleTransactionalContext)
}

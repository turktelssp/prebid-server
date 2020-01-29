package turktelekom

import (
	"text/template"

	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/usersync"
)

func NewTurktelekomSyncer(temp *template.Template) usersync.Usersyncer {
	return adapters.NewSyncer("turktelekom", 0, temp, adapters.SyncTypeRedirect)
}

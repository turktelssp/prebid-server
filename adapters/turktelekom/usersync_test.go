package turktelekom

import (
	"testing"
	"text/template"

	"github.com/prebid/prebid-server/privacy"
	"github.com/prebid/prebid-server/privacy/ccpa"
	"github.com/prebid/prebid-server/privacy/gdpr"
	"github.com/stretchr/testify/assert"
)

func TestTurktelekomSyncer(t *testing.T) {
	syncURL := "http://sp.programattik.com/s2s_sync?gdpr={{.GDPR}}&gdpr_consent={{.GDPRConsent}}&us_privacy={{.USPrivacy}}&redir=%2Fsetuid%3Fbidder%3Dturktelekom%26gdpr%3D{{.GDPR}}%26gdpr_consent%3D{{.GDPRConsent}}%26uid%3D%24%7BUUID%7D"
	syncURLTemplate := template.Must(
		template.New("sync-template").Parse(syncURL),
	)

	syncer := NewTurktelekomSyncer(syncURLTemplate)
	syncInfo, err := syncer.GetUsersyncInfo(privacy.Policies{
		GDPR: gdpr.Policy{
			Signal:  "A",
			Consent: "B",
		},
		CCPA: ccpa.Policy{
			Value: "C",
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, "http://sp.programattik.com/s2s_sync?gdpr=A&gdpr_consent=B&us_privacy=C&redir=%2Fsetuid%3Fbidder%3Dturktelekom%26gdpr%3DA%26gdpr_consent%3DB%26uid%3D%24%7BUUID%7D", syncInfo.URL)
	assert.Equal(t, "redirect", syncInfo.Type)
	assert.EqualValues(t, 0, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
}

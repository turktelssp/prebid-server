package turktelekom

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "turktelekomtest", NewTurktelekomBidder("http://localhost/prebid"))
}

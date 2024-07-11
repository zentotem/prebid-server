package zentotem

import (
	"github.com/prebid/prebid-server/v2/adapters/adapterstest"
	"github.com/prebid/prebid-server/v2/config"
	"github.com/prebid/prebid-server/v2/openrtb_ext"
	"testing"
)

// todo endpoint replace
func TestJsonSamples(t *testing.T) {
	bidder, buildErr := Builder(openrtb_ext.BidderZentotem, config.Adapter{
		Endpoint: "https://rtb.dev.zentotem.net/bid?sspuid=cq5r5n0dputspujutpmg"}, config.Server{ExternalUrl: "http://hosturl.com", GvlID: 1, DataCenter: "2"})

	if buildErr != nil {
		t.Fatalf("Builder returned unexpected error %v", buildErr)
	}

	adapterstest.RunJSONBidderTest(t, "zentotemtest", bidder)
}

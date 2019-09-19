package azure_resources

import (
	"flag"
	"testing"
)

var (
	clientID     string
	clientSecret string
	tenantID     string
)

func init() {
	testing.Init()
	flag.StringVar(&clientID, "client", "", "client to make the petition")
	flag.StringVar(&clientSecret, "secret", "", "secret to make the petition")
	flag.StringVar(&tenantID, "tenant", "", "tenant to make the petition")
	flag.Parse()
}

func TestGetAllByGroupName(t *testing.T) {
	const (
		subscriptionID = "72166d64-f454-46c1-a2f7-e84b57df44b8"
		groupName      = "terraform-resources"
	)

	if err := SetAuthorizer(tenantID, clientID, clientSecret); err != nil {
		t.Error(err)
	}

	resources, err := GetAllByGroupName(subscriptionID, groupName)
	if err != nil {
		t.Error(err)
	}

	for _, r := range resources {
		props, err := r.GetProperties()
		if err != nil {
			t.Error(err)
		}

		t.Logf("%s\n", props)
	}
}

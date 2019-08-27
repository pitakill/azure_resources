package azure_resources

import "testing"

func TestGetAllByGroupName(t *testing.T) {
	const (
		subscriptionID = "72166d64-f454-46c1-a2f7-e84b57df44b8"
		groupName      = "azure-test"
	)

	resources := GetAllByGroupName(subscriptionID, groupName)

	for _, r := range resources {
		props, err := r.GetProperties()
		if err != nil {
			t.Error(err)
		}

		t.Logf("%s\n", props)
	}
}

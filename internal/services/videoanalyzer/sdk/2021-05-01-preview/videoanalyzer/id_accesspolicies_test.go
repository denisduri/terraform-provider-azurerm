package videoanalyzer

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = AccessPoliciesId{}

func TestAccessPoliciesIDFormatter(t *testing.T) {
	actual := NewAccessPoliciesID("{subscriptionId}", "{resourceGroupName}", "{accountName}", "{accessPolicyName}").ID()
	expected := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies/{accessPolicyName}"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestParseAccessPoliciesID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AccessPoliciesId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/{subscriptionId}/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/{subscriptionId}/resourceGroups/",
			Error: true,
		},

		{
			// missing VideoAnalyzerName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/",
			Error: true,
		},

		{
			// missing value for VideoAnalyzerName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/",
			Error: true,
		},

		{
			// missing AccessPolicyName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/",
			Error: true,
		},

		{
			// missing value for AccessPolicyName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies/{accessPolicyName}",
			Expected: &AccessPoliciesId{
				SubscriptionId:    "{subscriptionId}",
				ResourceGroup:     "{resourceGroupName}",
				VideoAnalyzerName: "{accountName}",
				AccessPolicyName:  "{accessPolicyName}",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/{SUBSCRIPTIONID}/RESOURCEGROUPS/{RESOURCEGROUPNAME}/PROVIDERS/MICROSOFT.MEDIA/VIDEOANALYZERS/{ACCOUNTNAME}/ACCESSPOLICIES/{ACCESSPOLICYNAME}",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAccessPoliciesID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.VideoAnalyzerName != v.Expected.VideoAnalyzerName {
			t.Fatalf("Expected %q but got %q for VideoAnalyzerName", v.Expected.VideoAnalyzerName, actual.VideoAnalyzerName)
		}
		if actual.AccessPolicyName != v.Expected.AccessPolicyName {
			t.Fatalf("Expected %q but got %q for AccessPolicyName", v.Expected.AccessPolicyName, actual.AccessPolicyName)
		}
	}
}

func TestParseAccessPoliciesIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *AccessPoliciesId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/{subscriptionId}/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/{subscriptionId}/resourceGroups/",
			Error: true,
		},

		{
			// missing VideoAnalyzerName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/",
			Error: true,
		},

		{
			// missing value for VideoAnalyzerName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/",
			Error: true,
		},

		{
			// missing AccessPolicyName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/",
			Error: true,
		},

		{
			// missing value for AccessPolicyName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoAnalyzers/{accountName}/accessPolicies/{accessPolicyName}",
			Expected: &AccessPoliciesId{
				SubscriptionId:    "{subscriptionId}",
				ResourceGroup:     "{resourceGroupName}",
				VideoAnalyzerName: "{accountName}",
				AccessPolicyName:  "{accessPolicyName}",
			},
		},

		{
			// lower-cased segment names
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/videoanalyzers/{accountName}/accesspolicies/{accessPolicyName}",
			Expected: &AccessPoliciesId{
				SubscriptionId:    "{subscriptionId}",
				ResourceGroup:     "{resourceGroupName}",
				VideoAnalyzerName: "{accountName}",
				AccessPolicyName:  "{accessPolicyName}",
			},
		},

		{
			// upper-cased segment names
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/VIDEOANALYZERS/{accountName}/ACCESSPOLICIES/{accessPolicyName}",
			Expected: &AccessPoliciesId{
				SubscriptionId:    "{subscriptionId}",
				ResourceGroup:     "{resourceGroupName}",
				VideoAnalyzerName: "{accountName}",
				AccessPolicyName:  "{accessPolicyName}",
			},
		},

		{
			// mixed-cased segment names
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/ViDeOaNaLyZeRs/{accountName}/AcCeSsPoLiCiEs/{accessPolicyName}",
			Expected: &AccessPoliciesId{
				SubscriptionId:    "{subscriptionId}",
				ResourceGroup:     "{resourceGroupName}",
				VideoAnalyzerName: "{accountName}",
				AccessPolicyName:  "{accessPolicyName}",
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseAccessPoliciesIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.VideoAnalyzerName != v.Expected.VideoAnalyzerName {
			t.Fatalf("Expected %q but got %q for VideoAnalyzerName", v.Expected.VideoAnalyzerName, actual.VideoAnalyzerName)
		}
		if actual.AccessPolicyName != v.Expected.AccessPolicyName {
			t.Fatalf("Expected %q but got %q for AccessPolicyName", v.Expected.AccessPolicyName, actual.AccessPolicyName)
		}
	}
}

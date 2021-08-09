package naming_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/naming"
)

func TestParseCloudFormationTypeName(t *testing.T) {
	testCases := []struct {
		TestName             string
		Value                string
		ExpectError          bool
		ExpectedOrganization string
		ExpectedService      string
		ExpectedResource     string
	}{
		{
			TestName:    "empty string",
			Value:       "",
			ExpectError: true,
		},
		{
			TestName:    "whitespace string",
			Value:       "  ",
			ExpectError: true,
		},
		{
			TestName:    "incorrect number of segments",
			Value:       "AWS::EC2",
			ExpectError: true,
		},
		{
			TestName:    "invalid type name",
			Value:       "AWS::KMS::WayTooLongAResourceName000000000000000000000000000000000000000012",
			ExpectError: true,
		},
		{
			TestName:             "valid type name",
			Value:                "AWS::KMS::Key",
			ExpectedOrganization: "AWS",
			ExpectedService:      "KMS",
			ExpectedResource:     "Key",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			gotOrganization, gotService, gotResource, err := naming.ParseCloudFormationTypeName(testCase.Value)

			if err == nil && testCase.ExpectError {
				t.Fatalf("expected error, got no error")
			}

			if err != nil && !testCase.ExpectError {
				t.Fatalf("got unexpected error: %s", err)
			}

			if gotOrganization != testCase.ExpectedOrganization {
				t.Errorf("expected Organization: %s, got: %s", testCase.ExpectedOrganization, gotOrganization)
			}
			if gotService != testCase.ExpectedService {
				t.Errorf("expected Service: %s, got: %s", testCase.ExpectedService, gotService)
			}
			if gotResource != testCase.ExpectedResource {
				t.Errorf("expected Resource: %s, got: %s", testCase.ExpectedResource, gotResource)
			}
		})
	}
}

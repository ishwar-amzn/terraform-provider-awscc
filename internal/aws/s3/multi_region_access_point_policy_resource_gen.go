// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_s3_multi_region_access_point_policy", multiRegionAccessPointPolicyResource)
}

// multiRegionAccessPointPolicyResource returns the Terraform awscc_s3_multi_region_access_point_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::S3::MultiRegionAccessPointPolicy resource.
func multiRegionAccessPointPolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: MrapName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Multi Region Access Point to apply policy",
		//	  "maxLength": 50,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9][-a-z0-9]{1,48}[a-z0-9]$",
		//	  "type": "string"
		//	}
		"mrap_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Multi Region Access Point to apply policy",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 50),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z0-9][-a-z0-9]{1,48}[a-z0-9]$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Policy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Policy document to apply to a Multi Region Access Point",
		//	  "type": "object"
		//	}
		"policy": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Policy document to apply to a Multi Region Access Point",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Policy Status associated with this Multi Region Access Point",
		//	  "properties": {
		//	    "IsPublic": {
		//	      "description": "Specifies whether the policy is public or not.",
		//	      "enum": [
		//	        "true",
		//	        "false"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "IsPublic"
		//	  ],
		//	  "type": "object"
		//	}
		"policy_status": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IsPublic
				"is_public": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether the policy is public or not.",
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Policy Status associated with this Multi Region Access Point",
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The policy to be attached to a Multi Region Access Point",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::MultiRegionAccessPointPolicy").WithTerraformTypeName("awscc_s3_multi_region_access_point_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"is_public":     "IsPublic",
		"mrap_name":     "MrapName",
		"policy":        "Policy",
		"policy_status": "PolicyStatus",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

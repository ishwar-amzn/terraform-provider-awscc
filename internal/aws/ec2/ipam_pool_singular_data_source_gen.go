// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_ipam_pool", iPAMPoolDataSource)
}

// iPAMPoolDataSource returns the Terraform awscc_ec2_ipam_pool data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::IPAMPool resource.
func iPAMPoolDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AddressFamily
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The address family of the address space in this pool. Either IPv4 or IPv6.",
		//	  "type": "string"
		//	}
		"address_family": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The address family of the address space in this pool. Either IPv4 or IPv6.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AllocationDefaultNetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.",
		//	  "type": "integer"
		//	}
		"allocation_default_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AllocationMaxNetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum allowed netmask length for allocations made from this pool.",
		//	  "type": "integer"
		//	}
		"allocation_max_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum allowed netmask length for allocations made from this pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AllocationMinNetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The minimum allowed netmask length for allocations made from this pool.",
		//	  "type": "integer"
		//	}
		"allocation_min_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The minimum allowed netmask length for allocations made from this pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AllocationResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When specified, an allocation will not be allowed unless a resource has a matching set of tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"allocation_resource_tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "When specified, an allocation will not be allowed unless a resource has a matching set of tags.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IPAM Pool.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IPAM Pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AutoImport
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.",
		//	  "type": "boolean"
		//	}
		"auto_import": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AwsService
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Limits which service in Amazon Web Services that the pool can be used in.",
		//	  "enum": [
		//	    "ec2"
		//	  ],
		//	  "type": "string"
		//	}
		"aws_service": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Limits which service in Amazon Web Services that the pool can be used in.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IpamArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IPAM this pool is a part of.",
		//	  "type": "string"
		//	}
		"ipam_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IPAM this pool is a part of.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the IPAM Pool.",
		//	  "type": "string"
		//	}
		"ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the IPAM Pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpamScopeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the scope this pool is a part of.",
		//	  "type": "string"
		//	}
		"ipam_scope_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the scope this pool is a part of.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpamScopeId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of the scope this pool is a part of.",
		//	  "type": "string"
		//	}
		"ipam_scope_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of the scope this pool is a part of.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpamScopeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Determines whether this scope contains publicly routable space or space for a private network",
		//	  "enum": [
		//	    "public",
		//	    "private"
		//	  ],
		//	  "type": "string"
		//	}
		"ipam_scope_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Determines whether this scope contains publicly routable space or space for a private network",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Locale
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The region of this pool. If not set, this will default to \"None\" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.",
		//	  "type": "string"
		//	}
		"locale": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The region of this pool. If not set, this will default to \"None\" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PoolDepth
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The depth of this pool in the source pool hierarchy.",
		//	  "type": "integer"
		//	}
		"pool_depth": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The depth of this pool in the source pool hierarchy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProvisionedCidrs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of cidrs representing the address space available for allocation in this pool.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An address space to be inserted into this pool. All allocations must be made from this address space.",
		//	    "properties": {
		//	      "Cidr": {
		//	        "description": "Represents a single IPv4 or IPv6 CIDR",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Cidr"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"provisioned_cidrs": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Cidr
					"cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Represents a single IPv4 or IPv6 CIDR",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of cidrs representing the address space available for allocation in this pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublicIpSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is `byoip`.",
		//	  "enum": [
		//	    "byoip",
		//	    "amazon"
		//	  ],
		//	  "type": "string"
		//	}
		"public_ip_source": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is `byoip`.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PubliclyAdvertisable
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.",
		//	  "type": "boolean"
		//	}
		"publicly_advertisable": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceIpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.",
		//	  "type": "string"
		//	}
		"source_ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of this pool. This can be one of the following values: \"create-in-progress\", \"create-complete\", \"modify-in-progress\", \"modify-complete\", \"delete-in-progress\", or \"delete-complete\"",
		//	  "enum": [
		//	    "create-in-progress",
		//	    "create-complete",
		//	    "modify-in-progress",
		//	    "modify-complete",
		//	    "delete-in-progress",
		//	    "delete-complete"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of this pool. This can be one of the following values: \"create-in-progress\", \"create-complete\", \"modify-in-progress\", \"modify-complete\", \"delete-in-progress\", or \"delete-complete\"",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StateMessage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An explanation of how the pool arrived at it current state.",
		//	  "type": "string"
		//	}
		"state_message": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An explanation of how the pool arrived at it current state.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::IPAMPool",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAMPool").WithTerraformTypeName("awscc_ec2_ipam_pool")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address_family":                    "AddressFamily",
		"allocation_default_netmask_length": "AllocationDefaultNetmaskLength",
		"allocation_max_netmask_length":     "AllocationMaxNetmaskLength",
		"allocation_min_netmask_length":     "AllocationMinNetmaskLength",
		"allocation_resource_tags":          "AllocationResourceTags",
		"arn":                               "Arn",
		"auto_import":                       "AutoImport",
		"aws_service":                       "AwsService",
		"cidr":                              "Cidr",
		"description":                       "Description",
		"ipam_arn":                          "IpamArn",
		"ipam_pool_id":                      "IpamPoolId",
		"ipam_scope_arn":                    "IpamScopeArn",
		"ipam_scope_id":                     "IpamScopeId",
		"ipam_scope_type":                   "IpamScopeType",
		"key":                               "Key",
		"locale":                            "Locale",
		"pool_depth":                        "PoolDepth",
		"provisioned_cidrs":                 "ProvisionedCidrs",
		"public_ip_source":                  "PublicIpSource",
		"publicly_advertisable":             "PubliclyAdvertisable",
		"source_ipam_pool_id":               "SourceIpamPoolId",
		"state":                             "State",
		"state_message":                     "StateMessage",
		"tags":                              "Tags",
		"value":                             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

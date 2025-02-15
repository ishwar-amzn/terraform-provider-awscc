// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_transfer_certificate", certificateDataSource)
}

// certificateDataSource returns the Terraform awscc_transfer_certificate data source.
// This Terraform data source corresponds to the CloudFormation AWS::Transfer::Certificate resource.
func certificateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ActiveDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the active date for the certificate.",
		//	  "type": "string"
		//	}
		"active_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the active date for the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the unique Amazon Resource Name (ARN) for the agreement.",
		//	  "maxLength": 1600,
		//	  "minLength": 20,
		//	  "pattern": "arn:.*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the unique Amazon Resource Name (ARN) for the agreement.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Certificate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the certificate body to be imported.",
		//	  "maxLength": 16384,
		//	  "minLength": 1,
		//	  "pattern": "^[\t\n\r -ÿ]*",
		//	  "type": "string"
		//	}
		"certificate": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the certificate body to be imported.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateChain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the certificate chain to be imported.",
		//	  "maxLength": 2097152,
		//	  "minLength": 1,
		//	  "pattern": "^[\t\n\r -ÿ]*",
		//	  "type": "string"
		//	}
		"certificate_chain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the certificate chain to be imported.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the certificate.",
		//	  "maxLength": 22,
		//	  "minLength": 22,
		//	  "pattern": "^cert-([0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"certificate_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A textual description for the certificate.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^[\\w\\- ]*$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A textual description for the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InactiveDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the inactive date for the certificate.",
		//	  "type": "string"
		//	}
		"inactive_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the inactive date for the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NotAfterDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the not after date for the certificate.",
		//	  "type": "string"
		//	}
		"not_after_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the not after date for the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NotBeforeDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the not before date for the certificate.",
		//	  "type": "string"
		//	}
		"not_before_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the not before date for the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PrivateKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the private key for the certificate.",
		//	  "maxLength": 16384,
		//	  "minLength": 1,
		//	  "pattern": "^[\t\n\r -ÿ]*",
		//	  "type": "string"
		//	}
		"private_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the private key for the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Serial
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies Certificate's serial.",
		//	  "maxLength": 48,
		//	  "minLength": 0,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"serial": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies Certificate's serial.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A status description for the certificate.",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "PENDING",
		//	    "INACTIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A status description for the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose.",
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
		//	  "maxItems": 50,
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
			Description: "Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Describing the type of certificate. With or without a private key.",
		//	  "enum": [
		//	    "CERTIFICATE",
		//	    "CERTIFICATE_WITH_PRIVATE_KEY"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Describing the type of certificate. With or without a private key.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Usage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the usage type for the certificate.",
		//	  "enum": [
		//	    "SIGNING",
		//	    "ENCRYPTION"
		//	  ],
		//	  "type": "string"
		//	}
		"usage": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the usage type for the certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Transfer::Certificate",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Transfer::Certificate").WithTerraformTypeName("awscc_transfer_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"active_date":       "ActiveDate",
		"arn":               "Arn",
		"certificate":       "Certificate",
		"certificate_chain": "CertificateChain",
		"certificate_id":    "CertificateId",
		"description":       "Description",
		"inactive_date":     "InactiveDate",
		"key":               "Key",
		"not_after_date":    "NotAfterDate",
		"not_before_date":   "NotBeforeDate",
		"private_key":       "PrivateKey",
		"serial":            "Serial",
		"status":            "Status",
		"tags":              "Tags",
		"type":              "Type",
		"usage":             "Usage",
		"value":             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_rds_global_cluster", globalClusterResource)
}

// globalClusterResource returns the Terraform awscc_rds_global_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::RDS::GlobalCluster resource.
func globalClusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DeletionProtection
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled.",
		//	  "type": "boolean"
		//	}
		"deletion_protection": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Engine
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
		//	  "enum": [
		//	    "aurora",
		//	    "aurora-mysql",
		//	    "aurora-postgresql"
		//	  ],
		//	  "type": "string"
		//	}
		"engine": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"aurora",
					"aurora-mysql",
					"aurora-postgresql",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version number of the database engine to use. If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
		//	  "type": "string"
		//	}
		"engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version number of the database engine to use. If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: GlobalClusterIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.",
		//	  "pattern": "^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$",
		//	  "type": "string"
		//	}
		"global_cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceDBClusterIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) to use as the primary cluster of the global database. This parameter is optional. This parameter is stored as a lowercase string.",
		//	  "oneOf": [
		//	    {},
		//	    {}
		//	  ],
		//	  "type": "string"
		//	}
		"source_db_cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) to use as the primary cluster of the global database. This parameter is optional. This parameter is stored as a lowercase string.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StorageEncrypted
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": " The storage encryption setting for the new global database cluster.\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
		//	  "type": "boolean"
		//	}
		"storage_encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: " The storage encryption setting for the new global database cluster.\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::RDS::GlobalCluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::GlobalCluster").WithTerraformTypeName("awscc_rds_global_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"deletion_protection":          "DeletionProtection",
		"engine":                       "Engine",
		"engine_version":               "EngineVersion",
		"global_cluster_identifier":    "GlobalClusterIdentifier",
		"source_db_cluster_identifier": "SourceDBClusterIdentifier",
		"storage_encrypted":            "StorageEncrypted",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}

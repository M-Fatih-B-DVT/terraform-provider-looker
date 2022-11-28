package provider

import (
	"context"
	"regexp"

	"github.com/devoteamgcloud/terraform-provider-looker/pkg/lookergo"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resouceAlert() *schema.Resource {
	return &schema.Resource{
		Description: "", 
		CreateContext: 	resourceAlertCreate,
		ReadContext:   	resourceAlertRead,
		UpdateContext: 	resourceAlertUpdate,
		DeleteContext: 	resourceAlertDelete,
		Importer: 		&schema.ResourceImporter{
								StateContext: schema.ImportStatePassthroughContext,
							},
		Schema: map[string]*schema.Schema{
			"applied_dashboard_filters": {
				Description: "Filters coming from the dashboard that are applied.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_title":	{
							Description: "Field Title. Refer to `DashboardFilter.title` in [DashboardFilter](#!/types/DashboardFilter).",
							Type: schema.TypeString,
							Required: true,
						},
						"filter_name":	{
							Description: "Field Name. Refer to `DashboardFilter.dimension` in [DashboardFilter](#!/types/DashboardFilter).",
							Type: schema.TypeString,
							Required: true,
						},
						"filter_value":	{
							Description: "Field Value. [Filter Expressions](https://docs.looker.com/reference/filter-expressions).",
							Type: schema.TypeString,
							Required: true,
						},
						"filter_description":	{
							Description: "Human Readable Filter Description. This may be null or auto-generated.",
							Type: schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"comparison_type": {
				Description: " This property informs the check what kind of comparison we are performing. Only certain condition types are valid for time series alerts. For details, refer to [Setting Alert Conditions](https://docs.looker.com/sharing-and-publishing/creating-alerts#setting_alert_conditions) ",
				Type:	schema.TypeString,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringMatch(
						func() *regexp.Regexp {
							ret, _ := regexp.Compile("[EQUAL_TO|GREATER_THAN|GREATER_THAN_OR_EQUAL_TO|LESS_THAN|LESS_THAN_OR_EQUAL_TO|INCREASES_BY|DECREASES_BY|CHANGES_BY]")
							return ret
						}(),
						`Valid values are: "EQUAL_TO", "GREATER_THAN", "GREATER_THAN_OR_EQUAL_TO", "LESS_THAN", "LESS_THAN_OR_EQUAL_TO", "INCREASES_BY", "DECREASES_BY", "CHANGES_BY".`,
					),
				),
			},  
		},
	}
}
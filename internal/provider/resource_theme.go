package provider

import (
	"context"
	"fmt"

	"github.com/devoteamgcloud/terraform-provider-looker/pkg/lookergo"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTheme() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceThemeCreate,
		ReadContext:   resourceThemeRead,
		UpdateContext: resourceThemeUpdate,
		DeleteContext: resourceThemeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		// The resource schema is based on the request body for a POST call: 
		// https://developers.looker.com/api/explorer/4.0/methods/Theme/create_theme?sdk=go
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Description: "Name for Theme of LookML Models",
				Type:        schema.TypeString,
				Required:    true,
			},
			"begin_at": {
				Description: "",
				// use TypeString for time: 
				// https://developer.hashicorp.com/terraform/plugin/sdkv2/schemas/schema-types#date-time-data
				Type: schema.TypeString,
				ValidateFunc: validation.IsRFC3339Time, 
			},
		},
	},
}


func resourceThemeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	conn := m.(*Config).Api // .(*lookergo.Client)
}

func resourceThemeRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	
}

func resourceThemeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {

}

func resourceThemeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {

}

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				Description: "Timestamp for when this theme becomes active. Null=always",
				// use TypeString for time:
				// https://developer.hashicorp.com/terraform/plugin/sdkv2/schemas/schema-types#date-time-data
				Type:         schema.TypeString,
				ValidateFunc: validation.IsRFC3339Time,
			},
			"end_at": {
				Description: "Timestamp for when this theme expires. Null=never",
				// use TypeString for time:
				// https://developer.hashicorp.com/terraform/plugin/sdkv2/schemas/schema-types#date-time-data
				Type:         schema.TypeString,
				ValidateFunc: validation.IsRFC3339Time,
			},
			"settings": {
				// Inspired by the trigger-template block from the cloud build trigger resource from the official google provider
				// https://github.com/hashicorp/terraform-provider-google/blob/main/google/resource_cloudbuild_trigger.go
				// https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloudbuild_trigger#trigger_template
				Description: "settings",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"background_color": {
							Type:        schema.TypeString,
							Description: "Default background color",
						},
						"base_font_size": {
							Type:        schema.TypeString,
							Description: "Base font size for scaling fonts (only supported by legacy dashboards)",
						},
						"color_collection_id": {
							Type:        schema.TypeString,
							Description: "Optional. ID of color collection to use with the thme. Use an empty String for none",
							Optional:    true,
						},
						"font_color": {
							Description: "Default font color",
							Type:        schema.TypeString,
						},
						"font_family": {
							Description: "Source specification for font",
							Type:        schema.TypeString,
						},
						"font_source": {
							Description: "Source specification for font",
							Type:        schema.TypeString,
						},
						"info_button_color": {
							Description: "info button color",
							Type:        schema.TypeString,
						},
						"primary_button_color": {
							Type:        schema.TypeString,
							Description: "Primary button color",
						},
						"show_filters_bar": {
							Description: "Toggle to show filters. Defaults to true",
							Type:        schema.TypeBool,
						},
						"show_title": {
							Description: "Toggle to show title. Defaults to true",
							Type:        schema.TypeBool,
						},
						"text_tile_text_color": {
							Description: "Text color for text tiles",
							Type:        schema.TypeString,
						},
						"tile_background_color": {
							Description: "Background color for tiles",
							Type:        schema.TypeString,
						},
						"tile_text_color": {
							Description: "Text folor for tiles",
							Type:        schema.TypeString,
						},
						"title_color": {
							Description: "Color for titles",
							Type:        schema.TypeString,
						},
						"warn_button_color": {
							Description: "Warning button color",
							Type:        schema.TypeString,
						},
						"tile_title_alignment": {
							Description: "The text alignment of tile titles (New Dashboards)",
							Type:        schema.TypeString,
						},
						"tile_shadow": {
							Type:        schema.TypeBool,
							Description: "Toggles the tile shadow (not supported)",
						},
					},
				},
			},
		},
	}
}

func resourceThemeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	c := m.(*Config).Api // .(*lookergo.Client)
	return resourceModelSetRead(ctx, d, m)
}

func resourceThemeRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return resourceModelSetRead(ctx, d, m)
}

func resourceThemeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return resourceModelSetRead(ctx, d, m)
}

func resourceThemeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return resourceModelSetRead(ctx, d, m)
}

package provider

import (
	"context"
	"regexp"
	"time"

	lookergo "github.com/devoteamgcloud/terraform-provider-looker/pkg/lookergo"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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
			"Can": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
				Computed: true,
			},
			"Id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"Name": {
				Description: "Name for Theme of LookML Models",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringMatch(
					func() *regexp.Regexp {
						ret, _ := regexp.Compile("^[a-zA-Z0-9_]+$")
						return ret
					}(),
					"Name can only contain alphanumeric characters or underscores",
				),
			},
			"Begin_at": {
				Description: "Timestamp for when this theme becomes active. Null=always",
				// use TypeString for time:
				// https://developer.hashicorp.com/terraform/plugin/sdkv2/schemas/schema-types#date-time-data
				Type:         schema.TypeString,
				ValidateFunc: validation.IsRFC3339Time,
			},
			"End_at": {
				Description: "Timestamp for when this theme expires. Null=never",
				// use TypeString for time:
				// https://developer.hashicorp.com/terraform/plugin/sdkv2/schemas/schema-types#date-time-data
				Type:         schema.TypeString,
				ValidateFunc: validation.IsRFC3339Time,
			},
			"Settings": {
				// Inspired by the trigger-template block from the cloud build trigger resource from the official google provider
				// https://github.com/hashicorp/terraform-provider-google/blob/main/google/resource_cloudbuild_trigger.go
				// https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloudbuild_trigger#trigger_template
				Description: "settings",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"Background_color": {
							Type:        schema.TypeString,
							Description: "Default background color",
						},
						"Base_font_size": {
							Type:        schema.TypeString,
							Description: "Base font size for scaling fonts (only supported by legacy dashboards)",
						},
						"Color_collection_id": {
							Type:        schema.TypeString,
							Description: "Optional. ID of color collection to use with the thme. Use an empty String for none",
							Optional:    true,
						},
						"Font_color": {
							Description: "Default font color",
							Type:        schema.TypeString,
						},
						"Font_family": {
							Description: "Source specification for font",
							Type:        schema.TypeString,
						},
						"Font_source": {
							Description: "Source specification for font",
							Type:        schema.TypeString,
						},
						"Info_button_color": {
							Description: "info button color",
							Type:        schema.TypeString,
						},
						"Primary_button_color": {
							Type:        schema.TypeString,
							Description: "Primary button color",
						},
						"Show_filters_bar": {
							Description: "Toggle to show filters. Defaults to true",
							Type:        schema.TypeBool,
						},
						"Show_title": {
							Description: "Toggle to show title. Defaults to true",
							Type:        schema.TypeBool,
						},
						"Text_tile_text_color": {
							Description: "Text color for text tiles",
							Type:        schema.TypeString,
						},
						"Tile_background_color": {
							Description: "Background color for tiles",
							Type:        schema.TypeString,
						},
						"Tile_text_color": {
							Description: "Text folor for tiles",
							Type:        schema.TypeString,
						},
						"Title_color": {
							Description: "Color for titles",
							Type:        schema.TypeString,
						},
						"Warn_button_color": {
							Description: "Warning button color",
							Type:        schema.TypeString,
						},
						"Tile_title_alignment": {
							Description: "The text alignment of tile titles (New Dashboards)",
							Type:        schema.TypeString,
						},
						"Tile_shadow": {
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
	// prepare connection to API
	c := m.(*Config).Api // .(*lookergo.Client)
	tflog.Info(ctx, "Creating Looker theme")

	// prepare request body
	theme := lookergo.WriteTheme{
		Name: castToPtr(d.Get("Name").(string)),
	}

	if t, isNotNil := d.GetOk("Begin_at"); isNotNil {
		tmp, _ := time.Parse(time.RFC3339, t.(string))
		theme.BeginAt = castToPtr(tmp)
	}

	if t, isNotNil := d.GetOk("End_at"); isNotNil {
		tmp, _ := time.Parse(time.RFC3339, t.(string))
		theme.EndAt = castToPtr(tmp)
	}

	if settings, isNotNil := d.GetOk("Settings"); isNotNil {
		l := settings.([]interface{})                // Settings is a list of a single map[string]interface{}
		settingsMap := l[0].(map[string]interface{}) //

		themeSettings := lookergo.ThemeSettings{}

		for key, elem := range settingsMap {
			// SetField uses reflect, which is relatively inefficient.
			// key is the name of the json field.
			SetField(&themeSettings, key, castToPtr(elem))
		}

		theme.Settings = &themeSettings
	}

	// Makes a POST request to the API using the looker sdk
	newTheme, _, err := c.Theme.Create(ctx, &theme)

	if err != nil {
		return diag.FromErr(err)
	}

	// Id and Can are Computed fields
	d.SetId(*newTheme.Id)
	d.Set("Can", *newTheme.Can)

	tflog.Info(ctx, "Created Looker Theme", map[string]interface{}{"id": *newTheme.Id, "name": *newTheme.Name})
	return resourceThemeRead(ctx, d, m)
}

func resourceThemeRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	// prepare connection to API
	c := m.(*Config).Api // .(*lookergo.Client)

	theme, _, err := c.Theme.Get(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Transfer remote changes to terraform state
	d.Set("Id", *theme.Id)
	d.Set("Can", *theme.Can)
	d.Set("Name", *theme.Name)
	d.Set("Begin_at", *theme.BeginAt)
	d.Set("End_at", *theme.EndAt)
	//might need flattening
	d.Set("Settings", *theme.Settings)

	return diags
}

func resourceThemeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return resourceThemeRead(ctx, d, m)
}

func resourceThemeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return resourceThemeRead(ctx, d, m)
}

package lookergo

import (
	"context"
	"time"
)

const themeBasePath = "4.0/theme"

type Theme struct {
	Can      *map[string]bool `json:"can,omitempty"`      // Operations the current user is able to perform on this object
	BeginAt  *time.Time       `json:"begin_at,omitempty"` // Timestamp for when this theme becomes active. Null=always
	EndAt    *time.Time       `json:"end_at,omitempty"`   // Timestamp for when this theme expires. Null=never
	Id       *string          `json:"id,omitempty"`       // Unique Id
	Name     *string          `json:"name,omitempty"`     // Name of theme. Can only be alphanumeric and underscores.
	Settings *ThemeSettings   `json:"settings,omitempty"`
}

type ThemeSettings struct {
	BackgroundColor     *string `json:"background_color,omitempty"`      // Default background color
	BaseFontSize        *string `json:"base_font_size,omitempty"`        // Base font size for scaling fonts (only supported by legacy dashboards)
	ColorCollectionId   *string `json:"color_collection_id,omitempty"`   // Optional. ID of color collection to use with the theme. Use an empty string for none.
	FontColor           *string `json:"font_color,omitempty"`            // Default font color
	FontFamily          *string `json:"font_family,omitempty"`           // Primary font family
	FontSource          *string `json:"font_source,omitempty"`           // Source specification for font
	InfoButtonColor     *string `json:"info_button_color,omitempty"`     // Info button color
	PrimaryButtonColor  *string `json:"primary_button_color,omitempty"`  // Primary button color
	ShowFiltersBar      *bool   `json:"show_filters_bar,omitempty"`      // Toggle to show filters. Defaults to true.
	ShowTitle           *bool   `json:"show_title,omitempty"`            // Toggle to show the title. Defaults to true.
	TextTileTextColor   *string `json:"text_tile_text_color,omitempty"`  // Text color for text tiles
	TileBackgroundColor *string `json:"tile_background_color,omitempty"` // Background color for tiles
	TileTextColor       *string `json:"tile_text_color,omitempty"`       // Text color for tiles
	TitleColor          *string `json:"title_color,omitempty"`           // Color for titles
	WarnButtonColor     *string `json:"warn_button_color,omitempty"`     // Warning button color
	TileTitleAlignment  *string `json:"tile_title_alignment,omitempty"`  // The text alignment of tile titles (New Dashboards)
	TileShadow          *bool   `json:"tile_shadow,omitempty"`           // Toggles the tile shadow (not supported)
}

type WriteTheme struct {
	// Can and id will be set by looker, hence they are omitted
	BeginAt  *time.Time     `json:"begin_at,omitempty"` // Timestamp for when this theme becomes active. Null=always
	EndAt    *time.Time     `json:"end_at,omitempty"`   // Timestamp for when this theme expires. Null=never
	Name     *string        `json:"name,omitempty"`     // Name of theme. Can only be alphanumeric and underscores.
	Settings *ThemeSettings `json:"settings,omitempty"`
}

// https://developers.looker.com/api/explorer/4.0/methods/Theme?sdk=go
type ThemeResource interface {
	List(ctx context.Context) ([]Theme, *Response, error)
	Get(ctx context.Context, themeId string) (*Theme, *Response, error)
	Create(ctx context.Context, theme *WriteTheme) (*Theme, *Response, error)
	Update(ctx context.Context, themeId string, theme *Theme) (*Theme, *Response, error)
	Delete(ctx context.Context, themeId string) (*Response, error)
}

type ThemeResourceOp struct {
	client *Client
}

func (t ThemeResourceOp) List(ctx context.Context) ([]Theme, *Response, error) {
	return doList(ctx, t.client, themeBasePath, nil, new([]Theme))
}

func (t ThemeResourceOp) Get(ctx context.Context, themeId string) (*Theme, *Response, error) {
	return doGetById(ctx, t.client, themeBasePath, themeId, new(Theme))
}

func (t ThemeResourceOp) Create(ctx context.Context, theme *WriteTheme) (*Theme, *Response, error) {
	return doCreate(ctx, t.client, themeBasePath, theme, new(Theme))
}

func (t ThemeResourceOp) Update(ctx context.Context, themeId string, theme *Theme) (*Theme, *Response, error) {
	return doUpdate(ctx, t.client, themeBasePath, themeId, theme, new(Theme))
}

func (t ThemeResourceOp) Delete(ctx context.Context, themeId string) (*Response, error) {
	return doDelete(ctx, t.client, themeBasePath, themeId)
}

package lookergo

import (
	"context"
	"time"
)

const themeBasePath = "4.0/theme"

// https://developers.looker.com/api/explorer/4.0/methods/Theme?sdk=go
type ThemeResource interface {
	List(ctx context.Context) ([]Theme, *Response, error)
	Get(ctx context.Context, themeId string) (*Theme, *Response, error)
	Create(ctx context.Context, theme *Theme) (*Theme, *Response, error)
	Update(ctx context.Context, themeId string, theme *Theme) (*Theme, *Response, error)
	Delete(ctx context.Context, themeId string) (*Response, error)
}

type ThemeResourceOp struct {
	client *Client
}

// https://developers.looker.com/api/explorer/4.0/types/Theme/Theme?sdk=go
type Theme struct {
	Can      map[string]bool `json:"can,omitempty"`      // Operations the current user is able to perform on this object
	BeginAt  *time.Time      `json:"begin_at,omitempty"` // Timestamp for when this theme becomes active. Null=always
	EndAt    *time.Time      `json:"end_at,omitempty"`   // Timestamp for when this theme expires. Null=never
	ID       string          `json:"id,omitempty"`       // Unique Id
	Name     string          `json:"name,omitempty"`     // Name of theme. Can only be alphanumeric and underscores.
	Settings ThemeSettings   `json:"settings,omitempty"` // ThemeSettings is defined below
}

type ThemeSettings struct {
	Background_color      string `json:"background_color"`
	Base_font_size        string `json:"base_font_size"`
	Color_collection_id   string `json:"color_collection_id"`
	Font_color            string `json:"font_color"`
	Font_family           string `json:"font_family"`
	Font_source           string `json:"font-source"`
	Info_button_color     string `json:"info_button_color"`
	Primary_button_color  string `json:"primary_button_color"`
	Show_filters_bar      bool   `json:"show_filters_bar"`
	Show_title            bool   `json:"show_title"`
	Text_tile_text_color  string `json:"text_tile_text_color"`
	Tile_background_color string `json:"tile_background_color"`
	Tile_text_color       string `json:"tile_text_color"`
	Tile_color            string `json:"tile_color"`
	Warn_button_color     string `json:"warn_button_color"`
	Tile_title_alignment  string `json:"tile_title_alignment"`
	Tile_shadow           bool   `json:"tile_shadow"`
}

func (t ThemeResourceOp) List(ctx context.Context) ([]Theme, *Response, error) {
	return doList(ctx, t.client, themeBasePath, nil, new([]Theme))
}

func (t ThemeResourceOp) Get(ctx context.Context, themeId string) (*Theme, *Response, error) {
	return doGetById(ctx, t.client, themeBasePath, themeId, new(Theme))
}

func (t ThemeResourceOp) Create(ctx context.Context, theme *Theme) (*Theme, *Response, error) {
	return doCreate(ctx, t.client, themeBasePath, theme, new(Theme))
}

func (t ThemeResourceOp) Update(ctx context.Context, themeId string, theme *Theme) (*Theme, *Response, error) {
	return doUpdate(ctx, t.client, themeBasePath, themeId, theme, new(Theme))
}

func (t ThemeResourceOp) Delete(ctx context.Context, themeId string) (*Response, error) {
	return doDelete(ctx, t.client, themeBasePath, themeId)
}

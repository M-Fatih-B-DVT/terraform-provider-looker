package lookergo

import (
	"context"

	sdk "github.com/looker-open-source/sdk-codegen/go/sdk/v4" // imports Types Theme, ThemeSettings and WriteTheme
)

const themeBasePath = "4.0/theme"

// https://developers.looker.com/api/explorer/4.0/methods/Theme?sdk=go
type ThemeResource interface {
	List(ctx context.Context) ([]sdk.Theme, *Response, error)
	Get(ctx context.Context, themeId string) (*sdk.Theme, *Response, error)
	Create(ctx context.Context, theme *sdk.Theme) (*sdk.Theme, *Response, error)
	Update(ctx context.Context, themeId string, theme *sdk.Theme) (*sdk.Theme, *Response, error)
	Delete(ctx context.Context, themeId string) (*Response, error)
}

type ThemeResourceOp struct {
	client *Client
}

func (t ThemeResourceOp) List(ctx context.Context) ([]sdk.Theme, *Response, error) {
	return doList(ctx, t.client, themeBasePath, nil, new([]sdk.Theme))
}

func (t ThemeResourceOp) Get(ctx context.Context, themeId string) (*sdk.Theme, *Response, error) {
	return doGetById(ctx, t.client, themeBasePath, themeId, new(sdk.Theme))
}

func (t ThemeResourceOp) Create(ctx context.Context, theme *sdk.Theme) (*sdk.Theme, *Response, error) {
	return doCreate(ctx, t.client, themeBasePath, theme, new(sdk.Theme))
}

func (t ThemeResourceOp) Update(ctx context.Context, themeId string, theme *sdk.Theme) (*sdk.Theme, *Response, error) {
	return doUpdate(ctx, t.client, themeBasePath, themeId, theme, new(sdk.Theme))
}

func (t ThemeResourceOp) Delete(ctx context.Context, themeId string) (*Response, error) {
	return doDelete(ctx, t.client, themeBasePath, themeId)
}

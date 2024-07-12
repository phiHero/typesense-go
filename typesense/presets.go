package typesense

import (
	"context"

	"github.com/typesense/typesense-go/typesense/api"
)

type PresetsInterface interface {
	Retrieve(ctx context.Context) (*api.PresetsRetrieveSchema, error)
	Delete(ctx context.Context, presetName string, presetValue *api.PresetUpsertSchema) (*api.PresetSchema, error)
}

type presets struct {
	apiClient APIClientInterface
}

func (a *presets) Retrieve(ctx context.Context) (*api.PresetsRetrieveSchema, error) {
	response, err := a.apiClient.RetrieveAllPresetsWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}
	return response.JSON200, nil
}

func (a *presets) Delete(ctx context.Context, presetName string, presetValue *api.PresetUpsertSchema) (*api.PresetSchema, error) {
	response, err := a.apiClient.UpsertPresetWithResponse(ctx, presetName, *presetValue)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &HTTPError{Status: response.StatusCode(), Body: response.Body}
	}
	return response.JSON200, nil
}

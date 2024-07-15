package typesense

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/typesense/typesense-go/typesense/api"
)

func TestAnalyticsRulesRetrieve(t *testing.T) {
	expectedData := []*api.AnalyticsRuleSchema{
		{
			Name: "test_name",
			Type: "test_type",
			Params: api.AnalyticsRuleParameters{
				Limit: 10,
			},
		},
	}

	server, client := newTestServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		validateRequestMetadata(t, r, "/analytics/rules", http.MethodGet)
		data := jsonEncode(t, api.AnalyticsRulesRetrieveSchema{
			Rules: &expectedData,
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
	defer server.Close()

	res, err := client.Analytics().Rules().Retrieve(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedData, res)
}

func TestAnalyticsRulesRetrieveOnHttpStatusErrorCodeReturnsError(t *testing.T) {
	server, client := newTestServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		validateRequestMetadata(t, r, "/analytics/rules", http.MethodGet)
		w.WriteHeader(http.StatusConflict)
	})
	defer server.Close()

	_, err := client.Analytics().Rules().Retrieve(context.Background())
	assert.NotNil(t, err)
}

func TestAnalyticsRulesUpsert(t *testing.T) {
	upsertData := &api.AnalyticsRuleUpsertSchema{
		Type: "type2",
		Params: api.AnalyticsRuleParameters{
			Limit: 100,
		},
	}
	expectedData := &api.AnalyticsRuleSchema{
		Name:   "test-rule",
		Type:   upsertData.Type,
		Params: upsertData.Params,
	}

	server, client := newTestServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		validateRequestMetadata(t, r, "/analytics/rules/test-rule", http.MethodPut)

		var reqBody api.AnalyticsRuleUpsertSchema
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		assert.NoError(t, err)
		assert.Equal(t, *upsertData, reqBody)

		data := jsonEncode(t, api.AnalyticsRuleSchema{
			Name:   expectedData.Name,
			Type:   upsertData.Type,
			Params: upsertData.Params,
		})

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	})
	defer server.Close()

	res, err := client.Analytics().Rules().Upsert(context.Background(), expectedData.Name, upsertData)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, res)
}

func TestAnalyticsRulesUpsertOnHttpStatusErrorCodeReturnsError(t *testing.T) {
	server, client := newTestServerAndClient(func(w http.ResponseWriter, r *http.Request) {
		validateRequestMetadata(t, r, "/analytics/rules/test-rule", http.MethodPut)
		w.WriteHeader(http.StatusConflict)
	})
	defer server.Close()

	_, err := client.Analytics().Rules().Upsert(context.Background(), "test-rule", &api.AnalyticsRuleUpsertSchema{})
	assert.NotNil(t, err)
}
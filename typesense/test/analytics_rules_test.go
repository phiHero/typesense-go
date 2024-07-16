//go:build integration
// +build integration

package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func analyticsRulesCleanUp() {
	result, _ := typesenseClient.Analytics().Rules().Retrieve(context.Background())
	for _, rule := range result {
		typesenseClient.Analytics().Rule(rule.Name).Delete(context.Background())
	}
}
func TestAnalyticsRulesUpsert(t *testing.T) {
	t.Cleanup(analyticsRulesCleanUp)

	collectionName := createNewCollection(t, "analytics-rules-collection")
	ruleSchema := newAnalyticsRuleUpsertSchema(collectionName)
	ruleName := newUUIDName("test-rule")
	expectedData := newAnalyticsRule(ruleName, collectionName)

	result, err := typesenseClient.Analytics().Rules().Upsert(context.Background(), ruleName, ruleSchema)

	require.NoError(t, err)
	require.Equal(t, expectedData, result)
}

func TestAnalyticsRulesRetrieve(t *testing.T) {
	t.Cleanup(analyticsRulesCleanUp)

	collectionName := createNewCollection(t, "analytics-rules-collection")
	expectedRule := createNewAnalyticsRule(t, collectionName)

	results, err := typesenseClient.Analytics().Rules().Retrieve(context.Background())

	require.NoError(t, err)
	require.True(t, len(results) == 1, "number of rules is invalid")

	require.Equal(t, expectedRule, results[0])
}

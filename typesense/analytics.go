package typesense

type AnalyticsInterface interface {
	Rules() AnalyticsRulesInterface
	Rule(ruleName string) AnalyticsRuleInterface
}

type analytics struct {
	apiClient APIClientInterface
}

func (a *analytics) Rules() AnalyticsRulesInterface {
	return &analyticsRules{apiClient: a.apiClient}
}

func (a *analytics) Rule(ruleName string) AnalyticsRuleInterface {
	return &analyticsRule{apiClient: a.apiClient, ruleName: ruleName}
}

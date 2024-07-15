package typesense

type AnalyticsInterface interface {
	Rules() AnalyticsRulesInterface
}

type analytics struct {
	apiClient APIClientInterface
}

func (a *analytics) Rules() AnalyticsRulesInterface {
	return &analyticsRules{apiClient: a.apiClient}
}

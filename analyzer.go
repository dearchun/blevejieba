package blevejieba

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/analysis/token/lowercase"
	"github.com/blevesearch/bleve/v2/registry"
)

func AnalyzerConstructor(config map[string]interface{}, cache *registry.Cache) (*analysis.Analyzer, error) {
	tokenizer, err := cache.TokenizerNamed(Name)
	if err != nil {
		return nil, err
	}
	toLowerFilter, err := cache.TokenFilterNamed(lowercase.Name)
	if err != nil {
		return nil, err
	}
	stopFilter, err := cache.TokenFilterNamed(Name)
	if err != nil {
		return nil, err
	}

	return &analysis.Analyzer{
		Tokenizer: tokenizer,
		TokenFilters: []analysis.TokenFilter{
			toLowerFilter,
			stopFilter,
		},
	}, nil
}

func init() {
	registry.RegisterAnalyzer(Name, AnalyzerConstructor)
}

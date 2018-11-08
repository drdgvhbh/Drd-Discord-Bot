package api

type SearchCharacterResponse struct {
	RequestHash        string       `json:"request_hash"`
	RequestCached      bool         `json:"request_cached"`
	RequestCacheExpiry int          `json:"request_cache_expiry"`
	Results            []*Character `json:"results"`
	LastPage           int          `json:"last_page"`
}

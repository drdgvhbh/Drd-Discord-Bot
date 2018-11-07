package api

type SearchAnimeResponse struct {
	RequestHash        string   `json:"request_hash"`
	RequestCached      bool     `json:"request_cached"`
	RequestCacheExpiry int      `json:"request_cache_expiry"`
	Results            []*Anime `json:"results"`
	LastPage           int      `json:"last_page"`
}

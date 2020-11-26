package config

const (
	// CacheOptionPrefix key prefix for option cache
	CacheOptionPrefix = "PUTI_OPTION_"

	// CachePostCounterKey key of post counter cache, store all post's count number as a map
	CachePostCounterKey = "PUTI_POST_VIEWS_CACHE"
	// CachePostCounterIPPoolKeyPrefix key prefix, IP Pool was be designed as a single key/value cache
	CachePostCounterIPPoolKeyPrefix = "PUTI_POST_VIEWS_CACHE_IP_"

	// CachePageDetailPrefix key prefix for page cache
	CachePageDetailPrefix = "PUTI_PAGE_"

	// CacheArticleDetailPrefix key prefix for article cache
	CacheArticleDetailPrefix = "PUTI_ARTICLE_"

	// CacheKnowledgeInfoPrefix key prefix for knowledge info
	CacheKnowledgeInfoPrefix = "PUTI_K_"
	// CacheKnowledgeItemListPrefix key prefix for knowledge item list
	CacheKnowledgeItemListPrefix = "PUTI_KITEM_"
	// CacheKnowledgeItemContentPrefix key prefix for knowledge item content
	CacheKnowledgeItemContentPrefix = "PUTI_KITEM_CONTENT"
)

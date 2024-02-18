package consts

const (
	API_USAGE_KEY = "api:%d:usage"

	USER_USAGE_COUNT_FIELD  = "user.usage_count"
	USER_USED_TOKENS_FIELD  = "user.used_tokens"
	USER_TOTAL_TOKENS_FIELD = "user.total_tokens"

	APP_USAGE_COUNT_FIELD    = "app.%d.usage_count"
	APP_USED_TOKENS_FIELD    = "app.%d.used_tokens"
	APP_TOTAL_TOKENS_FIELD   = "app.%d.total_tokens"
	APP_IS_LIMIT_QUOTA_FIELD = "app.%d.is_limit_quota"

	KEY_USAGE_COUNT_FIELD    = "key.%d.%s.usage_count"
	KEY_USED_TOKENS_FIELD    = "key.%d.%s.used_tokens"
	KEY_TOTAL_TOKENS_FIELD   = "key.%d.%s.total_tokens"
	KEY_IS_LIMIT_QUOTA_FIELD = "key.%d.%s.is_limit_quota"
)

const (
	ERROR_MODEL_KEY       = "api:error:model:key:%s"
	ERROR_MODEL_AGENT     = "api:error:model:agent:%s"
	ERROR_MODEL_AGENT_KEY = "api:error:model:agent:key:%s"
)
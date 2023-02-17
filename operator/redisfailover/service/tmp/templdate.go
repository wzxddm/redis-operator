package tmp

import (
	_ "embed"
)

//go:embed redis.conf
var redisConfigTemplate string

//go:embed sentinel.conf
var sentinelConfigFileContent string

func GetRedisConfigTemplate() string {
	return redisConfigTemplate
}

func GetSentinelConfigFileContent() string {
	return sentinelConfigFileContent
}

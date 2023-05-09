package constant

import "time"

const (
	CacheExpiresInForever     = 0
	CacheExpiresInOneMinutes  = time.Minute
	CacheExpiresInTenMinutes  = time.Minute * 10
	CacheExpiresInOneHour     = time.Hour
	CacheExpiresInThreeHour   = time.Hour * 3
	CacheExpiresInOneDay      = CacheExpiresInOneHour * 24
	CacheExpiresInThreeDay    = CacheExpiresInOneDay * 3
	CacheExpiresInOneWeek     = CacheExpiresInOneDay * 7
	CacheExpiresInOneMonth    = CacheExpiresInOneDay * 30
	CacheExpiresInThreeMonths = CacheExpiresInOneMonth * 3
	CacheExpiresInSixMonths   = CacheExpiresInOneMonth * 6
	CacheExpiresInOneYear     = CacheExpiresInOneMonth * 12
)

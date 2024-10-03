package driver

const (
	UADefalut    = "Mozilla/5.0 (NE2210; 14; zh;) 115disk/32.3.1"
	UA115Browser = "Mozilla/5.0 (NE2210; 14; zh;) 115disk/32.3.1"
	UA115Disk    = "Mozilla/5.0 (NE2210; 14; zh;) 115disk/32.3.1"
	UA115Desktop = "Mozilla/5.0 (NE2210; 14; zh;) 115disk/32.3.1"
	UAIosApp     = "Mozilla/5.0 (NE2210; 14; zh;) 115disk/32.3.1"
)

const (
	CookieDomain115 = ".115.com"

	CookieUrl = "https://115.com"

	CookieNameUid  = "UID"
	CookieNameCid  = "CID"
	CookieNameSeid = "SEID"
)

const (
	OSSRegionID = "oss-cn-shenzhen"
	OSSEndpoint = "cn-shenzhen.oss.aliyuncs.com" // 双栈域名

	OSSUserAgent               = "aliyun-sdk-android/2.9.1"
	OssSecurityTokenHeaderName = "X-OSS-Security-Token"
)

const (
	KB = 1 << (10 * (iota + 1))
	MB
	GB
)

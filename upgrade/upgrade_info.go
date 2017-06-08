package upgrade

func NewUpgradeInfo(tip, url, md5, version string) *UpgradeInfo {
	return &UpgradeInfo{
		Tip:     tip,
		URL:     url,
		MD5:     md5,
		Version: version,
	}
}

type UpgradeInfo struct {
	Tip     string `json:"tip"`
	URL     string `json:"url""`
	MD5     string `json:"md5"`
	Version string `json:"version"`
}

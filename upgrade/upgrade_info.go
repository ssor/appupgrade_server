package upgrade

const (
	fake_company = "fakecompany"
)

func NewUpgradeInfo(upgrade, tip, url, md5 string) UpgradeInfo {
	return UpgradeInfo{
		Tip:     tip,
		URL:     url,
		MD5:     md5,
		Upgrade: upgrade,
	}
}

type UpgradeInfo struct {
	Tip     string `json:"tip" bson:"tip"`
	URL     string `json:"url" bson:"url"`
	MD5     string `json:"md5" bson:"md5"`
	Upgrade string `json:"version" bson:"version"` // yes or no
}

package messageHelpers

import "github.com/mt1976/frantic-core/commonConfig"

var cfg *commonConfig.Settings

func init() {
	cfg = commonConfig.Get()
}

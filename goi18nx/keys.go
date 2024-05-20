package goi18nx

const I18nKey = "SpectatorNan/goi18nx"
const I18nCurrentLangKey = "SpectatorNan/goi18nx/currentLang"

var (
	defaultLangHeaderKey        = "Accept-Language"
	defaultErrCode       uint32 = 10001
)

// SetDefaultLangHeaderKey sets the default value of the lang header key.
// need to be set before use
func SetDefaultLangHeaderKey(key string) {
	defaultLangHeaderKey = key
}

// SetDefaultErrCode sets the default value of the err code.
// need to be set before use
func SetDefaultErrCode(code uint32) {
	defaultErrCode = code
}

package goi18nx

const I18nKey = "SpectatorNan/goi18nx"

var (
	defaultLangHeaderKey = "Accept-Language"
)

// SetDefaultLangHeaderKey sets the default value of the lang header key.
// need to be set before use
func SetDefaultLangHeaderKey(key string) {
	defaultLangHeaderKey = key
}

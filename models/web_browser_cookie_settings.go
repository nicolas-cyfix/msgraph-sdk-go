package models
import (
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type WebBrowserCookieSettings int

const (
    // Browser default value, no intent.
    BROWSERDEFAULT_WEBBROWSERCOOKIESETTINGS WebBrowserCookieSettings = iota
    // Always block cookies.
    BLOCKALWAYS_WEBBROWSERCOOKIESETTINGS
    // Allow cookies from current Web site.
    ALLOWCURRENTWEBSITE_WEBBROWSERCOOKIESETTINGS
    // Allow Cookies from websites visited.
    ALLOWFROMWEBSITESVISITED_WEBBROWSERCOOKIESETTINGS
    // Always allow cookies.
    ALLOWALWAYS_WEBBROWSERCOOKIESETTINGS
)

func (i WebBrowserCookieSettings) String() string {
    return []string{"browserDefault", "blockAlways", "allowCurrentWebSite", "allowFromWebsitesVisited", "allowAlways"}[i]
}
func ParseWebBrowserCookieSettings(v string) (interface{}, error) {
    result := BROWSERDEFAULT_WEBBROWSERCOOKIESETTINGS
    switch v {
        case "browserDefault":
            result = BROWSERDEFAULT_WEBBROWSERCOOKIESETTINGS
        case "blockAlways":
            result = BLOCKALWAYS_WEBBROWSERCOOKIESETTINGS
        case "allowCurrentWebSite":
            result = ALLOWCURRENTWEBSITE_WEBBROWSERCOOKIESETTINGS
        case "allowFromWebsitesVisited":
            result = ALLOWFROMWEBSITESVISITED_WEBBROWSERCOOKIESETTINGS
        case "allowAlways":
            result = ALLOWALWAYS_WEBBROWSERCOOKIESETTINGS
        default:
            return 0, errors.New("Unknown WebBrowserCookieSettings value: " + v)
    }
    return &result, nil
}
func SerializeWebBrowserCookieSettings(values []WebBrowserCookieSettings) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package util

import (
	"net/url"
	"strconv"
	"strings"
)

func HttpUrlToWsUrl(httpBaseUrl string, paths ...string) string {
	wsBaseUrl := strings.Replace(httpBaseUrl, "http://", "ws://", 1)
	wsBaseUrl = strings.Replace(wsBaseUrl, "https://", "wss://", 1)
	parsedURL, _ := url.Parse(wsBaseUrl)
	path := parsedURL.JoinPath(paths...)
	//parsedURL = parsedURL.ResolveReference(&url.URL{Path: path.Join(paths...)})
	return path.String()
}

func GetPlayerIdFromMessage(msg map[string]any) (uint, error) {
	var playerId uint
	f, ok := msg["playerId"].(float64)
	if ok {
		playerId = uint(f)
	} else {
		parseUint, err := strconv.ParseUint(msg["playerId"].(string), 10, 32)
		if err != nil {
			return 0, nil
		}
		playerId = uint(parseUint)
	}
	return playerId, nil
}

func SanitizePlayerId(playerId uint) string {
	return strconv.Itoa(int(playerId))
}

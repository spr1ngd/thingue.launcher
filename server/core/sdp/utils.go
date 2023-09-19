package sdp

import "strconv"

func getPlayerIdFromMessage(msg map[string]any) (uint, error) {
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

func sanitizePlayerId(playerId uint) string {
	return strconv.Itoa(int(playerId))
}

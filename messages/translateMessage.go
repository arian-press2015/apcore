package messages

func TranslateMessage(message string, locale string) string {
	var translatedMessage string
	var exists bool

	switch locale {
	case "en":
		translatedMessage, exists = EnglishMessages[message]
	case "fa":
		translatedMessage, exists = FarsiMessages[message]
	}

	if !exists {
		return message
	}

	return translatedMessage
}

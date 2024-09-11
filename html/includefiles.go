package html

var IncludeFiles = map[string][]string{
	"index": {
		"html/index.html",
		"html/training/settings.html"},
	"training": {
		"html/training/training.html",
		"html/hand/hand.html",
		"html/hand/card.html",
		"html/training/button.html"},
	"card": {"html/hand/card.html"},
	"training-message": {
		"html/training/message.html",
		"html/training/button.html"},
}
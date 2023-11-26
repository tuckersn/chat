package trigger

type WebhookEntry struct {
	Name     *string
	Url      string
	Headers  map[string]string
	Subjects []string
}

var registry = map[string]WebhookEntry{}

func RegisterWebhook(name string, entry WebhookEntry) {
	registry[name] = entry
}

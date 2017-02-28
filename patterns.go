package clockwork

const (
	matchErrorNumber  = "([0-9])\\w+"
	matchTo           = "(To: [0-9])\\w+"
	matchID           = "(ID: [A-Z0-9])\\w+"
	matchCurrency     = "[-+]?([0-9]*\\.[0-9]+|[0-9]+)"
	matchCurrencyCode = "([A-Z]{2})\\w+"
)

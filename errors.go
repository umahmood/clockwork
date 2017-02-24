package clockwork

import "errors"

// You might get one of the following errors back when making a call to the API.
// Check the following link for more information:
// https://www.clockworksms.com/doc/reference/faqs/api-error-codes/

// ErrInternal internal error something went wrong in our API
var ErrInternal = errors.New("clockwork: internal error")

// ErrInvalidUsernamePassword invalid username or password
var ErrInvalidUsernamePassword = errors.New("clockwork: invalid username or password")

// ErrInsufficientCredit insufficient credits available
var ErrInsufficientCredit = errors.New("clockwork: insufficient credits available")

// ErrAuthFail authentication failure
var ErrAuthFail = errors.New("clockwork: authentication failure")

// ErrInvalidMsgType invalid 'MsgType' parameter
var ErrInvalidMsgType = errors.New("clockwork: invalid 'MsgType'")

// ErrMissingTo 'To' parameter not specified
var ErrMissingTo = errors.New("clockwork: 'To' parameter not specified")

// ErrMissingContent 'Content' parameter not specified
var ErrMissingContent = errors.New("clockwork: 'Content' parameter not specified")

// ErrMissingMessageID 'MessageID' parameter not specified
var ErrMissingMessageID = errors.New("clockwork: 'MessageID' parameter not specified")

// ErrUnknownMessageID unknown 'MessageID'
var ErrUnknownMessageID = errors.New("clockwork: unknown 'MessageID'")

// ErrInvalidTo invalid 'To' parameter
var ErrInvalidTo = errors.New("clockwork: invalid 'To' parameter some or all numbers were not correct")

// ErrInvalidFrom invalid 'From' parameter
var ErrInvalidFrom = errors.New("clockwork: invalid 'From' parameter")

// ErrMessageTooLong message text is too long
var ErrMessageTooLong = errors.New("clockwork: message text is too long")

// ErrRoutingMessage cannot route message
var ErrRoutingMessage = errors.New("clockwork: cannot route message")

// ErrMessageExpired message expired
var ErrMessageExpired = errors.New("clockwork: message expired")

// ErrNoRoute no route defined for this number
var ErrNoRoute = errors.New("clockwork: no route defined for this number")

// ErrMissingURL 'URL' parameter not set
var ErrMissingURL = errors.New("clockwork: 'URL' parameter not set")

// ErrInvalidSourceIP invalid source IP
var ErrInvalidSourceIP = errors.New("clockwork: invalid source IP")

// ErrMissingUDH 'UDH' parameter not specified
var ErrMissingUDH = errors.New("clockwork: 'UDH' parameter not specified")

// ErrInvalidServType invalid 'ServType' parameter
var ErrInvalidServType = errors.New("clockwork: invalid 'ServType' parameter")

// ErrInvalidExpiryTime invalid 'ExpiryTime' parameter
var ErrInvalidExpiryTime = errors.New("clockwork: invalid 'ExpiryTime' parameter")

// ErrDuplicateClientID duplicate 'ClientID' received
var ErrDuplicateClientID = errors.New("clockwork: duplicate 'ClientID' received")

// ErrInvalidTimeStamp invalid ‘TimeStamp’ parameter
var ErrInvalidTimeStamp = errors.New("clockwork: invalid ‘TimeStamp’ parameter")

// ErrInvalidAbsExpiry invalid ‘AbsExpiry’ parameter
var ErrInvalidAbsExpiry = errors.New("clockwork: invalid ‘AbsExpiry’ parameter")

// ErrInvalidDlrType invalid 'DlrType' parameter
var ErrInvalidDlrType = errors.New("clockwork: invalid 'DlrType' parameter")

// ErrInvalidConcat invalid 'Concat' parameter
var ErrInvalidConcat = errors.New("clockwork: invalid 'Concat' parameter")

// ErrInvalidUniqueID invalid 'UniqueID' parameter
var ErrInvalidUniqueID = errors.New("clockwork: invalid 'UniqueId' parameter")

// ErrClientIDRequired client id required. your account is setup to check for a
// unique client id on every message, one wasn't supplied in this send.
var ErrClientIDRequired = errors.New("clockwork: 'ClientID' Required - your account is setup to check for a unique client ID on every message, one wasn't supplied in this send")

// ErrInvalidCharInContent invalid character in 'Content' parameter
var ErrInvalidCharInContent = errors.New("clockwork: invalid character in 'Content' parameter")

// ErrInvalidTextPayload invalid 'Text' Payload MMS text has an invalid character
var ErrInvalidTextPayload = errors.New("clockwork: invalid 'TextPayload' MMS text has an invalid character")

// ErrInvalidHexPayload invalid 'Hex' Payload MMS payload can't be decoded as hex
var ErrInvalidHexPayload = errors.New("clockwork: invalid 'HexPayload' MMS Payload cant be decoded as hex")

// ErrInvalidBase64Payload invalid 'Base64' Payload MMS payload can't be decoded
// as base64.
var ErrInvalidBase64Payload = errors.New("clockwork: invalid 'Base64Payload' MMS payload cant be decoded as base64")

// ErrMissingContentType missing content type. No content type provided on MMS
// payload.
var ErrMissingContentType = errors.New("clockwork: missing 'ContentType'")

// ErrMissingID missing 'ID' all MMS payload parts must have an id
var ErrMissingID = errors.New("clockwork: missing 'ID' all MMS Payload parts must have an ID")

// ErrMMSMessageTooLarge MMS message too large the combined parts are too large
// to send
var ErrMMSMessageTooLarge = errors.New("clockwork: The combined parts are too large to send")

// ErrInvalidPayloadID invalid 'Payload' id
var ErrInvalidPayloadID = errors.New("clockwork: invalid payload ID")

// ErrDuplicatePayloadID duplicate payload id
var ErrDuplicatePayloadID = errors.New("clockwork: duplicate payload ID")

// ErrNoPayloadOnMMS no payload on MMS
var ErrNoPayloadOnMMS = errors.New("clockwork: no payload on MMS")

// ErrDuplicateFileName duplicate 'filename' attribute on payload. All MMS parts
// must have unique filenames.
var ErrDuplicateFileName = errors.New("clockwork: duplicate 'filename' attribute on payload -all MMS parts must have unique filenames")

// ErrMissingItemID 'ItemId' parameter not specified
var ErrMissingItemID = errors.New("clockwork: 'ItemId' parameter not specified")

// ErrInvalidItemID invalid 'ItemId' parameter
var ErrInvalidItemID = errors.New("clockwork: invalid 'ItemId' parameter")

// ErrGenerateFileName unable to generate filename for content-type
var ErrGenerateFileName = errors.New("clockwork: unable to generate filename for 'Content-Type'")

// ErrInvalidCharAction invalid 'InvalidCharAction' parameter
var ErrInvalidCharAction = errors.New("clockwork: invalid 'InvalidCharAction' parameter")

// ErrInvalidDlrEnroute invalid 'DlrEnroute' parameter
var ErrInvalidDlrEnroute = errors.New("clockwork: invalid 'DlrEnroute' parameter")

// ErrInvalidTruncate invalid 'Truncate' parameter
var ErrInvalidTruncate = errors.New("clockwork: invalid 'Truncate' parameter")

// ErrInvalidLong invalid 'Long' parameter
var ErrInvalidLong = errors.New("clockwork: invalid 'Long' parameter")

// ErrNoAPIKey no API key provided. You need to provide an API key or a
// username and password when calling the API.
var ErrNoAPIKey = errors.New("clockwork: no API key provided you need to provide an API Key or a user name and password when calling the API")

// ErrInvalidAPIKey invalid API key. Log in to your API account to check the key
// or create a new one
var ErrInvalidAPIKey = errors.New("clockwork: invalid API key - log in to your API account to check the key or create a new one")

// ErrMustUseAPIKeys account must use API keys. This account isn't allowed to
// use a username and password, log in to your account to create a key.
var ErrMustUseAPIKeys = errors.New("clockwork: account must use API keys")

// ErrBlockedSpam blocked by Spam filter. Sometimes your message will be caught
// by our Spam filter. If you're having trouble because of this error - get in
// touch.
var ErrBlockedSpam = errors.New("clockwork: blocked by spam filter - sometimes messages will be caught by our spam filter")

// ErrInvalidXML invalid XML API post can't be parsed as XML.
var ErrInvalidXML = errors.New("clockwork: invalid XML API post can't be parsed as XML")

// ErrInvalidXMLDoc XML document does not validate
var ErrInvalidXMLDoc = errors.New("clockwork: XML document does not validate")

// ErrLongClientID client id too long
var ErrLongClientID = errors.New("clockwork: Client ID too long")

// ErrRateExceeded query throttling rate exceeded. You've sent too many status
// requests this hour
var ErrRateExceeded = errors.New("clockwork: query throttling rate exceeded - you've sent too many status requests this hour")

// ErrStatusCode Clockwork SMS API returned a non 200 HTTP status code
var ErrStatusCode = errors.New("clockwork: API request did not return HTTP 200 OK")

// ErrUnknown if error is not in errorMap then this error will be returned
var ErrUnknown = errors.New("clockwork: unknown API error code")

// errorMap maps Clockwork API error codes to error messages. The keys (numbers)
// are important, they match the API error codes documented here:
// https://www.clockworksms.com/doc/reference/faqs/api-error-codes/
var errorMap = map[int]error{
	1:  ErrInternal,
	2:  ErrInvalidUsernamePassword,
	3:  ErrInsufficientCredit,
	4:  ErrAuthFail,
	5:  ErrInvalidMsgType,
	6:  ErrMissingTo,
	7:  ErrMissingContent,
	8:  ErrMissingMessageID,
	9:  ErrUnknownMessageID,
	10: ErrInvalidTo,

	11: ErrInvalidFrom,
	12: ErrMessageTooLong,
	13: ErrRoutingMessage,
	14: ErrMessageExpired,
	15: ErrNoRoute,
	16: ErrMissingURL,
	17: ErrInvalidSourceIP,
	18: ErrMissingUDH,
	19: ErrInvalidServType,
	20: ErrInvalidExpiryTime,

	25: ErrDuplicateClientID,
	26: ErrInternal,
	27: ErrInvalidTimeStamp,
	28: ErrInvalidAbsExpiry,
	29: ErrInvalidDlrType,

	31: ErrInvalidConcat,
	32: ErrInvalidUniqueID,
	33: ErrClientIDRequired,

	39: ErrInvalidCharInContent,
	40: ErrInvalidTextPayload,
	41: ErrInvalidHexPayload,
	42: ErrInvalidBase64Payload,
	43: ErrMissingContentType,
	44: ErrMissingID,
	45: ErrMMSMessageTooLarge,
	46: ErrInvalidPayloadID,
	47: ErrDuplicatePayloadID,
	48: ErrNoPayloadOnMMS,
	49: ErrDuplicateFileName,
	50: ErrMissingItemID,
	51: ErrInvalidItemID,
	52: ErrGenerateFileName,
	53: ErrInvalidCharAction,
	54: ErrInvalidDlrEnroute,
	55: ErrInvalidTruncate,
	56: ErrInvalidLong,
	57: ErrNoAPIKey,
	58: ErrInvalidAPIKey,
	59: ErrMustUseAPIKeys,
	60: ErrBlockedSpam,

	100: ErrInternal,
	101: ErrInternal,
	102: ErrInvalidXML,
	103: ErrInvalidXMLDoc,

	300: ErrLongClientID,
	305: ErrRateExceeded,
}

// errorFromCode looks up and returns the correct Clockwork API error. If the
// error code is unknown, returns ErrUnknown.
func errorFromCode(c int) error {
	if err, ok := errorMap[c]; ok {
		if err != nil {
			return err
		}
	}
	return ErrUnknown
}

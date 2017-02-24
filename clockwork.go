package clockwork

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// SendURL Clockwork SMS send http end point
const SendURL = "https://api.clockworksms.com/http/send.aspx"

// CreditURL Clockwork SMS credit http end point
const CreditURL = "https://api.clockworksms.com/http/balance"

// MsgType options
const (
	// TEXT Standard SMS message. Any character from the GSM character set can
	// be used. A single SMS can contain 160 standard GSM characters.
	// https://www.clockworksms.com/doc/reference/faqs/gsm-character-set/
	TEXT = "TEXT"
	// UCS2 Unicode SMS Message. Any Unicode characters from the UCS-2 character
	// set can be sent.
	UCS2 = "UCS2"
)

// Concat options
const (
	// OnePart SMS can contain 160 standard GSM characters
	OnePart = iota + 1
	// TwoParts SMS can contain 306 standard GSM characters
	TwoParts
	// ThreeParts SMS can contain 459 standard GSM characters
	ThreeParts
)

// InvalidCharAction options
const (
	// ErrorOnInvalidChars if invalid characters in the message are encountered
	ErrorOnInvalidChars = iota + 1
	// RemoveInvalidChars invalid characters in the message if encountered
	RemoveInvalidChars
	// ReplaceInvalidChars invalid characters where possible, remove the rest
	ReplaceInvalidChars
)

// Truncate options
const (
	// ErrorIfContentTooLong return error if content is too long
	ErrorIfContentTooLong = iota + 1
	// ReplaceExtraText remove the extra text
	ReplaceExtraText
)

// SMS represents a single SMS message
type SMS struct {
	// To list of up to 50 numbers. Phone numbers must be in international number
	// format without a leading ‘+’ or international dialing prefix such as ‘00’,
	// e.g. 441234567890.
	To []string
	// Content the message you want to send. Mobile networks only support
	// characters listed in the GSM character set.
	Content string
	// The text or phone number displayed when a text message is received on a
	// phone. This can be either a 12 digit number or 11 characters long. You
	// can set a default by logging in to Clockwork.
	From string
	// MsgType message type the default is TEXT.
	MsgType string
	// Concat The maximum number of parts for concatenated messages. Defaults to
	// 1 part, maximum 3. This parameter only affects TEXT message types. Each
	// part is billed as an individual message. When set to 1 the platform will
	// only allow 1 SMS to be sent. When set to 2, it allows either one or two
	// SMS to be sent. When set to 3, the platform allows 1,2 or 3 SMS to be
	// sent depending on the size of the content. Possible values - OnePart,
	// TwoParts, ThreeParts.
	Concat int
	// ClientID a unique Message ID specified by the connecting application,
	// maximum length: 50 characters.
	ClientID string
	// Expiry the number of minutes before the message expires. The minimum is
	// 10 and the maximum is 2160 (36 hours). Expiry time may not be honored by
	// some mobile networks.
	Expiry time.Duration
	// AbsExpiry the Absolute Expiry time for the message. An absolute time
	// should be specified in UTC. Expiry time may not be honored by some mobile
	// networks.
	AbsExpiry time.Time
	// UniqueIDChecks enable unique ID checks. If enabled, the ClientID specified
	// by the connecting application must be unique within the last 12 hours.
	// This is to prevent the connecting application from falsely sending
	// duplicate messages to a phone.
	UniqueIDChecks bool
	// InvalidCharAction what to do with any invalid characters in the message
	// content. Possible values - Error, Remove, Replace.
	InvalidCharAction int
	// Truncate trims the message content to the maximum length if it’s too
	// long. Truncate only works with standard text messages (MsgType=TEXT).
	// Possible values - ErrorIfContentTooLong, ReplaceExtraText.
	Truncate int
}

// SMSResponse stores telephone numbers and there associated meta data
type SMSResponse map[string]map[string]string

// Doer implemented by any type which can do HTTP requests
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Numbers slice of telephone numbers
type Numbers []string

// Clockwork instance
type Clockwork struct {
	apiKey string
}

// New creates a new instance of Clockwork SMS
func New(apiKey string) *Clockwork {
	return &Clockwork{
		apiKey: apiKey,
	}
}

// Send sends a single SMS message. If a message contains a mixture of valid and
// invalid numbers, then the method will return ErrInvalidTo. The valid numbers
// will be in the 'SMSResponse' map. To determine which numbers are invalid,
// check for there existent in the 'SMSResponse' map. For example:
//
//      msg := clockwork.SMS{
//              To: clockwork.Numbers{  "13052645330",   // valid
//                                      "4412345678910", // valid
//                                      "123",           // invalid
//                                      "meh!",          // invalid :)
//                                  }
//              ...
//              }
//
//      resp, err := cw.Send(msg)
//      if err == ErrInvalidTo {
//          // all/some of the numbers are invalid
//			// ...
//      }
//
//      // 'resp' map holds only the valid numbers SMS messages were sent to
//      resp["13052645330"] == true
//      resp["meh!"] == false
//
func (c *Clockwork) Send(sms SMS) (SMSResponse, error) {
	return DoSendRequestHelper(c, c.apiKey, SendURL, sms)
}

// Credit check how much credit you have left on your account
func (c *Clockwork) Credit() (credit float64, code string, err error) {
	return DoCreditRequestHelper(c, c.apiKey, CreditURL)
}

// Do performs a HTTP request
func (c *Clockwork) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DoSendRequestHelper helper to make a HTTP Get request using 'sms' values
func DoSendRequestHelper(d Doer, key string, url string, sms SMS) (SMSResponse, error) {
	m := smsSetOptions(sms)
	m["Key"] = key
	q := urlEncode(m)

	req, err := http.NewRequest("GET", url+"?"+q, nil)
	if err != nil {
		return nil, err
	}

	// identify this client to the clockwork SMS API
	userAgent := "Clockwork Go wrapper/" + Version()
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")

	resp, err := d.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		return nil, ErrStatusCode
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseSendResponseBody(string(body))
}

// DoCreditRequestHelper helper to make a HTTP Get request
func DoCreditRequestHelper(d Doer, key string, url string) (credit float64, code string, err error) {

	req, err := http.NewRequest("GET", url+"?key="+key, nil)
	if err != nil {
		return 0, "", err
	}

	// identify this client to the clockwork SMS API
	userAgent := "Clockwork Go wrapper/" + Version()
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")

	resp, err := d.Do(req)
	if err != nil {
		return 0, "", err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		return 0, "", ErrStatusCode
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}

	return parseCreditResponseBody(string(body))
}

// parseSendResponseBody parse the plain text response body from a clockwork
// /send HTTP call.
func parseSendResponseBody(body string) (SMSResponse, error) {
	// if some of the numbers provided to SMS message contained bad numbers,
	// flag it. So we can return the correct error to the caller.
	var badNumbers bool
	nums := make(SMSResponse)
	scanner := bufio.NewScanner(strings.NewReader(body))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Error") {
			// error response
			if strings.Contains(line, "To:") {
				// special case - if user specified invalid 'To' number, error
				// response is formatted as:
				//
				//       To: 123 Error 10: Invalid 'To' Parameter
				//       To: 456 Error 10: Invalid 'To' Parameter
				//       To: 441234567890 ID: VE_439221450
				//
				// the first two lines are bad numbers the last line is a valid
				// number.
				badNumbers = true
				continue
			} else {
				// all other error responses are formated as:
				//
				//      Error <number>: <message>
				//
				r := regexp.MustCompile("([0-9])\\w+")
				i, err := strconv.Atoi(r.FindString(line))
				if err != nil {
					return nil, err
				}
				return nil, errorFromCode(i)
			}
		} else {
			// valid responses are formatted as:
			//
			//      To: <number> ID: <id>
			//
			extractTo := regexp.MustCompile("(To: [0-9])\\w+")
			extractID := regexp.MustCompile("(ID: [A-Z0-9])\\w+")

			to := strings.Split(extractTo.FindString(line), " ")[1]
			id := strings.Split(extractID.FindString(line), " ")[1]

			nums[to] = map[string]string{"ID": id}
		}
	}

	if badNumbers {
		if len(nums) > 0 {
			return nums, ErrInvalidTo
		}
		return nil, ErrInvalidTo
	}

	return nums, nil
}

// parseCreditResponseBody parse the plain text response body from a clockwork
// /credit HTTP call.
func parseCreditResponseBody(body string) (float64, string, error) {
	if strings.Contains(body, "Error") {
		// error response is in the following plain-text format:
		//
		//	Error 58: Invalid API Key
		//
		r := regexp.MustCompile("([0-9])\\w+")
		n := r.FindString(body)
		i, err := strconv.Atoi(n)
		if err != nil {
			return 0, "", err
		}
		return 0, "", errorFromCode(i)
	}
	// valid response body is in the following plain-text format:
	//
	//	Balance: 287.58 (GBP)
	//
	// extract amount 287.58 and currency code GBP
	extractAmount := regexp.MustCompile("[-+]?([0-9]*\\.[0-9]+|[0-9]+)")
	extractCurrencyCode := regexp.MustCompile("([A-Z]{2})\\w+")
	a := extractAmount.FindString(body)
	c := extractCurrencyCode.FindString(body)
	s, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, "", err
	}
	return s, c, nil
}

// urlEncode encodes key/value pairs in a map as a HTTP Get query string. e.g.
// {"Message": "Hello World"}  -> "Message=Hello+World"
func urlEncode(queryValues map[string]string) string {
	u := url.Values{}
	for k, v := range queryValues {
		u.Add(k, v)
	}
	return u.Encode()
}

// smsSetOptions returns all the user set fields on the SMS type.
func smsSetOptions(sms SMS) map[string]string {
	vals := make(map[string]string)
	if sms.To != nil {
		vals["To"] = strings.Join(sms.To, ",")
	}
	if sms.Content != "" {
		vals["Content"] = sms.Content
	}
	if sms.From != "" {
		vals["From"] = sms.From
	}
	if sms.MsgType != "" {
		vals["MsgType"] = sms.MsgType
	}
	if sms.Concat != 0 {
		vals["Concat"] = strconv.Itoa(sms.Concat)
	}
	if sms.ClientID != "" {
		vals["ClientID"] = sms.ClientID
	}
	if sms.Expiry.Minutes() > 10 { // 10 = minimum expiry time
		vals["ExpiryTime"] = strconv.Itoa(int(sms.Expiry.Minutes()))
	}
	if !sms.AbsExpiry.IsZero() {
		vals["AbsExpiry"] = formatTime(sms.AbsExpiry)
	}
	if sms.UniqueIDChecks {
		vals["UniqueId"] = "1"
	}
	if sms.InvalidCharAction != 0 {
		vals["InvalidCharAction"] = strconv.Itoa(sms.InvalidCharAction)
	}
	if sms.Truncate != 0 {
		vals["Truncate"] = strconv.Itoa(sms.Truncate - 1)
	}
	return vals
}

// formatTime formats a time instance in the form yyyyMMddHHmm e.g. 201110201530
func formatTime(t time.Time) string {
	return fmt.Sprintf("%04d%02d%02d%02d%02d", t.Year(), t.Month(),
		t.Day(), t.Hour(), t.Minute())
}

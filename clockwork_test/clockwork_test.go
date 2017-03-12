package clockwork_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/umahmood/clockwork"
)

const testAPIKey = "TEST-KEY"

// mockClockwork test instance
type mockClockwork struct {
	f      func(req *http.Request) (*http.Response, error)
	apiKey string
}

// Send mock SMS
func (m *mockClockwork) Send(sms clockwork.SMS) (clockwork.SMSResponse, error) {
	return clockwork.DoSendRequestHelper(m, m.apiKey, "https://test.com/send", sms)
}

// Credit mock credit
func (m *mockClockwork) Credit() (amount float64, code string, err error) {
	return clockwork.DoCreditRequestHelper(m, m.apiKey, "https://test.com/credit")
}

// Do mock HTTP request
func (m *mockClockwork) Do(req *http.Request) (*http.Response, error) {
	return m.f(req)
}

// TestSingleValidSMSNumber
func TestSingleValidSMSNumber(t *testing.T) {
	mock := &mockClockwork{
		apiKey: testAPIKey,
		f: func(req *http.Request) (*http.Response, error) {
			body := "To: 1234567890 ID: VE_439333520"
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(body)),
			}, nil
		},
	}

	msg := clockwork.SMS{
		To:      clockwork.Numbers{"1234567890"},
		From:    "Gopher",
		Content: "Gophers rule!",
	}

	wantNumber := "1234567890"
	wantID := "VE_439333520"

	resp, err := mock.Send(msg)
	if err != nil {
		t.Error(err)
	}

	for gotNumber, gotInfo := range resp {
		if gotNumber != wantNumber {
			t.Errorf("Fail: got %s want %s", gotNumber, wantNumber)
		}
		gotID := gotInfo["ID"]
		if gotID != wantID {
			t.Errorf("Fail: got %s want %s", gotID, wantID)
		}
	}
}

// TestSingleInvalidSMSNumber
func TestSingleInvalidSMSNumber(t *testing.T) {
	mock := &mockClockwork{
		apiKey: testAPIKey,
		f: func(req *http.Request) (*http.Response, error) {
			body := "To: 123 Error 10: Invalid 'To' Parameter"
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(body)),
			}, nil
		},
	}

	msg := clockwork.SMS{
		To:      clockwork.Numbers{"123"},
		From:    "Gopher",
		Content: "Gophers rule!",
	}

	resp, err := mock.Send(msg)
	if err != clockwork.ErrInvalidTo {
		t.Errorf("Fail: err - got %s want %s", err, clockwork.ErrInvalidTo)
	}

	if resp != nil {
		t.Errorf("Fail: resp - got %v want nil", resp)
	}
}

// TestBothValidAndInvalidNumbers
func TestBothValidAndInvalidNumbers(t *testing.T) {
	mock := &mockClockwork{
		apiKey: testAPIKey,
		f: func(req *http.Request) (*http.Response, error) {
			body := `To: 123 Error 10: Invalid 'To' Parameter
					 To: 456 Error 10: Invalid 'To' Parameter
					 To: 13053696625 ID: LA_360224205
					 To: 44123456789 ID: VE_360224242`
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(body)),
			}, nil
		},
	}

	msg := clockwork.SMS{
		To:      clockwork.Numbers{"123", "456", "13053696625", "44123456789"},
		From:    "Gopher",
		Content: "Gophers rule!",
	}

	resp, err := mock.Send(msg)
	if err != clockwork.ErrInvalidTo {
		t.Errorf("Fail: err - got %s want %s", err, clockwork.ErrInvalidTo)
	}

	if resp["123"] != nil {
		t.Errorf("Fail: invalid number 123 in map")
	} else if resp["456"] != nil {
		t.Errorf("Fail: invalid number 456 in map")
	}

	if resp["13053696625"] == nil {
		t.Errorf("Fail: valid number 13053696625 not in map")

	} else if resp["44123456789"] == nil {
		t.Errorf("Fail: valid number 44123456789 not in map")
	}

	if len(resp) != 2 {
		t.Errorf("Fail: len map got %d want 2", len(resp))
	}
}

// TestCredit
func TestCredit(t *testing.T) {
	testCases := []struct {
		body       string
		wantCode   string
		wantCredit float64
	}{
		{
			body:       "Balance: 2048.12 (GBP)",
			wantCode:   "GBP",
			wantCredit: 2048.12,
		},
		{
			body:       "Balance: 234.56 (GBP)",
			wantCode:   "GBP",
			wantCredit: 234.56,
		},
		{
			body:       "Balance: 42.58 (GBP)",
			wantCode:   "GBP",
			wantCredit: 42.58,
		},
		{
			body:       "Balance: 1.42 (GBP)",
			wantCode:   "GBP",
			wantCredit: 1.42,
		},
		{
			body:       "Balance: 0.45 (GBP)",
			wantCode:   "GBP",
			wantCredit: 0.45,
		},
		{
			body:       "Balance: 0.00 (GBP)",
			wantCode:   "GBP",
			wantCredit: 0.00,
		},
		{
			body:       "Balance: 0 (GBP)",
			wantCode:   "GBP",
			wantCredit: 0,
		},
		{
			body:       "Balance: -94.23 (GBP)",
			wantCode:   "GBP",
			wantCredit: -94.23,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input_%s", tc.body), func(t *testing.T) {
			mock := &mockClockwork{
				apiKey: testAPIKey,
				f: func(req *http.Request) (*http.Response, error) {
					body := tc.body
					return &http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader(body)),
					}, nil
				},
			}
			gotCredit, gotCode, err := mock.Credit()
			if err != nil {
				t.Errorf("Fail err - got %v want nil", err)
			}

			if gotCredit != tc.wantCredit {
				t.Errorf("Fail credit - got %v want %v", gotCredit, tc.wantCredit)
			}

			if gotCode != tc.wantCode {
				t.Errorf("Fail code - got %s want %s", gotCode, tc.wantCode)
			}

		})
	}
}

// TestCreditError
func TestCreditError(t *testing.T) {
	mock := &mockClockwork{
		apiKey: testAPIKey,
		f: func(req *http.Request) (*http.Response, error) {
			body := "Error 58: Invalid API Key"
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(body)),
			}, nil
		},
	}

	wantError := clockwork.ErrInvalidAPIKey

	gotCredit, gotCode, gotError := mock.Credit()

	if gotCredit != 0 {
		t.Errorf("Fail: credit got %v want %v", gotCredit, 0)
	}

	if gotCode != "" {
		t.Errorf("Fail: code got %s want %s", gotCode, "")
	}

	if gotError != wantError {
		t.Errorf("Fail: error got %v want %v", gotError, wantError)
	}
}

// TestDeliveryReceipts
func TestDeliveryReceipts(t *testing.T) {
	onDeliveryReceipt := func(got clockwork.Receipt) {
		var wantID = "LA_424242"
		var wantTo = "441234567890"

		var wantStatus = clockwork.Delivered
		var wantErr error

		if got.ID != wantID {
			t.Errorf("Fail: id - got %s want %s", got.ID, wantID)
		}

		if got.To != wantTo {
			t.Errorf("Fail: to - got %s want %s", got.To, wantTo)
		}

		if got.Status != wantStatus {
			t.Errorf("Fail: status - got %v want %v", got.Status, wantStatus)
		}

		if got.Err != wantErr {
			t.Errorf("Fail: err - got %v want %v", got.Err, wantErr)
		}
	}

	testPath := "/receipts"
	testPort := 9090
	testQueryParams := "msg_id=LA_424242&status=DELIVRD&detail=0&to=441234567890"

	testURL := fmt.Sprintf("http://localhost:%d%s?%s", testPort, testPath, testQueryParams)

	go func() {
		clockwork.DeliveryReceiptListen(&clockwork.ReceiptHandler{
			Path:     testPath,
			Port:     testPort,
			Callback: onDeliveryReceipt,
		})
	}()

	testCases := []struct {
		method string
		url    string
		buf    *bytes.Buffer
	}{
		{
			method: "GET",
			url:    testURL,
			buf:    bytes.NewBufferString(""),
		},
		{
			method: "POST",
			url:    testURL,
			buf:    bytes.NewBufferString(testQueryParams),
		},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(tc.method, tc.url, tc.buf)
		if err != nil {
			t.Errorf("Fail: err %v", err)
		}
		client := &http.Client{}
		client.Do(req)
	}
}

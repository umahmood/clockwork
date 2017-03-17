package clockwork

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// DeliveryState status of a message
type DeliveryState byte

const (
	// Queued for delivery to mobile networks
	Queued DeliveryState = iota
	// Enroute sent to mobile network
	Enroute
	// Delivered to destination
	Delivered
	// Expired Message validity period has expired. Handset turned off or out of
	// range
	Expired
	// Deleted Message has been deleted
	Deleted
	// Undelivered Message could not be delivered
	Undelivered
	// Accepted Message is in accepted state. Message has been read manually on
	// behalf of the subscriber by customer service
	Accepted
	// Unknown No final delivery status received from the network
	Unknown
	// Rejected Message rejected by the mobile network
	Rejected
)

// String returns delivery state as a string
func (d DeliveryState) String() string {
	switch d {
	case Queued:
		return "Queued"
	case Enroute:
		return "Enroute"
	case Delivered:
		return "Delivered"
	case Expired:
		return "Expired"
	case Deleted:
		return "Deleted"
	case Undelivered:
		return "Undelivered"
	case Accepted:
		return "Accepted"
	case Unknown:
		return "Unknown"
	case Rejected:
		return "Rejected"
	default:
		return "!ERROR!" // debug
	}
}

// Receipt message delivery receipt
type Receipt struct {
	ID     string
	To     string
	Status DeliveryState
	Time   time.Time
	Err    error
}

// ReceiptCallback callback for new delivery receipt
type ReceiptCallback func(Receipt)

// ReceiptHandler custom handler which processes delivery receipts
type ReceiptHandler struct {
	// Path prefixed with "/" i.e. "/delivery-receipts", "/sms"
	Path string
	// Port to listen on
	Port int
	// Callback called when a new receipt has been received
	Callback ReceiptCallback
}

// ServeHTTP process delivery receipts
func (rc *ReceiptHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vals := url.Values{}
	if r.Method == "GET" {
		vals = r.URL.Query()
	}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err)
			return
		}

		vals, err = url.ParseQuery(string(body))
		if err != nil {
			log.Println(err)
			return
		}
	}

	if len(vals) == 0 {
		log.Println("serve http: query values map is zero")
		return
	}

	var receipt Receipt
	if v, ok := vals["msg_id"]; ok {
		receipt.ID = v[0]
	}

	if v, ok := vals["to"]; ok {
		receipt.To = v[0]
	}

	if v, ok := vals["status"]; ok {
		receipt.Status = toDeliveryState(v[0])
	}

	if v, ok := vals["detail"]; ok {
		receipt.Err = errorFromDetailCode(v[0])
	}

	receipt.Time = time.Now().UTC()

	// invoke callback
	rc.Callback(receipt)

	// need to respond with a 200 OK status code to acknowledge receipt of the
	// message, otherwise the clockwork API will retry at regular intervals.
	w.WriteHeader(http.StatusOK)
}

// DeliveryReceiptListen listen for incoming delivery receipts.
func DeliveryReceiptListen(rh *ReceiptHandler) error {
	if rh.Callback == nil {
		return errors.New("callback is nil")
	}
	if rh.Path == "" {
		return errors.New("path is empty")
	}
	mux := http.NewServeMux()
	mux.Handle(rh.Path, rh)
	p := fmt.Sprintf(":%d", rh.Port)
	return http.ListenAndServe(p, mux)
}

// toDeliveryState converts clockwork delivery status strings to the
// DeliveryState type. See:
// https://www.clockworksms.com/doc/reference/faqs/delivery-states/
func toDeliveryState(s string) DeliveryState {
	switch s {
	case "QUEUED":
		return Queued
	case "ENROUTE":
		return Enroute
	case "DELIVRD":
		return Delivered
	case "EXPIRED":
		return Expired
	case "DELETED":
		return Deleted
	case "UNDELIV":
		return Undelivered
	case "ACCEPTD":
		return Accepted
	case "UNKNOWN":
		return Unknown
	case "REJECTD":
		return Rejected
	default:
		return 126 // debug
	}
}

// errorFromDetailCode provides more information on why a message has failed.
// Sometimes the mobile network doesn’t provide a reason, in these cases it will
// be set to nil. See:
// https://www.clockworksms.com/doc/reference/faqs/delivery-states/
func errorFromDetailCode(code string) error {
	switch code {
	case "2":
		return ErrMessageDetailsWrong
	case "3":
		return ErrPermOperator
	case "4":
		return ErrTempOperator
	case "5":
		return ErrPermAbsentSub
	case "6":
		return ErrTempAbsentSub
	case "9":
		return ErrPermPhone
	case "10":
		return ErrTempPhone
	default:
		return nil
	}
}

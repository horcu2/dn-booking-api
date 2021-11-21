package models

import (
	"math/rand"
	"time"
)

type Consultation int64
type Conversation int64
type BookingStatus int64
type BookingPaymentStatus int64

const (
	Online Consultation = iota
	Visit
)

func (s Consultation) String() string {
	switch s {
	case Online:
		return "online"
	}
	return "visit"
}

const (
	Videocall Conversation = iota
	Voicecall
	Chat
)

func (s Conversation) String() string {
	switch s {
	case Videocall:
		return "video"
	case Voicecall:
		return "video"
	}
	return "chat"
}

const (
	Confirm BookingStatus = iota
	Waiting
	Cancelled
	Start
	Stop
	Complete
)

func (s BookingStatus) String() string {
	switch s {
	case Confirm:
		return "confirm"
	case Waiting:
		return "waiting"
	case Cancelled:
		return "cancelled"
	case Start:
		return "start"
	case Stop:
		return "stop"
	}
	return "complete"
}

const (
	Pending BookingPaymentStatus = iota
	Success
)

func (s BookingPaymentStatus) String() string {
	switch s {
	case Pending:
		return "pending"
	}
	return "success"
}

type Booking struct {
	Id                   int                  `json:"id"`
	Consultation         Consultation         `json:"consultation"`
	Conversation         Conversation         `json:"conversation"`
	BookingStatus        BookingStatus        `json:"booking-status"`
	BookingPaymentStatus BookingPaymentStatus `json:"booking-payment-status"`
	Patient              Patient              `json:"Patient"`
	Doctor               Doctor               `json:"Doctor"`
	Address              string               `json:"address"`
	AppointmentTime      time.Time			  `json:"appointment_time"`
	EstimatedCost		 float64				`json:"estimated_cost"`
	CancelReason		 string					`json:"cancel_reason"`
	Comments			 []Comment      `json:"comments"`
	Created 			 time.Time   `json:"created"`
	Updated			time.Time   `json:"updated"`
    Service			Service  `json:"service"`
	BookingPatient        Patient `json:"booking_patient"`
	Mobile				int64 `json:"mobile"`
	DoctorCost			float64 `json:"doctor_cost"`

	//properties below are for the Doctor object not booking
	//AdminCommission      float64 `json:"admin_commission"`
	//ProfessionalStatement string	`json:"professional_statement"`
	//Sex                   Sex `json:"sex"`
	//AvailableTimes        []time.Time `json:"available_times"`

}

func (booking Booking) EmergencyRequestCode() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 3)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (booking Booking) GetScheduledBookings() []Booking {
	return []Booking{}
}
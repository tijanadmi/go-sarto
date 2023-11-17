package models

import (
	"time"
)

// User is the user model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Phone     string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Room is the room model
type Room struct {
	ID        int
	RoomNameSr  string
	RoomNameEn  string
	RoomNameBg  string
	RoomShortDescSr  string
	RoomShortDescEn  string
	RoomShortDescBg  string
	RoomDescSr  string
	RoomDescEn  string
	RoomDescBg  string
	RoomPictureFolder  string
	RoomGuestNumber  string
	RoomPriceEn  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restriction is the restriction model
type Restriction struct {
	ID              int
	RestrictionNameSr string
	RestrictionNameEn string
	RestrictionNameBg string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservation is the reservation model
type Reservation struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Room      Room
	Processed int
}

// RoomRestriction is the room restriction model
type RoomRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	RoomID        int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Room          Room
	Reservation   Reservation
	Restriction   Restriction
}

//MailData holds an email message
type MailData struct {
	To      string
	From    string
	Subject string
	Content string
	Template string
}

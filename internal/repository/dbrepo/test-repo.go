package dbrepo

import (
	"errors"
	"log"
	"time"

	"github.com/Anahita-ghloo/foyer_international/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// If room_id is 2 it will fail
	if res.RoomID == 2 {
		return 0, errors.New("some error insert reser")
	}

	return 1, nil
}

// InsertRoomRestriction inserts a RoomRestriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	if r.RoomID == 1000 {
		return errors.New("some error insert rest")
	}
	return nil
}

// SearchAvailabilityByDates returns true if availability exists for roomID returns false if no availability existe
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	if roomID == 1000 {
		return false, errors.New("some error search")
	}

	// set up a test time
	layout := "2006-01-02"
	str := "2049-12-31"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	// this is our test to fail the query -- specify 2060-01-01 as start
	testDateToFail, err := time.Parse(layout, "2060-01-01")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return false, errors.New("some error search2")
	}

	// if the start date is after 2049-12-31, then return false,
	// indicating no availability;
	if start.After(t) {
		return false, nil
	}

	// otherwise, we have availability
	return true, nil

}

// SearchAvailabilityForAllRooms returns a slice of available rooms if any for the given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	// if the start date is after 2049-12-31, then return empty slice,
	// indicating no rooms are available;
	layout := "2006-01-02"
	str := "2049-12-31"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	testDateToFail, err := time.Parse(layout, "2060-01-01")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return rooms, errors.New("some error search avail")
	}

	if start.After(t) {
		return rooms, nil
	}

	// otherwise, put an entry into the slice, indicating that some room is
	// available for search dates
	room := models.Room{
		ID: 1,
	}
	rooms = append(rooms, room)

	return rooms, nil
}

// Get a room by ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {

	var room models.Room
	if id > 2 {
		return room, errors.New("some error search avail2")
	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	if email == "me@here.ca" {
		return 1, "", nil
	}
	return 0, "", errors.New("some error authe")
}

func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {

	var reservations []models.Reservation

	return reservations, nil
}

// AllNewReservation returns a slice of all reservations
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {

	var reservations []models.Reservation

	return reservations, nil
}

// GetReservationByID returns on e reservation by ID
func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {

	var res models.Reservation
	return res, nil
}

// UpdateReservation, updates a reservation in the database.
func (m *testDBRepo) UpdateReservation(u models.Reservation) error {

	return nil
}

// DeleteReservation deletes one reservation by id
func (m *testDBRepo) DeleteReservation(id int) error {

	return nil
}

func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {

	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}

// GetRestrictionsForRoomByDate, returns restrictions for a room by date range
func (m *testDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error) {

	var restrictions []models.RoomRestriction

	return restrictions, nil

}

// InsertBlockForRoom inserts a room restriction
func (m *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	return nil
}

// Delete block for a room by given ID
func (m *testDBRepo) DeleteBlockByID(id int) error {
	return nil
}

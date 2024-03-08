package controllers

import (
	m "UTS/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := gorm_connect(w)

	var rooms []m.Rooms
	result := db.Find(&rooms)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			sendErrorResponse(w, 404, "Rooms tidak ditemukan")
		} else {
			log.Println(result.Error)
			sendErrorResponse(w, 500, "Kesalahan internal server")
		}
		return
	}

	sendGetRoomsResponse(w, 200, "Berhasil Get Data", rooms)
}

func GetAllRoomDetails(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db := connect(w)
	defer db.Close()

	// Define the query to fetch room details and participants
	query := `
	SELECT r.id, r.room_name, p.id AS participant_id, p.id_account, u.username
	FROM rooms r
	INNER JOIN participants p ON r.id = p.id_room
	INNER JOIN accounts u ON p.id_account = u.id;
	`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 500, "internal error")
		return
	}
	defer rows.Close()

	// Initialize variables
	var account m.Accounts
	var room m.Rooms
	var participant m.Participants
	// var roomDetails []m.Rooms

	// Loop through each row and populate structures
	for rows.Next() {
		err := rows.Scan(&room.Id, &room.Room_Name, &participant.Id, &participant.ID_Account, &account.Username)
		if err != nil {
			log.Println(err)
			sendErrorResponse(w, 500, "internal error")
			return
		}
	}
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	// Connect to database
	db := gorm_connect(w)

	// Get participant data from request body
	var participant m.Participants
	err := json.NewDecoder(r.Body).Decode(&participant)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, 400, "Bad request, format data tidak sesuai")
		return
	}

	// Get ID_Game from Rooms
	var idGame int
	db.Table("rooms").Select("id_game").Where("id = ?", participant.ID_Room).Scan(&idGame)

	// Get MaxPlayer
	var maxPlayer int64
	db.Table("games").Select("max_player").Where("id = ?", idGame).Scan(&maxPlayer)

	// Count participants in the room
	var count int64
	db.Table("participants").Where("room_id = ?", participant.ID_Room).Count(&count)

	// Fail if MaxPlayer is reached
	if count >= maxPlayer {
		sendErrorResponse(w, 400, "Jumlah pemain dalam room sudah mencapai batas maksimal.")
		return
	}

	// Insert participant data
	result := db.Create(&participant)
	if result.Error != nil {
		log.Println(result.Error)
		sendErrorResponse(w, 500, "Gagal memasukan ke room")
		return
	}

	// Success response
	sendSuccessResponse(w, 201, "Berhasil masuk room")
}

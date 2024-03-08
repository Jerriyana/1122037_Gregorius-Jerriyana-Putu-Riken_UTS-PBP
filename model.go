package models

// Model Accounts
type Accounts struct {
	Id       int `json:"id"`
	Username int `json:"Username"`
}
type AccountsResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Accounts `json:"data"`
}

// Model Games
type Games struct {
	Id         int `json:"id"`
	Name       int `json:"Name"`
	Max_Player int `json:"Max_Player"`
}
type GamesResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Games `json:"data"`
}

// Model Participants
type Participants struct {
	Id         int `json:"id"`
	ID_Room    int `json:"ID_Room"`
	ID_Account int `json:"ID_Account"`
}
type ParticipantsResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    []Participants `json:"data"`
}

// Model Rooms
type Rooms struct {
	Id        int    `json:"id"`
	Room_Name string `json:"Room_Name"`
	Id_game   int    `json:"id_game"`
}
type RoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}

// Error response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Success response
type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Login Response
type LoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

package cache

const (
	ImUserOffline = 0
	ImUserOnline = 1
)

type UserLoginInfo struct {
	Status int8
	LoginToken string
	DeviceToken string
	LinkToken string
}


type CacheDao interface {
	GetRoomNumbers(roomId int64) (map[string]int, error)
	AddRoomNumbers(roomId int64, members map[string]int) error
	RemoveNumbers(roomId int64, members []string) error
	GetRoomNumbersByRang(roomId int64, min, max interface{}, limit int) (map[string]int, error)
}


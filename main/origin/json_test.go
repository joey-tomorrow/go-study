package origin

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func Test_Json(t *testing.T) {
	_ = json.NewEncoder(os.Stdout).Encode(
		&MyUser{1, "Ken", time.Now()},
	)
}

func (u *MyUser) MarshalJSON1() ([]byte, error) {
	return json.Marshal(&struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		LastSeen int64  `json:"lastSeen"`
	}{
		ID:       u.ID,
		Name:     u.Name,
		LastSeen: u.LastSeen.Unix(),
	})
}

func (u *MyUser) MarshalJSON() ([]byte, error) {
	type Alias MyUser
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		LastSeen: u.LastSeen.Unix(),
		Alias:    (*Alias)(u),
	})
}

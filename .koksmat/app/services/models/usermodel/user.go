/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package usermodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/nexi-tools/database/databasetypes"
)

func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type User struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Email string `json:"email"`
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
    Language_id int `json:"language_id"`
    Country_id int `json:"country_id"`
    Region_id int `json:"region_id"`
    Status string `json:"status"`

}


/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package iconmodel
import (
	"encoding/json"
	"time"
    // 
)

func UnmarshalIcon(data []byte) (Icon, error) {
	var r Icon
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Icon) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Icon struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Icon string `json:"icon"`
    Base64 string `json:"base64"`

}


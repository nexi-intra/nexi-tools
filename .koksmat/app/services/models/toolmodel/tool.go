/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package toolmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/nexi-tools/database/databasetypes"
)

func UnmarshalTool(data []byte) (Tool, error) {
	var r Tool
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Tool) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Tool struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Category_id int `json:"category_id"`

}


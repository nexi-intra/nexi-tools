/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package translationmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/nexi-tools/database/databasetypes"
)

func UnmarshalTranslation(data []byte) (Translation, error) {
	var r Translation
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Translation) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Translation struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Language_id int `json:"language_id"`
    Translation string `json:"translation"`
    Tablename string `json:"tablename"`
    Fieldname string `json:"fieldname"`
    Recordid int `json:"recordid"`

}


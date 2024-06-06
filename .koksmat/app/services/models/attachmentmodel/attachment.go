/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v1
package attachmentmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/nexi-tools/database/databasetypes"
)

func UnmarshalAttachment(data []byte) (Attachment, error) {
	var r Attachment
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Attachment) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Attachment struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Sortorder string `json:"sortorder"`
    Filename string `json:"filename"`
    Link string `json:"link"`
    Mimetype string `json:"mimetype"`
    Tool_id int `json:"tool_id"`

}


/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateMapModel v1.1
package applogic
import (
	//"encoding/json"
	//"time"
	"github.com/magicbutton/nexi-tools/database"
	"github.com/magicbutton/nexi-tools/services/models/toolmodel"
   
)


func MapToolOutgoing(db database.Tool) toolmodel.Tool {
    return toolmodel.Tool{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                Category_id : db.Category_id,
        Url : db.Url,
        Status : db.Status,
                Icon_id : db.Icon_id,

    }
}

func MapToolIncoming(in toolmodel.Tool) database.Tool {
    return database.Tool{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                Category_id : in.Category_id,
        Url : in.Url,
        Status : in.Status,
                Icon_id : in.Icon_id,
        Searchindex : in.Name,

    }
}

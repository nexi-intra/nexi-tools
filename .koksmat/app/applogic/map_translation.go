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
	"github.com/magicbutton/nexi-tools/services/models/translationmodel"
   
)


func MapTranslationOutgoing(db database.Translation) translationmodel.Translation {
    return translationmodel.Translation{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                Language_id : db.Language_id,
        Translation : db.Translation,
        Tablename : db.Tablename,
        Fieldname : db.Fieldname,

    }
}

func MapTranslationIncoming(in translationmodel.Translation) database.Translation {
    return database.Translation{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                Language_id : in.Language_id,
        Translation : in.Translation,
        Tablename : in.Tablename,
        Fieldname : in.Fieldname,
        Searchindex : in.Name,

    }
}

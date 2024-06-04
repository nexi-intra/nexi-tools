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
	"github.com/magicbutton/nexi-tools/services/models/countrymodel"
   
)


func MapCountryOutgoing(db database.Country) countrymodel.Country {
    return countrymodel.Country{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                Region_id : db.Region_id,

    }
}

func MapCountryIncoming(in countrymodel.Country) database.Country {
    return database.Country{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                Region_id : in.Region_id,
        Searchindex : in.Name,

    }
}

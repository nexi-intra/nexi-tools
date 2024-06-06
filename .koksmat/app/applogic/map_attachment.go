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
	"github.com/magicbutton/nexi-tools/services/models/attachmentmodel"
   
)


func MapAttachmentOutgoing(db database.Attachment) attachmentmodel.Attachment {
    return attachmentmodel.Attachment{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Sortorder : db.Sortorder,
        Filename : db.Filename,
        Link : db.Link,
        Mimetype : db.Mimetype,
                Tool_id : db.Tool_id,

    }
}

func MapAttachmentIncoming(in attachmentmodel.Attachment) database.Attachment {
    return database.Attachment{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Sortorder : in.Sortorder,
        Filename : in.Filename,
        Link : in.Link,
        Mimetype : in.Mimetype,
                Tool_id : in.Tool_id,
        Searchindex : in.Name,

    }
}

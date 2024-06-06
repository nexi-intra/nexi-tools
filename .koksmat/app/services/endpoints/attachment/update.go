/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package attachment

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/attachmentmodel"

)

func AttachmentUpdate(item attachmentmodel.Attachment) (*attachmentmodel.Attachment, error) {
    log.Println("Calling Attachmentupdate")

    return applogic.Update[database.Attachment, attachmentmodel.Attachment](item.ID,item, applogic.MapAttachmentIncoming, applogic.MapAttachmentOutgoing)

}

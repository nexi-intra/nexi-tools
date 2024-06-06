/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package attachment

import (
    "log"
   
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/attachmentmodel"

)

func AttachmentCreate(item attachmentmodel.Attachment) (*attachmentmodel.Attachment, error) {
    log.Println("Calling Attachmentcreate")

    return applogic.Create[database.Attachment, attachmentmodel.Attachment](item, applogic.MapAttachmentIncoming, applogic.MapAttachmentOutgoing)

}

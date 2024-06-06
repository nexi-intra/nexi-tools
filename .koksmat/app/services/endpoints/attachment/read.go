/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package attachment

import (
    "log"
    "strconv"
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/attachmentmodel"

)

func AttachmentRead(arg0 string) (*attachmentmodel.Attachment, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Attachmentread")

    return applogic.Read[database.Attachment, attachmentmodel.Attachment](id, applogic.MapAttachmentOutgoing)

}

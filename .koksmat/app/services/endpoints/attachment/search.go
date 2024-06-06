/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package attachment

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/attachmentmodel"
    . "github.com/magicbutton/nexi-tools/utils"
)

func AttachmentSearch(query string) (*Page[attachmentmodel.Attachment], error) {
    log.Println("Calling Attachmentsearch")

    return applogic.Search[database.Attachment, attachmentmodel.Attachment]("searchindex", query, applogic.MapAttachmentOutgoing)

}

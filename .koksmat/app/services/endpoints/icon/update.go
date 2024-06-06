/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package icon

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/iconmodel"

)

func IconUpdate(item iconmodel.Icon) (*iconmodel.Icon, error) {
    log.Println("Calling Iconupdate")

    return applogic.Update[database.Icon, iconmodel.Icon](item.ID,item, applogic.MapIconIncoming, applogic.MapIconOutgoing)

}

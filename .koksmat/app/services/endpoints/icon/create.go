/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package icon

import (
    "log"
   
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/iconmodel"

)

func IconCreate(item iconmodel.Icon) (*iconmodel.Icon, error) {
    log.Println("Calling Iconcreate")

    return applogic.Create[database.Icon, iconmodel.Icon](item, applogic.MapIconIncoming, applogic.MapIconOutgoing)

}

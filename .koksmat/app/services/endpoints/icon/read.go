/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package icon

import (
    "log"
    "strconv"
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/iconmodel"

)

func IconRead(arg0 string) (*iconmodel.Icon, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Iconread")

    return applogic.Read[database.Icon, iconmodel.Icon](id, applogic.MapIconOutgoing)

}

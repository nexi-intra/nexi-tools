/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package tool

import (
    "log"
   
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/toolmodel"

)

func ToolCreate(item toolmodel.Tool) (*toolmodel.Tool, error) {
    log.Println("Calling Toolcreate")

    return applogic.Create[database.Tool, toolmodel.Tool](item, applogic.MapToolIncoming, applogic.MapToolOutgoing)

}

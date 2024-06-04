/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package tool

import (
    "log"
    "strconv"
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/toolmodel"

)

func ToolRead(arg0 string) (*toolmodel.Tool, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Toolread")

    return applogic.Read[database.Tool, toolmodel.Tool](id, applogic.MapToolOutgoing)

}

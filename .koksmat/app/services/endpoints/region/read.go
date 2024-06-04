/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package region

import (
    "log"
    "strconv"
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/regionmodel"

)

func RegionRead(arg0 string) (*regionmodel.Region, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Regionread")

    return applogic.Read[database.Region, regionmodel.Region](id, applogic.MapRegionOutgoing)

}

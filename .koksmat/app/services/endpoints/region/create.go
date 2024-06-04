/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package region

import (
    "log"
   
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/regionmodel"

)

func RegionCreate(item regionmodel.Region) (*regionmodel.Region, error) {
    log.Println("Calling Regioncreate")

    return applogic.Create[database.Region, regionmodel.Region](item, applogic.MapRegionIncoming, applogic.MapRegionOutgoing)

}

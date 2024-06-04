/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package region

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/regionmodel"
    . "github.com/magicbutton/nexi-tools/utils"
)

func RegionSearch(query string) (*Page[regionmodel.Region], error) {
    log.Println("Calling Regionsearch")

    return applogic.Search[database.Region, regionmodel.Region]("searchindex", query, applogic.MapRegionOutgoing)

}

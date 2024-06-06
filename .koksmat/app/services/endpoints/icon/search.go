/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package icon

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/iconmodel"
    . "github.com/magicbutton/nexi-tools/utils"
)

func IconSearch(query string) (*Page[iconmodel.Icon], error) {
    log.Println("Calling Iconsearch")

    return applogic.Search[database.Icon, iconmodel.Icon]("searchindex", query, applogic.MapIconOutgoing)

}

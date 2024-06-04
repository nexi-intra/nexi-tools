/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package tool

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/toolmodel"
    . "github.com/magicbutton/nexi-tools/utils"
)

func ToolSearch(query string) (*Page[toolmodel.Tool], error) {
    log.Println("Calling Toolsearch")

    return applogic.Search[database.Tool, toolmodel.Tool]("searchindex", query, applogic.MapToolOutgoing)

}

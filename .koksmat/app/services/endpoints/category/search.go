/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package category

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/categorymodel"
    . "github.com/magicbutton/nexi-tools/utils"
)

func CategorySearch(query string) (*Page[categorymodel.Category], error) {
    log.Println("Calling Categorysearch")

    return applogic.Search[database.Category, categorymodel.Category]("searchindex", query, applogic.MapCategoryOutgoing)

}

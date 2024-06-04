/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package category

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/categorymodel"

)

func CategoryUpdate(item categorymodel.Category) (*categorymodel.Category, error) {
    log.Println("Calling Categoryupdate")

    return applogic.Update[database.Category, categorymodel.Category](item.ID,item, applogic.MapCategoryIncoming, applogic.MapCategoryOutgoing)

}

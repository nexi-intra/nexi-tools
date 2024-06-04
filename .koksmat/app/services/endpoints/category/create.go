/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package category

import (
    "log"
   
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/categorymodel"

)

func CategoryCreate(item categorymodel.Category) (*categorymodel.Category, error) {
    log.Println("Calling Categorycreate")

    return applogic.Create[database.Category, categorymodel.Category](item, applogic.MapCategoryIncoming, applogic.MapCategoryOutgoing)

}

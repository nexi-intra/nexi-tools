/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package translation

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/translationmodel"

)

func TranslationUpdate(item translationmodel.Translation) (*translationmodel.Translation, error) {
    log.Println("Calling Translationupdate")

    return applogic.Update[database.Translation, translationmodel.Translation](item.ID,item, applogic.MapTranslationIncoming, applogic.MapTranslationOutgoing)

}

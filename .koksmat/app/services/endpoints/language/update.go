/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package language

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/languagemodel"

)

func LanguageUpdate(item languagemodel.Language) (*languagemodel.Language, error) {
    log.Println("Calling Languageupdate")

    return applogic.Update[database.Language, languagemodel.Language](item.ID,item, applogic.MapLanguageIncoming, applogic.MapLanguageOutgoing)

}

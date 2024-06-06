/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package language

import (
    "log"
   
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/languagemodel"

)

func LanguageCreate(item languagemodel.Language) (*languagemodel.Language, error) {
    log.Println("Calling Languagecreate")

    return applogic.Create[database.Language, languagemodel.Language](item, applogic.MapLanguageIncoming, applogic.MapLanguageOutgoing)

}

/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package translation

import (
    "log"
    "strconv"
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/translationmodel"

)

func TranslationRead(arg0 string) (*translationmodel.Translation, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Translationread")

    return applogic.Read[database.Translation, translationmodel.Translation](id, applogic.MapTranslationOutgoing)

}

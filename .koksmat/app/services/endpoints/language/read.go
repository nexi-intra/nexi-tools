/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package language

import (
    "log"
    "strconv"
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/languagemodel"

)

func LanguageRead(arg0 string) (*languagemodel.Language, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Languageread")

    return applogic.Read[database.Language, languagemodel.Language](id, applogic.MapLanguageOutgoing)

}

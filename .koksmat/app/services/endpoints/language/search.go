/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package language

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/languagemodel"
    . "github.com/magicbutton/nexi-tools/utils"
)

func LanguageSearch(query string) (*Page[languagemodel.Language], error) {
    log.Println("Calling Languagesearch")

    return applogic.Search[database.Language, languagemodel.Language]("searchindex", query, applogic.MapLanguageOutgoing)

}

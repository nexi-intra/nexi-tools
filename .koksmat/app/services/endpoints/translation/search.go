/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package translation

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/translationmodel"
    . "github.com/magicbutton/nexi-tools/utils"
)

func TranslationSearch(query string) (*Page[translationmodel.Translation], error) {
    log.Println("Calling Translationsearch")

    return applogic.Search[database.Translation, translationmodel.Translation]("searchindex", query, applogic.MapTranslationOutgoing)

}

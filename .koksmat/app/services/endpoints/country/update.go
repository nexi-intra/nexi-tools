/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package country

import (
    "log"

    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/countrymodel"

)

func CountryUpdate(item countrymodel.Country) (*countrymodel.Country, error) {
    log.Println("Calling Countryupdate")

    return applogic.Update[database.Country, countrymodel.Country](item.ID,item, applogic.MapCountryIncoming, applogic.MapCountryOutgoing)

}

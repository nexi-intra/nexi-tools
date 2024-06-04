/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package country

import (
    "log"
    "strconv"
    "github.com/magicbutton/nexi-tools/applogic"
    "github.com/magicbutton/nexi-tools/database"
    "github.com/magicbutton/nexi-tools/services/models/countrymodel"

)

func CountryRead(arg0 string) (*countrymodel.Country, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Countryread")

    return applogic.Read[database.Country, countrymodel.Country](id, applogic.MapCountryOutgoing)

}

    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/koksmat/useservice";
    import { useState } from "react";
    import {RegionForm} from "./form";
    
    import {RegionItem} from "../applogic/model";
    export default function CreateRegion() {
       
        const region = {} as RegionItem;
        return (
          <div>{region && 
          <RegionForm region={region} editmode="create"/>}
         
          </div>
        );
    }

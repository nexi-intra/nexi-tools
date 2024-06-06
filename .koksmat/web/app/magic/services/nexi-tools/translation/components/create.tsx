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
    import {TranslationForm} from "./form";
    
    import {TranslationItem} from "../applogic/model";
    export default function CreateTranslation() {
       
        const translation = {} as TranslationItem;
        return (
          <div>{translation && 
          <TranslationForm translation={translation} editmode="create"/>}
         
          </div>
        );
    }

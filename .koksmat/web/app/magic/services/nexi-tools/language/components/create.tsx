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
    import {LanguageForm} from "./form";
    
    import {LanguageItem} from "../applogic/model";
    export default function CreateLanguage() {
       
        const language = {} as LanguageItem;
        return (
          <div>{language && 
          <LanguageForm language={language} editmode="create"/>}
         
          </div>
        );
    }

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
import {LanguageItem} from "../applogic/model";


/* yankiebar */

export default function ReadLanguage(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<LanguageItem>(
      "nexi-tools.language",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const language = readResult.data;
    return (
      <div>
          
    {language && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{language.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{language.description}</div>
    </div>    <div>
        <div className="font-bold" >Code</div>
        <div>{language.code}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{language.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{language.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{language.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{language.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{language.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      

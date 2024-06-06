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
import {TranslationItem} from "../applogic/model";


/* yankiebar */

export default function ReadTranslation(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TranslationItem>(
      "nexi-tools.translation",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const translation = readResult.data;
    return (
      <div>
          
    {translation && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{translation.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{translation.description}</div>
    </div>    <div>
        <div className="font-bold" >Language</div>
        <div>{translation.language_id}</div>
    </div>    <div>
        <div className="font-bold" >Translation</div>
        <div>{translation.translation}</div>
    </div>    <div>
        <div className="font-bold" >Table Name</div>
        <div>{translation.tablename}</div>
    </div>    <div>
        <div className="font-bold" >Field Name</div>
        <div>{translation.fieldname}</div>
    </div>    <div>
        <div className="font-bold" >Record Id</div>
        <div>{translation.recordid}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{translation.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{translation.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{translation.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{translation.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{translation.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      

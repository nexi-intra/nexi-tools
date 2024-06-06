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
import {IconItem} from "../applogic/model";


/* yankiebar */

export default function ReadIcon(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<IconItem>(
      "nexi-tools.icon",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const icon = readResult.data;
    return (
      <div>
          
    {icon && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{icon.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{icon.description}</div>
    </div>    <div>
        <div className="font-bold" >Icon</div>
        <div>{icon.icon}</div>
    </div>    <div>
        <div className="font-bold" >Base64</div>
        <div>{icon.base64}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{icon.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{icon.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{icon.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{icon.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{icon.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      

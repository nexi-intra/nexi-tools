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
import {RegionItem} from "../applogic/model";


/* yankiebar */

export default function ReadRegion(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<RegionItem>(
      "nexi-tools.region",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const region = readResult.data;
    return (
      <div>
          
    {region && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{region.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{region.description}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{region.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{region.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{region.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{region.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{region.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      

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
import {ToolItem} from "../applogic/model";


/* yankiebar */

export default function ReadTool(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ToolItem>(
      "nexi-tools.tool",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const tool = readResult.data;
    return (
      <div>
          
    {tool && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{tool.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{tool.description}</div>
    </div>    <div>
        <div className="font-bold" >Category</div>
        <div>{tool.category_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{tool.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{tool.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{tool.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{tool.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{tool.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      

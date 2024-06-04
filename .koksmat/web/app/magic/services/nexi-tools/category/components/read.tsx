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
import {CategoryItem} from "../applogic/model";


/* yankiebar */

export default function ReadCategory(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<CategoryItem>(
      "nexi-tools.category",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const category = readResult.data;
    return (
      <div>
          
    {category && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{category.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{category.description}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{category.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{category.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{category.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{category.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{category.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      

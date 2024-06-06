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
import {AttachmentItem} from "../applogic/model";


/* yankiebar */

export default function ReadAttachment(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AttachmentItem>(
      "nexi-tools.attachment",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const attachment = readResult.data;
    return (
      <div>
          
    {attachment && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{attachment.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{attachment.description}</div>
    </div>    <div>
        <div className="font-bold" >Sort Order</div>
        <div>{attachment.sortorder}</div>
    </div>    <div>
        <div className="font-bold" >Filename</div>
        <div>{attachment.filename}</div>
    </div>    <div>
        <div className="font-bold" >Link</div>
        <div>{attachment.link}</div>
    </div>    <div>
        <div className="font-bold" >Mimetype</div>
        <div>{attachment.mimetype}</div>
    </div>    <div>
        <div className="font-bold" >Tool</div>
        <div>{attachment.tool_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{attachment.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{attachment.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{attachment.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{attachment.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{attachment.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      

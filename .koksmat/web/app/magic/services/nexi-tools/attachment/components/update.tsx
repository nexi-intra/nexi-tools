/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/koksmat/useservice";
import { useState } from "react";
import {AttachmentForm} from "./form";

import {AttachmentItem} from "../applogic/model";
export default function UpdateAttachment(props: { id: number }) {
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
      <div>{attachment && 
      <AttachmentForm attachment={attachment} editmode="update"/>}
     
      </div>
    );
}

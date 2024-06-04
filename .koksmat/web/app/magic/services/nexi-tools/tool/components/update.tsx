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
import {ToolForm} from "./form";

import {ToolItem} from "../applogic/model";
export default function UpdateTool(props: { id: number }) {
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
      <div>{tool && 
      <ToolForm tool={tool} editmode="update"/>}
     
      </div>
    );
}

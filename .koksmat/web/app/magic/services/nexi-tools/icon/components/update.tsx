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
import {IconForm} from "./form";

import {IconItem} from "../applogic/model";
export default function UpdateIcon(props: { id: number }) {
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
      <div>{icon && 
      <IconForm icon={icon} editmode="update"/>}
     
      </div>
    );
}

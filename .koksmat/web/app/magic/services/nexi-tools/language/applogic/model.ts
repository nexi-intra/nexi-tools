/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/
"use client";
import { z } from "zod";
// spunk
// Language
export interface LanguageItem {
  id: number;
  created_at: string;
  created_by: string;
  updated_at: string;
  updated_by: string;
  name: string;
  description: string;
  code: string;
}

// Language
export const LanguageSchema = z.object({
  name: z.string(),
  description: z.string().optional(),
  code: z.string(),
});

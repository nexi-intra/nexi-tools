/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.translation
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by character varying COLLATE pg_catalog."default"  ,

    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying COLLATE pg_catalog."default" ,

    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,searchindex character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,language_id int   NOT NULL
    ,translation character varying COLLATE pg_catalog."default"  NOT NULL
    ,tablename character varying COLLATE pg_catalog."default"  NOT NULL
    ,fieldname character varying COLLATE pg_catalog."default"  NOT NULL
    ,recordid character varying COLLATE pg_catalog."default"   NOT NULL


);

                ALTER TABLE IF EXISTS public.translation
                ADD FOREIGN KEY (language_id)
                REFERENCES public.language (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----

DROP TABLE public.translation;


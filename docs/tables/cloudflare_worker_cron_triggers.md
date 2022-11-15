# Table: cloudflare_worker_cron_triggers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| worker_meta_data_id | string | X | √ |  | 
| cron | string | X | √ |  | 
| created_on | timestamp | X | √ |  | 
| modified_on | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| cloudflare_worker_meta_data_selefra_id | string | X | X | fk to cloudflare_worker_meta_data.selefra_id | 



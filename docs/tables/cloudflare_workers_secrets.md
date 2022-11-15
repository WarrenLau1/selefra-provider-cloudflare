# Table: cloudflare_workers_secrets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| worker_meta_data_id | string | X | √ |  | 
| name | string | X | √ |  | 
| secret_text | string | X | √ |  | 
| cloudflare_worker_meta_data_selefra_id | string | X | X | fk to cloudflare_worker_meta_data.selefra_id | 
| selefra_id | string | √ | √ | random id | 



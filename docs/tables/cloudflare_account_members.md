# Table: cloudflare_account_members

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | random id | 
| cloudflare_accounts_selefra_id | string | X | X | fk to cloudflare_accounts.selefra_id | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| code | string | X | √ |  | 
| user | json | X | √ |  | 
| status | string | X | √ |  | 
| roles | json | X | √ |  | 



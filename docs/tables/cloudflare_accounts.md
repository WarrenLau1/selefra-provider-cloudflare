# Table: cloudflare_accounts

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| created_on | timestamp | X | √ |  | 
| settings | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 



# Table: cloudflare_worker_routes

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ | `The Account ID of the resource.` | 
| zone_id | string | X | √ | `Zone identifier tag.` | 
| id | string | √ | √ |  | 
| pattern | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| script | string | X | √ |  | 



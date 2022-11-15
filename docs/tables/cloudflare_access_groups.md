# Table: cloudflare_access_groups

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ | `The Account ID of the resource.` | 
| zone_id | string | X | √ | `Zone identifier tag.` | 
| id | string | √ | √ |  | 
| updated_at | timestamp | X | √ |  | 
| exclude | json | X | √ |  | 
| require | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_at | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| include | json | X | √ |  | 



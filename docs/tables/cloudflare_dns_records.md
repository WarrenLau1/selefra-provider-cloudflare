# Table: cloudflare_dns_records

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| zone_name | string | X | √ |  | 
| ttl | int | X | √ |  | 
| meta | json | X | √ | `Extra Cloudflare-specific information about the record.` | 
| created_on | timestamp | X | √ |  | 
| type | string | X | √ |  | 
| name | string | X | √ |  | 
| locked | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| zone_id | string | X | √ |  | 
| proxied | bool | X | √ |  | 
| proxiable | bool | X | √ |  | 
| data | json | X | √ | `Metadata about the record.` | 
| account_id | string | X | √ | `The Account ID of the resource.` | 
| modified_on | timestamp | X | √ |  | 
| content | string | X | √ |  | 
| id | string | √ | √ |  | 
| priority | int | X | √ |  | 



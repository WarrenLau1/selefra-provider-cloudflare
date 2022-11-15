# Table: cloudflare_waf_overrides

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | √ | √ |  | 
| urls | string_array | X | √ |  | 
| priority | int | X | √ |  | 
| rewrite_action | json | X | √ |  | 
| paused | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ | `The Account ID of the resource.` | 
| zone_id | string | X | √ | `Zone identifier tag.` | 
| description | string | X | √ |  | 
| groups | json | X | √ |  | 
| rules | json | X | √ |  | 



# Table: cloudflare_certificate_packs

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| hosts | string_array | X | √ |  | 
| certificates | json | X | √ |  | 
| primary_certificate | string | X | √ |  | 
| validation_errors | json | X | √ |  | 
| validation_records | json | X | √ |  | 
| account_id | string | X | √ | `The Account ID of the resource.` | 
| zone_id | string | X | √ | `Zone identifier tag.` | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 



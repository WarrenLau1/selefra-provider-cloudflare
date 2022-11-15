# Table: cloudflare_images

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| variants | string_array | X | √ |  | 
| uploaded | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ | `The Account ID of the resource.` | 
| id | string | √ | √ |  | 
| filename | string | X | √ |  | 
| metadata | json | X | √ |  | 
| require_signed_urls | bool | X | √ |  | 



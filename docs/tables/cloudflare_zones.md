# Table: cloudflare_zones

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| verification_key | string | X | √ |  | 
| account_id | string | X | √ | `The Account ID of the resource.` | 
| original_name_servers | string_array | X | √ |  | 
| plan | json | X | √ |  | 
| plan_pending | json | X | √ |  | 
| betas | string_array | X | √ |  | 
| created_on | timestamp | X | √ |  | 
| modified_on | timestamp | X | √ |  | 
| type | string | X | √ |  | 
| host | json | X | √ |  | 
| account | json | X | √ |  | 
| deactivation_reason | string | X | √ |  | 
| name | string | X | √ |  | 
| development_mode | int | X | √ |  | 
| original_dnshost | string | X | √ |  | 
| permissions | string_array | X | √ |  | 
| vanity_name_servers | string_array | X | √ |  | 
| paused | bool | X | √ |  | 
| meta | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | √ | √ |  | 
| original_registrar | string | X | √ |  | 
| name_servers | string_array | X | √ |  | 
| owner | json | X | √ |  | 
| status | string | X | √ |  | 



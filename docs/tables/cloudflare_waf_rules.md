# Table: cloudflare_waf_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| priority | string | X | √ |  | 
| cloudflare_waf_packages_selefra_id | string | X | X | fk to cloudflare_waf_packages.selefra_id | 
| allowed_modes | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| waf_package_id | string | X | √ |  | 
| description | string | X | √ |  | 
| package_id | string | X | √ |  | 
| group | json | X | √ |  | 
| mode | string | X | √ |  | 
| default_mode | string | X | √ |  | 



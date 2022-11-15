# Table: cloudflare_waf_groups

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cloudflare_waf_packages_selefra_id | string | X | X | fk to cloudflare_waf_packages.selefra_id | 
| rules_count | int | X | √ |  | 
| modified_rules_count | int | X | √ |  | 
| package_id | string | X | √ |  | 
| mode | string | X | √ |  | 
| allowed_modes | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| waf_package_id | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 



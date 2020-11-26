# 视图API

这是一个对视图进行创建、修改、检索、和删除的API

Contact:

Name: 杜锐

Email: durui@zdns.cn

## API

### /views

#### DELETE

删除视图

**Responses**

|Status|Description|
|---|---|
|200|删除成功|

**Response Schema**

Status Code 200

Body Type: application/json

|Field|Type|Required|Description|
|---|---|---|---|

**Response Example**

200 Response

```json
""
```
#### GET

查看视图信息

**Parameters**

|Field|In|Type|Required|Description|
|---|---|---|---|---|
|with_add|query|string|False|是否显示ADD视图信息|
|search_attrs[name|comment]|query|string|False|search_attrs[name|comment]|
|sort_col|query|string|False|按照某字段排序|
|sort_ord|query|string|False|sort_ord|
|page_num|query|integer|False|page_num|
|page_size|query|integer|False|page_size|

**Responses**

|Status|Description|
|---|---|
|200|查看成功|

**Response Schema**

Status Code 200

Body Type: application/json

|Field|Type|Required|Description|
|---|---|---|---|
|page_num|string|False||
|page_size|string|False||
|resources||False||
|total_size|integer|False||

**Response Example**

200 Response

```json
{
    "page_num": "string",
    "page_size": "string",
    "resources": "",
    "total_size": 0
}
```
#### POST

添加新视图信息

**Request Body**

Body Type: application/json

Body Parameter

|Field|Type|Required|Description|
|---|---|---|---|

Body Example

```json
""
```
**Responses**

|Status|Description|
|---|---|
|200|添加成功|

**Response Schema**

Status Code 200

Body Type: application/json

|Field|Type|Required|Description|
|---|---|---|---|

**Response Example**

200 Response

```json
""
```
#### PUT

更改视图的内容

**Request Body**

Body Type: application/json

Body Parameter

|Field|Type|Required|Description|
|---|---|---|---|

Body Example

```json
""
```
**Responses**

|Status|Description|
|---|---|
|200|更改成功|

**Response Schema**

Status Code 200

Body Type: application/json

|Field|Type|Required|Description|
|---|---|---|---|

**Response Example**

200 Response

```json
""
```
## Schemas

### view

**Properties**

|Field|Type|Required|Description|
|---|---|---|---|
|acls||False||
|allow_recursive|string|False||
|backup_query_sources||False||
|bind_ips||False||
|black_acls||False||
|comment|string|False||
|dns64s||False||
|ecs_exact_match|string|False||
|ecs_recurse_domains||False||
|fail_forwarder|string|False||
|filter_aaaa|string|False||
|filter_aaaa_exempt||False||
|filter_aaaa_ips||False||
|forward_zones||False||
|href|string|False||
|id|string|False||
|limit_ips||False||
|name|string|False||
|need_tsig_key|string|False||
|non_recursive_acls||False||
|owners||False||
|priority|integer|False||
|query_source|string|False||
|recursion_enable|string|False||
|recursive_acls||False||
|state|string|False||
|stub_zones||False||
|try_final_after_forward|string|False||
|tsig_algorithm|string|False||
|tsig_host||False||
|tsig_name|string|False||
|tsig_secret|string|False||
|working_query_source|string|False||
|zones||False||

**Example**

```json
{
    "acls": "",
    "allow_recursive": "string",
    "backup_query_sources": "",
    "bind_ips": "",
    "black_acls": "",
    "comment": "string",
    "dns64s": "",
    "ecs_exact_match": "string",
    "ecs_recurse_domains": "",
    "fail_forwarder": "string",
    "filter_aaaa": "string",
    "filter_aaaa_exempt": "",
    "filter_aaaa_ips": "",
    "forward_zones": "",
    "href": "string",
    "id": "string",
    "limit_ips": "",
    "name": "string",
    "need_tsig_key": "string",
    "non_recursive_acls": "",
    "owners": "",
    "priority": 0,
    "query_source": "string",
    "recursion_enable": "string",
    "recursive_acls": "",
    "state": "string",
    "stub_zones": "",
    "try_final_after_forward": "string",
    "tsig_algorithm": "string",
    "tsig_host": "",
    "tsig_name": "string",
    "tsig_secret": "string",
    "working_query_source": "string",
    "zones": ""
}
```

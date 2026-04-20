# 家谱管理系统 API 文档

## 基础信息

- **Base URL**: `http://localhost:8080/api`
- **Swagger文档**: `http://localhost:8080/swagger/index.html`
- **数据格式**: JSON
- **编码格式**: UTF-8

---

## 一、人员管理接口

### 1.1 创建人员
- **接口**: `POST /api/person/create`
- **描述**: 创建新的家族成员
- **请求体**:
```json
{
  "uuid": "string",
  "name": "string",
  "firstName": "string",
  "lastName": "string",
  "generationName": "string",
  "alias": "string",
  "englishName": "string",
  "gender": "string",
  "isAlive": true,
  "birthDate": "string",
  "birthLunar": "string",
  "birthPlace": "string",
  "birthPlaceLongitude": "string",
  "birthPlaceLatitude": "string",
  "deathDate": "string",
  "deathLunar": "string",
  "deathPlace": "string",
  "deathPlaceLongitude": "string",
  "deathPlaceLatitude": "string",
  "deathCause": "string",
  "generation": "string",
  "occupation": "string",
  "education": "string",
  "biography": "string",
  "avatarUrl": "string",
  "coverPhotoUrl": "string",
  "isPublic": true,
  "verificationStatus": "string",
  "createdBy": "string"
}
```
- **响应**: `201 Created` 返回人员对象

### 1.2 更新人员信息
- **接口**: `PUT /api/person/update`
- **描述**: 更新家族成员信息
- **请求体**: 同创建人员
- **响应**: `200 OK` 返回更新后的人员对象

### 1.3 删除人员
- **接口**: `DELETE /api/person/delete`
- **描述**: 删除家族成员
- **请求参数**:
  - `id` (query, required): 人员ID
- **响应**: `204 No Content`

### 1.4 根据ID获取人员
- **接口**: `GET /api/person/getById`
- **描述**: 根据ID获取家族成员详细信息
- **请求参数**:
  - `id` (query, required): 人员ID
- **响应**: `200 OK` 返回人员对象

### 1.5 根据UUID获取人员
- **接口**: `GET /api/person/getByUuid`
- **描述**: 根据UUID获取家族成员信息
- **请求参数**:
  - `uuid` (query, required): 人员UUID
- **响应**: `200 OK` 返回人员对象

### 1.6 列出人员
- **接口**: `GET /api/person/list`
- **描述**: 分页列出所有家族成员
- **请求参数**:
  - `page` (query, optional, default=1): 页码
  - `pageSize` (query, optional, default=10): 每页数量
- **响应**:
```json
{
  "data": [],
  "total": 100,
  "page": 1,
  "size": 10
}
```

### 1.7 搜索人员
- **接口**: `GET /api/person/search`
- **描述**: 根据姓名搜索家族成员
- **请求参数**:
  - `name` (query, required): 搜索姓名
  - `page` (query, optional, default=1): 页码
  - `pageSize` (query, optional, default=10): 每页数量
- **响应**: 同列出人员接口

---

## 二、亲子关系管理接口

### 2.1 创建亲子关系
- **接口**: `POST /api/relationship/createParentChild`
- **描述**: 创建父母与子女之间的关系
- **请求体**:
```json
{
  "parentId": 1,
  "childId": 2,
  "relationType": "string",
  "parentType": "string"
}
```
- **响应**: `201 Created` 返回亲子关系对象

### 2.2 删除亲子关系
- **接口**: `DELETE /api/relationship/deleteParentChild`
- **描述**: 删除父母与子女之间的关系
- **请求参数**:
  - `parentId` (query, required): 父母ID
  - `childId` (query, required): 子女ID
- **响应**: `204 No Content`

### 2.3 获取亲子关系
- **接口**: `GET /api/relationship/getParentChild`
- **描述**: 获取指定的亲子关系
- **请求参数**:
  - `parentId` (query, required): 父母ID
  - `childId` (query, required): 子女ID
- **响应**: `200 OK` 返回亲子关系对象

### 2.4 获取父母的子女
- **接口**: `GET /api/relationship/getChildrenByParentId`
- **描述**: 根据父母ID获取其所有子女
- **请求参数**:
  - `parentId` (query, required): 父母ID
- **响应**: `200 OK` 返回亲子关系数组

### 2.5 获取子女的父母
- **接口**: `GET /api/relationship/getParentsByChildId`
- **描述**: 根据子女ID获取其所有父母
- **请求参数**:
  - `childId` (query, required): 子女ID
- **响应**: `200 OK` 返回亲子关系数组

---

## 三、家庭关系管理接口

### 3.1 创建家庭关系
- **接口**: `POST /api/familyRelation/create`
- **描述**: 创建夫妻/伴侣关系
- **请求体**:
```json
{
  "personId": 1,
  "spouseId": 2,
  "familyId": "string",
  "relationType": "string",
  "marriageDate": "string",
  "marriagePlace": "string",
  "divorceDate": "string",
  "isActive": true
}
```
- **响应**: `201 Created` 返回家庭关系对象

### 3.2 更新家庭关系
- **接口**: `PUT /api/familyRelation/update`
- **描述**: 更新夫妻/伴侣关系信息
- **请求体**: 同创建家庭关系
- **响应**: `200 OK` 返回更新后的家庭关系对象

### 3.3 删除家庭关系
- **接口**: `DELETE /api/familyRelation/delete`
- **描述**: 删除夫妻/伴侣关系
- **请求参数**:
  - `id` (query, required): 家庭关系ID
- **响应**: `204 No Content`

### 3.4 根据ID获取家庭关系
- **接口**: `GET /api/familyRelation/getById`
- **描述**: 根据ID获取夫妻/伴侣关系信息
- **请求参数**:
  - `id` (query, required): 家庭关系ID
- **响应**: `200 OK` 返回家庭关系对象

### 3.5 根据人员ID获取家庭关系
- **接口**: `GET /api/familyRelation/getByPersonId`
- **描述**: 根据人员ID获取其所有夫妻/伴侣关系
- **请求参数**:
  - `personId` (query, required): 人员ID
- **响应**: `200 OK` 返回家庭关系数组

### 3.6 根据配偶ID获取家庭关系
- **接口**: `GET /api/familyRelation/getBySpouseId`
- **描述**: 根据配偶ID获取其所有夫妻/伴侣关系
- **请求参数**:
  - `spouseId` (query, required): 配偶ID
- **响应**: `200 OK` 返回家庭关系数组

### 3.7 根据家庭ID获取家庭关系
- **接口**: `GET /api/familyRelation/getByFamilyId`
- **描述**: 根据家庭ID获取其所有夫妻/伴侣关系
- **请求参数**:
  - `familyId` (query, required): 家庭ID
- **响应**: `200 OK` 返回家庭关系数组

---

## 四、事件管理接口

### 4.1 创建事件
- **接口**: `POST /api/event/create`
- **描述**: 创建新的家族事件
- **请求体**:
```json
{
  "personId": 1,
  "eventType": "string",
  "eventName": "string",
  "eventDate": "string",
  "eventLunar": "string",
  "eventPlace": "string",
  "eventPlaceLongitude": "string",
  "eventPlaceLatitude": "string",
  "description": "string",
  "isImportant": true
}
```
- **响应**: `201 Created` 返回事件对象

### 4.2 更新事件
- **接口**: `PUT /api/event/update`
- **描述**: 更新家族事件信息
- **请求体**: 同创建事件
- **响应**: `200 OK` 返回更新后的事件对象

### 4.3 删除事件
- **接口**: `DELETE /api/event/delete`
- **描述**: 删除家族事件
- **请求参数**:
  - `id` (query, required): 事件ID
- **响应**: `204 No Content`

### 4.4 根据ID获取事件
- **接口**: `GET /api/event/getById`
- **描述**: 根据ID获取家族事件信息
- **请求参数**:
  - `id` (query, required): 事件ID
- **响应**: `200 OK` 返回事件对象

### 4.5 根据人员ID获取事件
- **接口**: `GET /api/event/getByPersonId`
- **描述**: 根据人员ID获取其所有事件
- **请求参数**:
  - `personId` (query, required): 人员ID
  - `page` (query, optional, default=1): 页码
  - `pageSize` (query, optional, default=10): 每页数量
- **响应**:
```json
{
  "data": [],
  "total": 100,
  "page": 1,
  "size": 10
}
```

### 4.6 根据事件类型获取事件
- **接口**: `GET /api/event/getByType`
- **描述**: 根据事件类型获取所有事件
- **请求参数**:
  - `type` (query, required): 事件类型
  - `page` (query, optional, default=1): 页码
  - `pageSize` (query, optional, default=10): 每页数量
- **响应**: 同根据人员ID获取事件接口

### 4.7 获取重要事件
- **接口**: `GET /api/event/getImportant`
- **描述**: 获取所有重要的家族事件
- **请求参数**:
  - `page` (query, optional, default=1): 页码
  - `pageSize` (query, optional, default=10): 每页数量
- **响应**: 同根据人员ID获取事件接口

---

## 五、多媒体管理接口

### 5.1 创建多媒体
- **接口**: `POST /api/media/create`
- **描述**: 创建新的家族多媒体文件
- **请求体**:
```json
{
  "personId": 1,
  "eventId": 1,
  "mediaType": "string",
  "fileName": "string",
  "fileUrl": "string",
  "fileSize": "string",
  "mimeType": "string",
  "thumbnailUrl": "string",
  "description": "string",
  "isPublic": true
}
```
- **响应**: `201 Created` 返回多媒体对象

### 5.2 更新多媒体
- **接口**: `PUT /api/media/update`
- **描述**: 更新家族多媒体文件信息
- **请求体**: 同创建多媒体
- **响应**: `200 OK` 返回更新后的多媒体对象

### 5.3 删除多媒体
- **接口**: `DELETE /api/media/delete`
- **描述**: 删除家族多媒体文件
- **请求参数**:
  - `id` (query, required): 多媒体ID
- **响应**: `204 No Content`

### 5.4 根据ID获取多媒体
- **接口**: `GET /api/media/getById`
- **描述**: 根据ID获取家族多媒体文件信息
- **请求参数**:
  - `id` (query, required): 多媒体ID
- **响应**: `200 OK` 返回多媒体对象

### 5.5 根据人员ID获取多媒体
- **接口**: `GET /api/media/getByPersonId`
- **描述**: 根据人员ID获取其所有多媒体文件
- **请求参数**:
  - `personId` (query, required): 人员ID
  - `page` (query, optional, default=1): 页码
  - `pageSize` (query, optional, default=10): 每页数量
- **响应**:
```json
{
  "data": [],
  "total": 100,
  "page": 1,
  "size": 10
}
```

### 5.6 根据事件ID获取多媒体
- **接口**: `GET /api/media/getByEventId`
- **描述**: 根据事件ID获取其所有多媒体文件
- **请求参数**:
  - `eventId` (query, required): 事件ID
  - `page` (query, optional, default=1): 页码
  - `pageSize` (query, optional, default=10): 每页数量
- **响应**: 同根据人员ID获取多媒体接口

### 5.7 根据媒体类型获取多媒体
- **接口**: `GET /api/media/getByType`
- **描述**: 根据多媒体类型获取所有多媒体文件
- **请求参数**:
  - `type` (query, required): 多媒体类型
  - `page` (query, optional, default=1): 页码
  - `pageSize` (query, optional, default=10): 每页数量
- **响应**: 同根据人员ID获取多媒体接口

### 5.8 获取公开的多媒体
- **接口**: `GET /api/media/getPublic`
- **描述**: 获取所有公开的家族多媒体文件
- **请求参数**:
  - `page` (query, optional, default=1): 页码
  - `pageSize` (query, optional, default=10): 每页数量
- **响应**: 同根据人员ID获取多媒体接口

---

## 接口统计

| 模块 | 接口数量 | 主要功能 |
|------|---------|---------|
| 人员管理 | 7 | 人员的CRUD、搜索、列表 |
| 亲子关系管理 | 5 | 亲子关系的CRUD、查询 |
| 家庭关系管理 | 7 | 家庭关系的CRUD、多维度查询 |
| 事件管理 | 7 | 事件的CRUD、多维度查询 |
| 多媒体管理 | 9 | 多媒体的CRUD、多维度查询 |
| **总计** | **35** | - |

---

## 通用响应格式

### 成功响应
```json
{
  "id": 1,
  "name": "example",
  // ... 其他字段
}
```

### 分页响应
```json
{
  "data": [],
  "total": 100,
  "page": 1,
  "size": 10
}
```

### 错误响应
```json
{
  "error": "错误信息描述"
}
```

### HTTP 状态码说明

- `200 OK`: 请求成功
- `201 Created`: 资源创建成功
- `204 No Content`: 请求成功但无返回内容
- `400 Bad Request`: 请求参数错误
- `500 Internal Server Error`: 服务器内部错误

import axios from 'axios'
import { 
  Person, 
  PersonListResponse, 
  ParentChildRelationship, 
  FamilyRelation, 
  Event, 
  EventListResponse, 
  Media, 
  MediaListResponse, 
  SearchParams 
} from '@/types/api'

// 创建axios实例
const apiClient = axios.create({
  baseURL: 'http://localhost:8080/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// 请求拦截器
apiClient.interceptors.request.use(
  (config) => {
    // 可以在这里添加认证token等
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
apiClient.interceptors.response.use(
  (response) => {
    // 返回完整的响应对象，让调用方处理数据
    return response
  },
  (error) => {
    console.error('API Error:', error)
    return Promise.reject(error)
  }
)

// 人员管理API
export const personApi = {
  // 创建人员
  create: (person: Person) => apiClient.post<Person>('/person/create', person),
  
  // 更新人员
  update: (person: Person) => apiClient.put<Person>('/person/update', person),
  
  // 删除人员
  delete: (id: number) => apiClient.delete(`/person/delete?id=${id}`),
  
  // 根据ID获取人员
  getById: (id: number) => apiClient.get<Person>(`/person/getById?id=${id}`),
  
  // 根据UUID获取人员
  getByUuid: (uuid: string) => apiClient.get<Person>(`/person/getByUuid?uuid=${uuid}`),
  
  // 列出人员
  list: (params?: SearchParams) => apiClient.get<PersonListResponse>('/person/list', { params }),
  
  // 搜索人员
  search: (params: SearchParams) => apiClient.get<PersonListResponse>('/person/search', { params }),
}

// 亲子关系管理API
export const relationshipApi = {
  // 创建亲子关系
  createParentChild: (relation: ParentChildRelationship) => 
    apiClient.post<ParentChildRelationship>('/relationship/createParentChild', relation),
  
  // 删除亲子关系
  deleteParentChild: (parentId: number, childId: number) => 
    apiClient.delete(`/relationship/deleteParentChild?parentId=${parentId}&childId=${childId}`),
  
  // 获取亲子关系
  getParentChild: (parentId: number, childId: number) => 
    apiClient.get<ParentChildRelationship>(`/relationship/getParentChild?parentId=${parentId}&childId=${childId}`),
  
  // 获取父母的子女
  getChildrenByParentId: (parentId: number) => 
    apiClient.get<ParentChildRelationship[]>(`/relationship/getChildrenByParentId?parentId=${parentId}`),
  
  // 获取子女的父母
  getParentsByChildId: (childId: number) => 
    apiClient.get<ParentChildRelationship[]>(`/relationship/getParentsByChildId?childId=${childId}`),
}

// 家庭关系管理API
export const familyRelationApi = {
  // 创建家庭关系
  create: (relation: FamilyRelation) => 
    apiClient.post<FamilyRelation>('/familyRelation/create', relation),
  
  // 更新家庭关系
  update: (relation: FamilyRelation) => 
    apiClient.put<FamilyRelation>('/familyRelation/update', relation),
  
  // 删除家庭关系
  delete: (id: number) => apiClient.delete(`/familyRelation/delete?id=${id}`),
  
  // 根据ID获取家庭关系
  getById: (id: number) => apiClient.get<FamilyRelation>(`/familyRelation/getById?id=${id}`),
  
  // 根据人员ID获取家庭关系
  getByPersonId: (personId: number) => 
    apiClient.get<FamilyRelation[]>(`/familyRelation/getByPersonId?personId=${personId}`),
  
  // 根据配偶ID获取家庭关系
  getBySpouseId: (spouseId: number) => 
    apiClient.get<FamilyRelation[]>(`/familyRelation/getBySpouseId?spouseId=${spouseId}`),
  
  // 根据家庭ID获取家庭关系
  getByFamilyId: (familyId: string) => 
    apiClient.get<FamilyRelation[]>(`/familyRelation/getByFamilyId?familyId=${familyId}`),
}

// 事件管理API
export const eventApi = {
  // 创建事件
  create: (event: Event) => apiClient.post<Event>('/event/create', event),
  
  // 更新事件
  update: (event: Event) => apiClient.put<Event>('/event/update', event),
  
  // 删除事件
  delete: (id: number) => apiClient.delete(`/event/delete?id=${id}`),
  
  // 根据ID获取事件
  getById: (id: number) => apiClient.get<Event>(`/event/getById?id=${id}`),
  
  // 根据人员ID获取事件
  getByPersonId: (personId: number, params?: SearchParams) => 
    apiClient.get<EventListResponse>(`/event/getByPersonId?personId=${personId}`, { params }),
  
  // 根据事件类型获取事件
  getByType: (type: string, params?: SearchParams) => 
    apiClient.get<EventListResponse>(`/event/getByType?type=${type}`, { params }),
  
  // 获取重要事件
  getImportant: (params?: SearchParams) => 
    apiClient.get<EventListResponse>('/event/getImportant', { params }),
}

// 媒体管理API
export const mediaApi = {
  // 创建媒体
  create: (media: Media) => apiClient.post<Media>('/media/create', media),
  
  // 更新媒体
  update: (media: Media) => apiClient.put<Media>('/media/update', media),
  
  // 删除媒体
  delete: (id: number) => apiClient.delete(`/media/delete?id=${id}`),
  
  // 根据ID获取媒体
  getById: (id: number) => apiClient.get<Media>(`/media/getById?id=${id}`),
  
  // 根据人员ID获取媒体
  getByPersonId: (personId: number, params?: SearchParams) => 
    apiClient.get<MediaListResponse>(`/media/getByPersonId?personId=${personId}`, { params }),
}
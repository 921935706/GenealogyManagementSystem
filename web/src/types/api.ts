// 基础响应类型
export interface ApiResponse<T> {
  data: T
  total?: number
  page?: number
  size?: number
}

// 分页参数
export interface PaginationParams {
  page?: number
  pageSize?: number
}

// 人员管理相关类型
export interface Person {
  id?: number
  uuid?: string
  name: string
  firstName?: string
  lastName?: string
  generationName?: string
  alias?: string
  englishName?: string
  gender: 'male' | 'female'
  isAlive: boolean
  birthDate?: string
  birthLunar?: string
  birthPlace?: string
  birthPlaceLongitude?: string
  birthPlaceLatitude?: string
  deathDate?: string
  deathLunar?: string
  deathPlace?: string
  deathPlaceLongitude?: string
  deathPlaceLatitude?: string
  deathCause?: string
  generation?: string
  occupation?: string
  education?: string
  biography?: string
  avatarUrl?: string
  coverPhotoUrl?: string
  isPublic: boolean
  verificationStatus?: string
  createdBy?: string
  createdAt?: string
  updatedAt?: string
  email?: string
}

export interface PersonListResponse extends ApiResponse<Person[]> {}

// 亲子关系相关类型
export interface ParentChildRelationship {
  id?: number
  parentId: number
  childId: number
  relationType: string
  parentType: string
  createdAt?: string
  updatedAt?: string
}

// 家庭关系相关类型
export interface FamilyRelation {
  id?: number
  personId: number
  spouseId: number
  familyId?: string
  relationType: string
  marriageDate?: string
  marriagePlace?: string
  divorceDate?: string
  isActive: boolean
  createdAt?: string
  updatedAt?: string
}

// 事件管理相关类型
export interface Event {
  id?: number
  personId: number
  eventType: string
  eventName: string
  eventDate?: string
  eventLunar?: string
  eventPlace?: string
  eventPlaceLongitude?: string
  eventPlaceLatitude?: string
  description?: string
  isImportant: boolean
  createdAt?: string
  updatedAt?: string
}

export interface EventListResponse extends ApiResponse<Event[]> {}

// 媒体管理相关类型
export interface Media {
  id?: number
  personId: number
  eventId?: number
  mediaType: string
  fileName: string
  fileUrl: string
  fileSize?: string
  mimeType?: string
  thumbnailUrl?: string
  description?: string
  isPublic: boolean
  createdAt?: string
  updatedAt?: string
}

export interface MediaListResponse extends ApiResponse<Media[]> {}

// 搜索参数
export interface SearchParams extends PaginationParams {
  name?: string
  type?: string
  personId?: number
}
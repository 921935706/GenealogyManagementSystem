import { create } from 'zustand'
import { ParentChildRelationship } from '@/types/api'
import { relationshipApi } from '@/lib/api'
import { notifications } from '@mantine/notifications'

interface RelationshipState {
  relationships: ParentChildRelationship[]
  loading: boolean
  error: string | null
  
  // 操作
  createParentChild: (relation: ParentChildRelationship) => Promise<boolean>
  deleteParentChild: (parentId: number, childId: number) => Promise<boolean>
  getChildrenByParentId: (parentId: number) => Promise<void>
  getParentsByChildId: (childId: number) => Promise<void>
  clearError: () => void
}

export const useRelationshipStore = create<RelationshipState>((set) => ({
  relationships: [],
  loading: false,
  error: null,

  createParentChild: async (relation: ParentChildRelationship) => {
    set({ loading: true, error: null })
    try {
      await relationshipApi.createParentChild(relation)
      set({ loading: false })
      notifications.show({
        title: '成功',
        message: '亲子关系创建成功',
        color: 'green',
      })
      return true
    } catch (error: any) {
      set({ 
        error: error.message || '创建亲子关系失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '创建亲子关系失败',
        color: 'red',
      })
      return false
    }
  },

  deleteParentChild: async (parentId: number, childId: number) => {
    set({ loading: true, error: null })
    try {
      await relationshipApi.deleteParentChild(parentId, childId)
      set({ loading: false })
      notifications.show({
        title: '成功',
        message: '亲子关系删除成功',
        color: 'green',
      })
      return true
    } catch (error: any) {
      set({ 
        error: error.message || '删除亲子关系失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '删除亲子关系失败',
        color: 'red',
      })
      return false
    }
  },

  getChildrenByParentId: async (parentId: number) => {
    set({ loading: true, error: null })
    try {
      const response = await relationshipApi.getChildrenByParentId(parentId)
      set({ 
        relationships: response?.data || [],
        loading: false 
      })
    } catch (error: any) {
      set({ 
        error: error.message || '获取子女列表失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '获取子女列表失败',
        color: 'red',
      })
    }
  },

  getParentsByChildId: async (childId: number) => {
    set({ loading: true, error: null })
    try {
      const response = await relationshipApi.getParentsByChildId(childId)
      set({ 
        relationships: response?.data || [],
        loading: false 
      })
    } catch (error: any) {
      set({ 
        error: error.message || '获取父母列表失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '获取父母列表失败',
        color: 'red',
      })
    }
  },

  clearError: () => {
    set({ error: null })
  },
}))
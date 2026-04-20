import { create } from 'zustand'
import { Person } from '@/types/api'
import { personApi } from '@/lib/api'
import { notifications } from '@mantine/notifications'

interface PersonState {
  persons: Person[]
  currentPerson: Person | null
  loading: boolean
  error: string | null
  total: number
  
  // 操作
  fetchPersons: (page?: number, pageSize?: number) => Promise<void>
  fetchPersonById: (id: number) => Promise<void>
  createPerson: (person: Person) => Promise<boolean>
  updatePerson: (person: Person) => Promise<boolean>
  deletePerson: (id: number) => Promise<boolean>
  searchPersons: (name: string, page?: number, pageSize?: number) => Promise<void>
  setCurrentPerson: (person: Person | null) => void
  clearError: () => void
}

export const usePersonStore = create<PersonState>((set) => ({
  persons: [],
  currentPerson: null,
  loading: false,
  error: null,
  total: 0,

  fetchPersons: async (page = 1, pageSize = 10) => {
    set({ loading: true, error: null })
    try {
      const response = await personApi.list({ page, pageSize })
      set({ 
        persons: response?.data?.data || [],
        total: response?.data?.total || 0,
        loading: false 
      })
    } catch (error: any) {
      set({ 
        error: error.message || '获取人员列表失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '获取人员列表失败',
        color: 'red',
      })
    }
  },

  fetchPersonById: async (id: number) => {
    set({ loading: true, error: null })
    try {
      const response = await personApi.getById(id)
      set({ currentPerson: response?.data, loading: false })
    } catch (error: any) {
      set({ 
        error: error.message || '获取人员信息失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '获取人员信息失败',
        color: 'red',
      })
    }
  },

  createPerson: async (person: Person) => {
    set({ loading: true, error: null })
    try {
      await personApi.create(person)
      set({ loading: false })
      notifications.show({
        title: '成功',
        message: '人员创建成功',
        color: 'green',
      })
      return true
    } catch (error: any) {
      set({ 
        error: error.message || '创建人员失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '创建人员失败',
        color: 'red',
      })
      return false
    }
  },

  updatePerson: async (person: Person) => {
    set({ loading: true, error: null })
    try {
      await personApi.update(person)
      set({ loading: false })
      notifications.show({
        title: '成功',
        message: '人员信息更新成功',
        color: 'green',
      })
      return true
    } catch (error: any) {
      set({ 
        error: error.message || '更新人员信息失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '更新人员信息失败',
        color: 'red',
      })
      return false
    }
  },

  deletePerson: async (id: number) => {
    set({ loading: true, error: null })
    try {
      await personApi.delete(id)
      set({ loading: false })
      notifications.show({
        title: '成功',
        message: '人员删除成功',
        color: 'green',
      })
      return true
    } catch (error: any) {
      set({ 
        error: error.message || '删除人员失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '删除人员失败',
        color: 'red',
      })
      return false
    }
  },

  searchPersons: async (name: string, page = 1, pageSize = 10) => {
    set({ loading: true, error: null })
    try {
      const response = await personApi.search({ name, page, pageSize })
      set({ 
        persons: response?.data?.data || [],
        total: response?.data?.total || 0,
        loading: false 
      })
    } catch (error: any) {
      set({ 
        error: error.message || '搜索人员失败',
        loading: false 
      })
      notifications.show({
        title: '错误',
        message: '搜索人员失败',
        color: 'red',
      })
    }
  },

  setCurrentPerson: (person: Person | null) => {
    set({ currentPerson: person })
  },

  clearError: () => {
    set({ error: null })
  },
}))
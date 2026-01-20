import type { LoginParams, UserInfo } from '@/types/user'
import { message } from 'ant-design-vue'
import { defineStore } from 'pinia'
import { getUserInfo, login, register } from '@/apis/modules/user'
import { clearToken, getToken, setToken as setTokenUtil } from '@/utils'

interface UserState {
  token: string
  userInfo: UserInfo | null
  roles: string[]
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    token: getToken() || '',
    userInfo: null,
    roles: [],
  }),

  getters: {
    isLoggedIn: state => !!state.token,
    getUserRoles: state => state.roles,
    getHomePath: state => state.userInfo?.homePath || '/dashboard',
  },

  actions: {
    setToken(token: string) {
      this.token = token
      setTokenUtil(token)
    },

    setUserInfo(userInfo: UserInfo) {
      this.userInfo = userInfo
      this.roles = userInfo.roles
    },

    async login(params: LoginParams) {
      try {
        // 调用实际的登录 API
        const response = await login({
          username: params.username,
          password: params.password,
        })

        if (response.data.code === 200) {
          const token = response.data.data.token
          this.setToken(token)

          // 获取用户信息
          const userInfoResponse = await getUserInfo()
          if (userInfoResponse.data.code === 200) {
            const userInfoData = userInfoResponse.data.data
            const userInfo: UserInfo = {
              id: String(userInfoData.id),
              username: userInfoData.username,
              nickname: userInfoData.username,
              avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${userInfoData.username}`,
              email: `${userInfoData.username}@example.com`,
              roles: [userInfoData.role],
              homePath: '/dashboard',
            }
            this.setUserInfo(userInfo)
          }

          message.success('登录成功')
          return response.data
        }
        else {
          message.error(response.data.msg || '登录失败')
          throw new Error(response.data.msg)
        }
      }
      catch (error: any) {
        message.error(error.response?.data?.msg || '登录失败')
        throw error
      }
    },

    async register(params: LoginParams) {
      try {
        const response = await register({
          username: params.username,
          password: params.password,
        })

        if (response.data.code === 200) {
          message.success('注册成功')
          return response.data
        }
        else {
          message.error(response.data.msg || '注册失败')
          throw new Error(response.data.msg)
        }
      }
      catch (error: any) {
        message.error(error.response?.data?.msg || '注册失败')
        throw error
      }
    },

    async getUserInfo() {
      try {
        const response = await getUserInfo()
        if (response.data.code === 200) {
          const userInfoData = response.data.data
          const userInfo: UserInfo = {
            id: String(userInfoData.id),
            username: userInfoData.username,
            nickname: userInfoData.username,
            avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${userInfoData.username}`,
            email: `${userInfoData.username}@example.com`,
            roles: [userInfoData.role],
            homePath: '/dashboard',
          }
          this.setUserInfo(userInfo)
          return userInfo
        }
      }
      catch (error) {
        console.error('获取用户信息失败:', error)
        throw error
      }
    },

    logout() {
      this.token = ''
      this.userInfo = null
      this.roles = []
      clearToken()
      message.success('已退出登录')
    },
  },

  persist: true,
})

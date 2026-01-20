import baseAxios from '..'

// 认证相关 API
export function login(data: { username: string; password: string }) {
  return baseAxios.post<DataResponse<{ token: string }>>('/api/v1/auth/login', data)
}

export function register(data: { username: string; password: string }) {
  return baseAxios.post<DataResponse<{ id: number; username: string }>>('/api/v1/auth/register', data)
}

export function checkUser(data: { username: string }) {
  return baseAxios.post<DataResponse<{ exists: boolean }>>('/api/v1/auth/check-user', data)
}

export function getUserInfo() {
  return baseAxios.get<DataResponse<{ id: number; username: string; role: string }>>('/api/v1/auth/info')
}

// 用户管理相关 API
export function getAllUsers(params: { pageNum: number; pageSize: number }) {
  return baseAxios.get<UserResponse>('/api/v1/hello/user/all', { params })
}

export function addUser(data: AddUserParams) {
  return baseAxios.post<DataResponse<User>>('/api/v1/hello/user/add', data)
}

export function deleteUsers(ids: number[]) {
  return baseAxios.delete<DataResponse<{ deletedCount: number }>>('/api/v1/hello/user/delete', {
    data: { ids },
  })
}

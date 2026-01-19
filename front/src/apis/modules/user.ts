import baseAxios from '..'

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

interface User {
  id: number
  name: string
  age: number
  email: string
  createTime: string
  role: string
  project: string
}

type UserResponse = TableResponse<User>

type AddUserParams = Pick<User, 'name' | 'role' | 'project'> &
  Partial<Pick<User, 'age' | 'email'>>

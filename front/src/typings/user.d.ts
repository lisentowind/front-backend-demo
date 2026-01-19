interface User {
  id: number
  name: string
  age: number
  email: string
  createTime: string
}

type UserResponse = TableResponse<User>

type AddUserParams = Pick<User, 'name'> & Partial<Pick<User, 'age', 'email'>>

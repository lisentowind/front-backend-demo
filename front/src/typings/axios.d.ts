interface DataResponse<T = unknown> {
  code: number
  data: T
  msg: string
}

interface TableResponse<T = unknown> {
  code: number
  data: {
    total: number
    list: T[]
    page: number
    size: number
  }
  msg: string
}

interface HttpResponse<T = DataResponse> {
  status: number
  msg: string
  code: number
  data: T
  result: T
  statusText: string
  headers: RawAxiosResponseHeaders | AxiosResponseHeaders
  config: AxiosRequestConfig
}

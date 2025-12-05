import axios from 'axios'

const API_BASE = 'http://localhost:3061'

// 创建axios实例
const request = axios.create({
  baseURL: API_BASE,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 查询日程（day可选，不传则查询整月）
export async function querySchedule(params) {
  const { user_id, year, month, day } = params
  const requestData = {
    user_id,
    year,
    month
  }
  // day参数可选，不传则查询整月
  if (day !== undefined && day !== null) {
    requestData.day = day
  }
  const response = await request.post('/schedule/query', requestData)
  return response.data
}

// 新增日程
export async function storeSchedule(params) {
  const { year, month, day, user_id, content, start, end, priority, status } = params
  const response = await request.post('/schedule/store', {
    year,
    month,
    day,
    user_id,
    content,
    start,
    end,
    priority,
    status
  })
  return response.data
}

// 更新日程
export async function updateSchedule(params) {
  const { id, year, month, day, start, end, content, status, user_id, priority } = params
  const response = await request.post('/schedule/update', {
    id,
    year,
    month,
    day,
    start,
    end,
    content,
    status,
    user_id,
    priority
  })
  return response.data
}

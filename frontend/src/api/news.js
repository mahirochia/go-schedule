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

// 查询某一天的新闻
export async function queryNews(date) {
  const response = await request.post('/news/query', {
    date  // 格式: "2025-12-04 00:00:00"
  })
  return response.data
}


import { defineStore } from 'pinia'
import { ref } from 'vue'
import { queryNews } from '@/api/news'

export const useNewsStore = defineStore('news', () => {
  // 当天的新闻列表
  const newsList = ref([])
  
  // 加载状态
  const loading = ref(false)

  // 格式化数字为两位
  function padZero(num) {
    return String(num).padStart(2, '0')
  }

  // 生成日期字符串 "YYYY-MM-DD 00:00:00"
  function formatDateString(year, month, day) {
    return `${year}-${padZero(month)}-${padZero(day)} 00:00:00`
  }
  
  // 查询新闻
  async function fetchNews(year, month, day) {
    loading.value = true
    try {
      const dateStr = formatDateString(year, month, day)
      const res = await queryNews(dateStr)
      if (res.code === 200) {
        newsList.value = res.data || []
      } else {
        console.error('查询新闻失败:', res.msg)
        newsList.value = []
      }
    } catch (error) {
      console.error('查询新闻出错:', error)
      newsList.value = []
    } finally {
      loading.value = false
    }
  }

  // 清空新闻列表
  function clearNews() {
    newsList.value = []
  }
  
  return {
    newsList,
    loading,
    fetchNews,
    clearNews
  }
})


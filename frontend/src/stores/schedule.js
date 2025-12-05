import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { querySchedule, storeSchedule, updateSchedule } from '@/api/schedule'

export const useScheduleStore = defineStore('schedule', () => {
  // 当前用户ID（模拟）
  const userId = ref(1)
  
  // 当前选中的日期
  const selectedDate = ref({
    year: new Date().getFullYear(),
    month: new Date().getMonth() + 1,
    day: new Date().getDate()
  })
  
  // 当前日期的日程列表
  const scheduleList = ref([])
  
  // 月度日程缓存 { "2025-12-01": [...schedules], "2025-12-02": [...] }
  const monthlySchedules = ref({})
  
  // 当前加载的月份
  const loadedMonth = ref({ year: 0, month: 0 })
  
  // 加载状态
  const loading = ref(false)
  const monthLoading = ref(false)
  
  // 格式化显示的日期
  const formattedDate = computed(() => {
    const { year, month, day } = selectedDate.value
    return `${year}年${month}月${day}日`
  })

  // 格式化数字为两位
  function padZero(num) {
    return String(num).padStart(2, '0')
  }

  // 生成日期key
  function getDateKey(year, month, day) {
    return `${year}-${padZero(month)}-${padZero(day)}`
  }

  // 将各种时间格式统一转换为本地时区的 "YYYY-MM-DD HH:mm:ss" 格式
  function normalizeDateTime(timeStr, year, month, day) {
    if (!timeStr) return ''
    
    let date
    
    // ISO 8601 格式: "2025-12-02T18:09:00+08:00"
    if (timeStr.includes('T')) {
      date = new Date(timeStr)
    }
    // 已经是 "YYYY-MM-DD HH:mm:ss" 格式
    else if (timeStr.includes(' ') && timeStr.includes('-')) {
      return timeStr
    }
    // 纯时间格式: "11:00:00" 或 "11:00"
    else if (timeStr.includes(':')) {
      const parts = timeStr.split(':')
      const hours = parseInt(parts[0]) || 0
      const minutes = parseInt(parts[1]) || 0
      return `${year}-${padZero(month)}-${padZero(day)} ${padZero(hours)}:${padZero(minutes)}:00`
    }
    else {
      return timeStr
    }
    
    // 从Date对象格式化为本地时间字符串
    if (date && !isNaN(date.getTime())) {
      const y = date.getFullYear()
      const m = padZero(date.getMonth() + 1)
      const d = padZero(date.getDate())
      const h = padZero(date.getHours())
      const min = padZero(date.getMinutes())
      return `${y}-${m}-${d} ${h}:${min}:00`
    }
    
    return timeStr
  }
  
  // 设置选中的日期
  function setSelectedDate(year, month, day) {
    selectedDate.value = { year, month, day }
  }
  
  // 查询某天的日程
  async function fetchSchedules() {
    loading.value = true
    try {
      const { year, month, day } = selectedDate.value
      const res = await querySchedule({
        user_id: userId.value,
        year,
        month,
        day
      })
      if (res.code === 200) {
        scheduleList.value = res.data || []
        // 同时更新月度缓存
        const dateKey = getDateKey(year, month, day)
        monthlySchedules.value[dateKey] = res.data || []
      } else {
        console.error('查询日程失败:', res.msg)
        scheduleList.value = []
      }
    } catch (error) {
      console.error('查询日程出错:', error)
      scheduleList.value = []
    } finally {
      loading.value = false
    }
  }

  // 查询整个月的日程（不传day参数）
  async function fetchMonthSchedules(year, month) {
    // 如果已经加载过这个月，直接返回
    if (loadedMonth.value.year === year && loadedMonth.value.month === month) {
      return
    }
    
    monthLoading.value = true
    
    try {
      // 不传day参数，查询整月数据
      const res = await querySchedule({
        user_id: userId.value,
        year,
        month
      })
      
      if (res.code === 200) {
        const allSchedules = res.data || []
        
        // 清空当前月的缓存
        Object.keys(monthlySchedules.value).forEach(key => {
          if (key.startsWith(`${year}-${padZero(month)}`)) {
            delete monthlySchedules.value[key]
          }
        })
        
        // 按日期分组存储
        allSchedules.forEach(schedule => {
          const dateKey = getDateKey(schedule.year, schedule.month, schedule.day)
          if (!monthlySchedules.value[dateKey]) {
            monthlySchedules.value[dateKey] = []
          }
          monthlySchedules.value[dateKey].push(schedule)
        })
        
        loadedMonth.value = { year, month }
      }
    } catch (error) {
      console.error('查询月度日程出错:', error)
    } finally {
      monthLoading.value = false
    }
  }

  // 获取某天的未完成日程（状态不是4-已完成）
  function getUnfinishedSchedules(year, month, day) {
    const dateKey = getDateKey(year, month, day)
    const daySchedules = monthlySchedules.value[dateKey] || []
    // 返回未完成的日程（状态不是4）
    return daySchedules.filter(s => s.status !== 4)
  }

  // 清除月度缓存
  function clearMonthlyCache() {
    monthlySchedules.value = {}
    loadedMonth.value = { year: 0, month: 0 }
  }
  
  // 新增日程
  async function addSchedule(scheduleData) {
    loading.value = true
    try {
      const { year, month, day } = selectedDate.value
      const res = await storeSchedule({
        year,
        month,
        day,
        user_id: userId.value,
        content: scheduleData.content,
        start: scheduleData.start,
        end: scheduleData.end,
        priority: scheduleData.priority,
        status: scheduleData.status || 1
      })
      if (res.code === 200) {
        await fetchSchedules()
        // 刷新月度缓存
        clearMonthlyCache()
        await fetchMonthSchedules(year, month)
        return { success: true, msg: res.msg }
      } else {
        return { success: false, msg: res.msg }
      }
    } catch (error) {
      console.error('新增日程出错:', error)
      return { success: false, msg: '新增日程失败' }
    } finally {
      loading.value = false
    }
  }
  
  // 更新日程
  async function editSchedule(scheduleData) {
    loading.value = true
    try {
      const res = await updateSchedule({
        id: scheduleData.id,
        year: scheduleData.year,
        month: scheduleData.month,
        day: scheduleData.day,
        start: scheduleData.start,
        end: scheduleData.end,
        content: scheduleData.content,
        status: scheduleData.status,
        user_id: userId.value,
        priority: scheduleData.priority
      })
      if (res.code === 200) {
        await fetchSchedules()
        // 刷新月度缓存
        clearMonthlyCache()
        const { year, month } = selectedDate.value
        await fetchMonthSchedules(year, month)
        return { success: true, msg: res.msg }
      } else {
        return { success: false, msg: res.msg }
      }
    } catch (error) {
      console.error('更新日程出错:', error)
      return { success: false, msg: '更新日程失败' }
    } finally {
      loading.value = false
    }
  }
  
  // 切换日程状态
  async function toggleScheduleStatus(schedule, newStatus) {
    const normalizedStart = normalizeDateTime(schedule.start_time, schedule.year, schedule.month, schedule.day)
    const normalizedEnd = normalizeDateTime(schedule.end_time, schedule.year, schedule.month, schedule.day)
    
    return await editSchedule({
      ...schedule,
      status: newStatus,
      start: normalizedStart,
      end: normalizedEnd
    })
  }
  
  return {
    userId,
    selectedDate,
    scheduleList,
    monthlySchedules,
    loading,
    monthLoading,
    formattedDate,
    setSelectedDate,
    fetchSchedules,
    fetchMonthSchedules,
    getUnfinishedSchedules,
    clearMonthlyCache,
    addSchedule,
    editSchedule,
    toggleScheduleStatus
  }
})

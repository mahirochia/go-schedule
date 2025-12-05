<script setup>
import { ref, watch, onMounted, computed } from 'vue'
import { useScheduleStore } from '@/stores/schedule'
import { useNewsStore } from '@/stores/news'
import { ArrowLeft, ArrowRight, Calendar as CalendarIcon, Clock, Edit, Plus, Notebook, Reading, View, ChatDotRound, Star } from '@element-plus/icons-vue'

const scheduleStore = useScheduleStore()
const newsStore = useNewsStore()
const emit = defineEmits(['add', 'edit'])

// 当前模式：schedule（日程） 或 news（新闻）
const currentMode = ref('schedule')

// 日历实例引用
const calendarRef = ref(null)

// 使用Date对象来控制日历
const calendarValue = ref(new Date())

// 当前显示的年月
const currentYear = ref(new Date().getFullYear())
const currentMonth = ref(new Date().getMonth() + 1)

// 详情对话框
const detailDialogVisible = ref(false)
const selectedDayDate = ref({ year: 0, month: 0, day: 0 })

// 状态映射
const statusMap = {
  1: { label: '未开始', type: 'info' },
  2: { label: '进行中', type: 'primary' },
  3: { label: '已结束', type: 'warning' },
  4: { label: '已完成', type: 'success' }
}

const priorityMap = {
  1: { label: '低', type: 'success' },
  2: { label: '中', type: 'warning' },
  3: { label: '高', type: 'danger' }
}

// 初始化时加载当月日程
onMounted(() => {
  scheduleStore.fetchMonthSchedules(currentYear.value, currentMonth.value)
})

// 监听日历值变化，更新当前年月并加载日程
watch(calendarValue, (newDate) => {
  const year = newDate.getFullYear()
  const month = newDate.getMonth() + 1
  
  if (year !== currentYear.value || month !== currentMonth.value) {
    currentYear.value = year
    currentMonth.value = month
    scheduleStore.fetchMonthSchedules(year, month)
  }
})

// 点击日期单元格 - 打开详情对话框
function handleDateClick(data) {
  const date = data.date
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  
  // 更新选中日期
  calendarValue.value = date
  scheduleStore.setSelectedDate(year, month, day)
  selectedDayDate.value = { year, month, day }
  
  // 根据模式加载数据
  if (currentMode.value === 'news') {
    newsStore.fetchNews(year, month, day)
  }
  
  detailDialogVisible.value = true
}

// 获取某天的所有日程
function getDaySchedules(date) {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const dateKey = `${year}-${String(month).padStart(2, '0')}-${String(day).padStart(2, '0')}`
  return scheduleStore.monthlySchedules[dateKey] || []
}

// 获取某天的未完成日程
function getUnfinishedTasks(date) {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  return scheduleStore.getUnfinishedSchedules(year, month, day)
}

// 当前选中日期的日程列表
const selectedDaySchedules = computed(() => {
  const { year, month, day } = selectedDayDate.value
  if (!year) return []
  const date = new Date(year, month - 1, day)
  return getDaySchedules(date)
})

// 判断是否为选中的日期
function isSelected(date) {
  return date.getFullYear() === scheduleStore.selectedDate.year &&
         date.getMonth() + 1 === scheduleStore.selectedDate.month &&
         date.getDate() === scheduleStore.selectedDate.day
}

// 判断是否为今天
function isToday(date) {
  const today = new Date()
  return date.getFullYear() === today.getFullYear() &&
         date.getMonth() === today.getMonth() &&
         date.getDate() === today.getDate()
}

// 格式化时间
function formatTime(time) {
  if (!time) return '--:--'
  if (time.includes('T')) {
    const timePart = time.split('T')[1]
    const timeWithoutTz = timePart.split('+')[0]
    return timeWithoutTz.substring(0, 5)
  }
  if (time.includes(' ')) {
    return time.split(' ')[1].substring(0, 5)
  }
  if (time.includes(':')) {
    return time.substring(0, 5)
  }
  return '--:--'
}

// 格式化发布时间（只显示时分秒）
function formatPublishTime(time) {
  if (!time) return ''
  // ISO 8601 格式: "2025-12-02T18:09:00+08:00"
  if (time.includes('T')) {
    const timePart = time.split('T')[1]
    const timeWithoutTz = timePart.split('+')[0].split('-')[0]  // 去掉时区
    return timeWithoutTz.substring(0, 8)  // 返回 "HH:mm:ss"
  }
  // 传统格式: "2025-12-01 11:00:00"
  if (time.includes(' ')) {
    return time.split(' ')[1].substring(0, 8)  // 返回 "HH:mm:ss"
  }
  // 已经是时间格式
  if (time.includes(':')) {
    return time.substring(0, 8)
  }
  return time
}

// 使用日历组件的selectDate方法进行导航
function selectDate(type) {
  if (calendarRef.value) {
    calendarRef.value.selectDate(type)
  }
}

// 回到今天
function goToToday() {
  selectDate('today')
  const today = new Date()
  scheduleStore.setSelectedDate(today.getFullYear(), today.getMonth() + 1, today.getDate())
}

// 新增日程
function handleAdd() {
  detailDialogVisible.value = false
  emit('add')
}

// 编辑日程
function handleEdit(schedule) {
  detailDialogVisible.value = false
  emit('edit', schedule)
}

// 切换完成状态
async function toggleComplete(schedule) {
  const newStatus = schedule.status === 4 ? 1 : 4
  await scheduleStore.toggleScheduleStatus(schedule, newStatus)
}

// 打开新闻链接
function openNews(news) {
  if (news.link) {
    window.open(news.link, '_blank')
  }
}

// 格式化对话框标题日期
const dialogTitle = computed(() => {
  const { year, month, day } = selectedDayDate.value
  const modeText = currentMode.value === 'schedule' ? '日程' : '新闻'
  return `${year}年${month}月${day}日 ${modeText}`
})

// 对话框内切换模式时加载数据
function handleModeChange(mode) {
  if (mode === 'news' && newsStore.newsList.length === 0) {
    const { year, month, day } = selectedDayDate.value
    newsStore.fetchNews(year, month, day)
  }
}
</script>

<template>
  <el-card class="calendar-card" shadow="never">
    <!-- 模式切换 -->
    <div class="mode-switch">
      <el-segmented v-model="currentMode" :options="[
        { label: '日程', value: 'schedule', icon: Notebook },
        { label: '新闻', value: 'news', icon: Reading }
      ]" />
    </div>
    
    <el-calendar ref="calendarRef" v-model="calendarValue" class="custom-calendar">
      <!-- 自定义头部 -->
      <template #header="{ date }">
        <div class="calendar-header">
          <div class="header-nav">
            <el-button-group>
              <el-button @click="selectDate('prev-year')" size="small">
                <el-icon><ArrowLeft /></el-icon>
                <el-icon><ArrowLeft /></el-icon>
              </el-button>
              <el-button @click="selectDate('prev-month')" size="small">
                <el-icon><ArrowLeft /></el-icon>
              </el-button>
            </el-button-group>
            
            <span class="header-title">{{ date }}</span>
            
            <el-button-group>
              <el-button @click="selectDate('next-month')" size="small">
                <el-icon><ArrowRight /></el-icon>
              </el-button>
              <el-button @click="selectDate('next-year')" size="small">
                <el-icon><ArrowRight /></el-icon>
                <el-icon><ArrowRight /></el-icon>
              </el-button>
            </el-button-group>
          </div>
          
          <el-button type="primary" @click="goToToday" style="width: 100%;">
            <el-icon><CalendarIcon /></el-icon>
            今天
          </el-button>
        </div>
      </template>
      
      <!-- 自定义日期单元格 -->
      <template #date-cell="{ data }">
        <div 
          class="date-cell"
          :class="{ 
            'is-selected': isSelected(data.date),
            'is-today': isToday(data.date),
            'is-other-month': data.type !== 'current-month'
          }"
          @click.stop="handleDateClick(data)"
        >
          <div class="date-header">
            <span class="date-number">{{ data.day.split('-')[2] }}</span>
            <span v-if="currentMode === 'schedule' && getUnfinishedTasks(data.date).length > 0" class="task-count">
              {{ getUnfinishedTasks(data.date).length }}
            </span>
          </div>
          
          <!-- 显示未完成事项（仅日程模式） -->
          <div class="task-list" v-if="currentMode === 'schedule' && data.type === 'current-month'">
            <div 
              v-for="task in getUnfinishedTasks(data.date).slice(0, 3)" 
              :key="task.id"
              class="task-item"
              :class="'priority-' + task.priority"
            >
              {{ task.content.length > 8 ? task.content.slice(0, 8) + '...' : task.content }}
            </div>
            <div 
              v-if="getUnfinishedTasks(data.date).length > 3" 
              class="task-more"
            >
              还有 {{ getUnfinishedTasks(data.date).length - 3 }} 项...
            </div>
          </div>
          
          <!-- 新闻模式提示 -->
          <div class="news-hint" v-if="currentMode === 'news' && data.type === 'current-month'">
            <el-icon><Reading /></el-icon>
            <span>点击查看新闻</span>
          </div>
        </div>
      </template>
    </el-calendar>
    
    <!-- 加载提示 -->
    <div v-if="scheduleStore.monthLoading" class="loading-overlay">
      <el-icon class="is-loading"><Loading /></el-icon>
    </div>
  </el-card>
  
  <!-- 详情对话框 -->
  <el-dialog
    v-model="detailDialogVisible"
    :title="dialogTitle"
    width="700px"
    class="detail-dialog"
  >
    <!-- 对话框内的模式切换 -->
    <div class="dialog-mode-switch">
      <el-segmented v-model="currentMode" @change="handleModeChange" :options="[
        { label: '日程', value: 'schedule', icon: Notebook },
        { label: '新闻', value: 'news', icon: Reading }
      ]" />
    </div>
    
    <div class="dialog-content">
      <!-- 日程列表 -->
      <template v-if="currentMode === 'schedule'">
        <el-empty v-if="selectedDaySchedules.length === 0" description="暂无日程安排">
          <el-button type="primary" @click="handleAdd">
            <el-icon><Plus /></el-icon>
            添加日程
          </el-button>
        </el-empty>
        
        <div v-else class="schedule-list">
          <div 
            v-for="schedule in selectedDaySchedules" 
            :key="schedule.id"
            class="schedule-item"
            :class="{ completed: schedule.status === 4 }"
          >
            <el-checkbox
              :model-value="schedule.status === 4"
              @change="toggleComplete(schedule)"
              size="large"
            />
            
            <div class="item-content">
              <div class="item-header">
                <span class="time-range">
                  <el-icon><Clock /></el-icon>
                  {{ formatTime(schedule.start_time) }} - {{ formatTime(schedule.end_time) }}
                </span>
                <el-tag :type="statusMap[schedule.status]?.type" size="small">
                  {{ statusMap[schedule.status]?.label }}
                </el-tag>
                <el-tag :type="priorityMap[schedule.priority]?.type" size="small" effect="plain">
                  {{ priorityMap[schedule.priority]?.label }}优先级
                </el-tag>
              </div>
              <div class="item-text" :class="{ 'is-completed': schedule.status === 4 }">
                {{ schedule.content }}
              </div>
            </div>
            
            <el-button :icon="Edit" circle size="small" @click="handleEdit(schedule)" />
          </div>
        </div>
      </template>
      
      <!-- 新闻列表 -->
      <template v-else>
        <div v-if="newsStore.loading" class="news-loading">
          <el-icon class="is-loading"><Loading /></el-icon>
          <span>加载中...</span>
        </div>
        
        <el-empty v-else-if="newsStore.newsList.length === 0" description="暂无新闻" />
        
        <div v-else class="news-list">
          <div 
            v-for="news in newsStore.newsList" 
            :key="news.id"
            class="news-item"
            @click="openNews(news)"
          >
            <div class="news-cover" v-if="news.cover">
              <el-image :src="news.cover" fit="cover" />
            </div>
            <div class="news-content">
              <h4 class="news-title">{{ news.title }}</h4>
              <p class="news-desc" v-if="news.desc">{{ news.desc }}</p>
              <div class="news-meta">
                <span class="news-source">{{ news.source }}</span>
                <span class="news-time">{{ formatPublishTime(news.publish_time) }}</span>
                <span class="news-creator" v-if="news.creator">{{ news.creator }}</span>
              </div>
              <div class="news-stats">
                <span><el-icon><View /></el-icon> {{ news.read_num || 0 }}</span>
                <span><el-icon><ChatDotRound /></el-icon> {{ news.comment_num || 0 }}</span>
                <span><el-icon><Star /></el-icon> {{ news.like_num || 0 }}</span>
              </div>
            </div>
          </div>
        </div>
      </template>
    </div>
    
    <template #footer>
      <el-button @click="detailDialogVisible = false">关闭</el-button>
      <el-button v-if="currentMode === 'schedule'" type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        添加日程
      </el-button>
    </template>
  </el-dialog>
</template>

<script>
import { Loading } from '@element-plus/icons-vue'
export default {
  components: { Loading }
}
</script>

<style scoped>
.calendar-card {
  border-radius: 12px;
  position: relative;
}

/* 模式切换 */
.mode-switch {
  display: flex;
  justify-content: center;
  margin-bottom: 16px;
  padding-bottom: 16px;
  border-bottom: 1px solid var(--el-border-color-light);
}

.custom-calendar {
  --el-calendar-border: none;
  --el-calendar-header-border-bottom: none;
  border: none;
}

/* 头部样式 */
.calendar-header {
  display: flex;
  flex-direction: column;
  gap: 12px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--el-border-color-light);
}

.header-nav {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
}

.header-title {
  font-size: 18px;
  font-weight: 600;
  color: #303133;
  text-align: center;
  flex: 1;
}

.header-nav .el-icon + .el-icon {
  margin-left: -6px;
}

/* 隐藏默认头部 */
:deep(.el-calendar__header) {
  display: none;
}

:deep(.el-calendar__body) {
  padding: 12px 0 0 0;
}

:deep(.el-calendar-table) {
  thead th {
    color: #606266;
    font-weight: 600;
    font-size: 14px;
    padding: 12px 0;
  }
}

:deep(.el-calendar-table td) {
  border: 1px solid #ebeef5 !important;
  vertical-align: top;
}

:deep(.el-calendar-table .el-calendar-day) {
  height: auto;
  min-height: 120px;
  padding: 0;
}

/* 日期单元格样式 */
.date-cell {
  min-height: 110px;
  padding: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.date-cell:hover {
  background: #f5f7fa;
}

.date-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 6px;
}

.date-number {
  font-size: 16px;
  font-weight: 600;
  color: #303133;
}

.task-count {
  background: #409eff;
  color: white;
  font-size: 11px;
  padding: 2px 6px;
  border-radius: 10px;
  font-weight: 500;
}

.date-cell.is-other-month {
  opacity: 0.4;
}

.date-cell.is-other-month .date-number {
  color: #c0c4cc;
}

.date-cell.is-today {
  background: #e6f7e6;
}

.date-cell.is-today .date-number {
  color: #67c23a;
  font-weight: 700;
}

.date-cell.is-selected {
  background: #ecf5ff;
  box-shadow: inset 0 0 0 2px #409eff;
}

.date-cell.is-selected .date-number {
  color: #409eff;
}

.date-cell.is-selected.is-today {
  background: #e6f7e6;
  box-shadow: inset 0 0 0 2px #67c23a;
}

/* 未完成事项列表 */
.task-list {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
}

.task-item {
  font-size: 12px;
  padding: 3px 6px;
  border-radius: 4px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  line-height: 1.4;
}

/* 优先级颜色 */
.task-item.priority-1 {
  background: #e6f7e6;
  color: #67c23a;
  border-left: 3px solid #67c23a;
}

.task-item.priority-2 {
  background: #fdf6ec;
  color: #e6a23c;
  border-left: 3px solid #e6a23c;
}

.task-item.priority-3 {
  background: #fef0f0;
  color: #f56c6c;
  border-left: 3px solid #f56c6c;
}

.task-more {
  font-size: 11px;
  color: #909399;
  padding: 2px 6px;
}

/* 新闻提示 */
.news-hint {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
  flex: 1;
  color: #909399;
  font-size: 12px;
}

.news-hint .el-icon {
  font-size: 20px;
  color: #c0c4cc;
}

/* 加载遮罩 */
.loading-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 12px;
  z-index: 10;
}

.loading-overlay .el-icon {
  font-size: 32px;
  color: #409eff;
}

/* 对话框样式 */
.dialog-mode-switch {
  display: flex;
  justify-content: center;
  margin-bottom: 20px;
}

.dialog-content {
  max-height: 60vh;
  overflow-y: auto;
}

/* 日程列表样式 */
.schedule-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.schedule-item {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  padding: 16px;
  background: #fafafa;
  border: 1px solid #ebeef5;
  border-radius: 8px;
  transition: all 0.2s;
}

.schedule-item:hover {
  background: #f5f7fa;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.schedule-item.completed {
  opacity: 0.6;
  background: #f0f9eb;
}

.item-content {
  flex: 1;
  min-width: 0;
}

.item-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
  flex-wrap: wrap;
}

.time-range {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #409eff;
  font-weight: 600;
  font-size: 13px;
}

.item-text {
  font-size: 15px;
  color: #303133;
  line-height: 1.6;
  word-break: break-word;
}

.item-text.is-completed {
  text-decoration: line-through;
  color: #909399;
}

/* 新闻列表样式 */
.news-loading {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px;
  color: #909399;
  gap: 12px;
}

.news-loading .el-icon {
  font-size: 32px;
  color: #409eff;
}

.news-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.news-item {
  display: flex;
  gap: 16px;
  padding: 16px;
  background: #fafafa;
  border: 1px solid #ebeef5;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s;
}

.news-item:hover {
  background: #f5f7fa;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  border-color: #409eff;
}

.news-cover {
  flex-shrink: 0;
  width: 120px;
  height: 80px;
  border-radius: 6px;
  overflow: hidden;
}

.news-cover .el-image {
  width: 100%;
  height: 100%;
}

.news-content {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.news-title {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #303133;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.news-desc {
  margin: 0;
  font-size: 13px;
  color: #606266;
  line-height: 1.5;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.news-meta {
  display: flex;
  gap: 12px;
  font-size: 12px;
  color: #909399;
}

.news-source {
  color: #409eff;
}

.news-stats {
  display: flex;
  gap: 16px;
  font-size: 12px;
  color: #909399;
}

.news-stats span {
  display: flex;
  align-items: center;
  gap: 4px;
}
</style>

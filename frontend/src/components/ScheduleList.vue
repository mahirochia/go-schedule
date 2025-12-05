<script setup>
import { useScheduleStore } from '@/stores/schedule'
import { ElMessage } from 'element-plus'

const emit = defineEmits(['edit'])
const store = useScheduleStore()

const priorityMap = {
  1: { label: '低', type: 'success' },
  2: { label: '中', type: 'warning' },
  3: { label: '高', type: 'danger' }
}

// 状态映射：1-未开始，2-进行中，3-已结束，4-已完成
const statusMap = {
  1: { label: '未开始', type: 'info' },
  2: { label: '进行中', type: 'primary' },
  3: { label: '已结束', type: 'warning' },
  4: { label: '已完成', type: 'success' }
}

function getPriorityInfo(priority) {
  return priorityMap[priority] || priorityMap[1]
}

function getStatusInfo(status) {
  return statusMap[status] || statusMap[1]
}

function formatTime(time) {
  if (!time) return '--:--'
  
  // 处理 ISO 8601 格式：2025-12-02T18:09:00+08:00
  if (time.includes('T')) {
    const timePart = time.split('T')[1]
    const timeWithoutTz = timePart.split('+')[0]
    return timeWithoutTz.substring(0, 5)
  }
  
  // 处理传统格式：2025-12-01 11:00:00
  if (time.includes(' ')) {
    const timePart = time.split(' ')[1]
    return timePart.substring(0, 5)
  }
  
  // 如果已经是时间格式
  if (time.includes(':')) {
    return time.substring(0, 5)
  }
  
  return '--:--'
}

// 快速切换到已完成状态
async function toggleComplete(schedule) {
  const newStatus = schedule.status === 4 ? 1 : 4
  const result = await store.toggleScheduleStatus(schedule, newStatus)
  if (result.success) {
    ElMessage.success(newStatus === 4 ? '已标记为完成' : '已标记为未开始')
  }
}

function handleEdit(schedule) {
  emit('edit', schedule)
}
</script>

<template>
  <el-card class="schedule-list-card" shadow="never">
    <template #header>
      <div class="card-header">
        <div class="header-title">
          <el-icon size="22" color="#409eff"><Calendar /></el-icon>
          <span>{{ store.formattedDate }}</span>
        </div>
        <el-tag type="info" effect="plain">{{ store.scheduleList.length }} 项日程</el-tag>
      </div>
    </template>
    
    <div v-loading="store.loading" class="list-content">
      <el-empty 
        v-if="!store.loading && store.scheduleList.length === 0"
        description="暂无日程安排"
      >
        <el-text type="info" size="small">点击右下角按钮添加新日程</el-text>
      </el-empty>
      
      <TransitionGroup name="list" tag="div" class="items-container" v-else>
        <div 
          v-for="schedule in store.scheduleList" 
          :key="schedule.id"
          class="schedule-item"
          :class="{ completed: schedule.status === 4 }"
        >
          <div class="item-left">
            <el-checkbox
              :model-value="schedule.status === 4"
              @change="toggleComplete(schedule)"
              size="large"
            />
          </div>
          
          <div class="item-content">
            <div class="item-header">
              <el-text class="time-range" type="primary">
                <el-icon><Clock /></el-icon>
                {{ formatTime(schedule.start_time) }} - {{ formatTime(schedule.end_time) }}
              </el-text>
              <el-tag 
                :type="getStatusInfo(schedule.status).type"
                size="small"
              >
                {{ getStatusInfo(schedule.status).label }}
              </el-tag>
              <el-tag 
                :type="getPriorityInfo(schedule.priority).type"
                size="small"
                effect="plain"
              >
                {{ getPriorityInfo(schedule.priority).label }}优先级
              </el-tag>
            </div>
            <el-text class="item-text" :class="{ 'is-completed': schedule.status === 4 }">
              {{ schedule.content }}
            </el-text>
          </div>
          
          <el-button 
            :icon="Edit"
            circle
            @click="handleEdit(schedule)"
          />
        </div>
      </TransitionGroup>
    </div>
  </el-card>
</template>

<script>
import { Edit, Calendar, Clock } from '@element-plus/icons-vue'
export default {
  components: { Edit, Calendar, Clock }
}
</script>

<style scoped>
.schedule-list-card {
  border-radius: 12px;
  min-height: 400px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 18px;
  font-weight: 600;
  color: #303133;
}

.list-content {
  min-height: 300px;
}

.items-container {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.schedule-item {
  display: flex;
  align-items: flex-start;
  gap: 14px;
  padding: 16px;
  background: #fafafa;
  border: 1px solid #ebeef5;
  border-radius: 12px;
  transition: all 0.2s ease;
}

.schedule-item:hover {
  background: #f5f7fa;
  border-color: #dcdfe6;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.schedule-item.completed {
  opacity: 0.7;
  background: #f0f9eb;
}

.item-left {
  padding-top: 2px;
}

.item-content {
  flex: 1;
  min-width: 0;
}

.item-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 8px;
  flex-wrap: wrap;
}

.time-range {
  display: flex;
  align-items: center;
  gap: 4px;
  font-weight: 600;
  font-size: 13px;
}

.item-text {
  font-size: 15px;
  line-height: 1.6;
  color: #303133;
  word-break: break-word;
}

.item-text.is-completed {
  text-decoration: line-through;
  color: #909399;
}

/* List transition animations */
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}
</style>

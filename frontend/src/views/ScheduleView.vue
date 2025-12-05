<script setup>
import { ref } from 'vue'
import { useScheduleStore } from '@/stores/schedule'
import Calendar from '@/components/Calendar.vue'
import ScheduleForm from '@/components/ScheduleForm.vue'
import { Calendar as CalendarIcon } from '@element-plus/icons-vue'

const store = useScheduleStore()

const showForm = ref(false)
const editingSchedule = ref(null)

function openAddForm() {
  editingSchedule.value = null
  showForm.value = true
}

function openEditForm(schedule) {
  editingSchedule.value = schedule
  showForm.value = true
}

function closeForm() {
  showForm.value = false
  editingSchedule.value = null
}
</script>

<template>
  <div class="schedule-view">
    <header class="app-header">
      <div class="logo">
        <el-icon :size="36" color="#409eff">
          <CalendarIcon />
        </el-icon>
        <h1>日程管理</h1>
      </div>
      <el-text type="info" size="large">高效管理您的每一天</el-text>
    </header>
    
    <main class="main-content">
      <Calendar @add="openAddForm" @edit="openEditForm" />
    </main>
    
    <ScheduleForm 
      :visible="showForm"
      :schedule="editingSchedule"
      @close="closeForm"
      @success="closeForm"
    />
  </div>
</template>

<style scoped>
.schedule-view {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #e4e7ed 100%);
  padding: 32px;
}

.app-header {
  text-align: center;
  margin-bottom: 32px;
}

.logo {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 14px;
  margin-bottom: 8px;
}

.logo h1 {
  margin: 0;
  font-size: 32px;
  font-weight: 700;
  color: #303133;
  letter-spacing: -0.5px;
}

.main-content {
  max-width: 1200px;
  margin: 0 auto;
}

@media (max-width: 768px) {
  .schedule-view {
    padding: 16px;
  }
  
  .logo h1 {
    font-size: 24px;
  }
}
</style>

<script setup>
import { ref, watch, computed } from 'vue'
import { useScheduleStore } from '@/stores/schedule'
import { ElMessage } from 'element-plus'

const props = defineProps({
  schedule: {
    type: Object,
    default: null
  },
  visible: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['close', 'success', 'update:visible'])

const store = useScheduleStore()
const formRef = ref(null)

const form = ref({
  content: '',
  start: null,
  end: null,
  priority: 1,
  status: 1
})

const isEdit = computed(() => !!props.schedule)

// 状态选项：1-未开始，2-进行中，3-已结束，4-已完成
const statusOptions = [
  { value: 1, label: '未开始', type: 'info' },
  { value: 2, label: '进行中', type: 'primary' },
  { value: 3, label: '已结束', type: 'warning' },
  { value: 4, label: '已完成', type: 'success' }
]

const dialogTitle = computed(() => isEdit.value ? '编辑日程' : '新增日程')

const rules = {
  content: [
    { required: true, message: '请输入日程内容', trigger: 'blur' }
  ],
  start: [
    { required: true, message: '请选择开始时间', trigger: 'change' }
  ],
  end: [
    { required: true, message: '请选择结束时间', trigger: 'change' }
  ]
}

// 格式化数字为两位
function padZero(num) {
  return String(num).padStart(2, '0')
}

// 将Date对象格式化为 "YYYY-MM-DD HH:mm:ss" 格式
function formatDateTime(date, year, month, day) {
  if (!date) return ''
  const hours = padZero(date.getHours())
  const minutes = padZero(date.getMinutes())
  const seconds = '00'
  return `${year}-${padZero(month)}-${padZero(day)} ${hours}:${minutes}:${seconds}`
}

// 解析各种时间格式为本地时区的Date对象
function parseTimeToDate(timeStr) {
  if (!timeStr) return null
  
  let hours = 9
  let minutes = 0
  
  // ISO 8601 格式
  if (timeStr.includes('T')) {
    const isoDate = new Date(timeStr)
    if (!isNaN(isoDate.getTime())) {
      hours = isoDate.getHours()
      minutes = isoDate.getMinutes()
    }
  }
  // 传统格式: "2025-12-01 11:00:00"
  else if (timeStr.includes(' ')) {
    const timePart = timeStr.split(' ')[1]
    const parts = timePart.split(':')
    hours = parseInt(parts[0]) || 9
    minutes = parseInt(parts[1]) || 0
  }
  // 纯时间格式: "11:00:00" 或 "11:00"
  else if (timeStr.includes(':')) {
    const parts = timeStr.split(':')
    hours = parseInt(parts[0]) || 9
    minutes = parseInt(parts[1]) || 0
  }
  
  const date = new Date()
  date.setHours(hours, minutes, 0, 0)
  return date
}

// 默认时间
function getDefaultStartTime() {
  const date = new Date()
  date.setHours(9, 0, 0, 0)
  return date
}

function getDefaultEndTime() {
  const date = new Date()
  date.setHours(10, 0, 0, 0)
  return date
}

// 监听schedule变化，填充表单
watch(() => props.schedule, (newVal) => {
  if (newVal) {
    form.value = {
      content: newVal.content || '',
      start: parseTimeToDate(newVal.start_time),
      end: parseTimeToDate(newVal.end_time),
      priority: newVal.priority || 1,
      status: newVal.status || 1
    }
  } else {
    resetForm()
  }
}, { immediate: true })

function resetForm() {
  form.value = {
    content: '',
    start: getDefaultStartTime(),
    end: getDefaultEndTime(),
    priority: 1,
    status: 1
  }
  formRef.value?.resetFields()
}

async function handleSubmit() {
  if (!formRef.value) return
  
  await formRef.value.validate(async (valid) => {
    if (!valid) return
    
    const { year, month, day } = isEdit.value 
      ? { year: props.schedule.year, month: props.schedule.month, day: props.schedule.day }
      : store.selectedDate
    
    const startFormatted = formatDateTime(form.value.start, year, month, day)
    const endFormatted = formatDateTime(form.value.end, year, month, day)
    
    let result
    if (isEdit.value) {
      result = await store.editSchedule({
        id: props.schedule.id,
        year: props.schedule.year,
        month: props.schedule.month,
        day: props.schedule.day,
        start: startFormatted,
        end: endFormatted,
        content: form.value.content,
        status: form.value.status,
        priority: form.value.priority
      })
    } else {
      result = await store.addSchedule({
        content: form.value.content,
        start: startFormatted,
        end: endFormatted,
        priority: form.value.priority,
        status: form.value.status
      })
    }
    
    if (result.success) {
      ElMessage.success(isEdit.value ? '日程更新成功' : '日程添加成功')
      emit('success')
      handleClose()
    } else {
      ElMessage.error(result.msg || '操作失败')
    }
  })
}

function handleClose() {
  resetForm()
  emit('close')
  emit('update:visible', false)
}
</script>

<template>
  <el-dialog
    :model-value="visible"
    :title="dialogTitle"
    width="500px"
    :before-close="handleClose"
    destroy-on-close
  >
    <el-form 
      ref="formRef" 
      :model="form" 
      :rules="rules" 
      label-position="top"
    >
      <el-form-item label="日程内容" prop="content">
        <el-input
          v-model="form.content"
          type="textarea"
          :rows="3"
          placeholder="请输入日程内容..."
          maxlength="200"
          show-word-limit
        />
      </el-form-item>
      
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="开始时间" prop="start">
            <el-time-picker
              v-model="form.start"
              placeholder="选择开始时间"
              format="HH:mm"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="结束时间" prop="end">
            <el-time-picker
              v-model="form.end"
              placeholder="选择结束时间"
              format="HH:mm"
              style="width: 100%"
            />
          </el-form-item>
        </el-col>
      </el-row>
      
      <el-form-item label="优先级">
        <el-radio-group v-model="form.priority" class="full-width-radio">
          <el-radio-button :value="1">
            <el-tag type="success" effect="plain" size="small">低</el-tag>
          </el-radio-button>
          <el-radio-button :value="2">
            <el-tag type="warning" effect="plain" size="small">中</el-tag>
          </el-radio-button>
          <el-radio-button :value="3">
            <el-tag type="danger" effect="plain" size="small">高</el-tag>
          </el-radio-button>
        </el-radio-group>
      </el-form-item>
      
      <el-form-item label="状态">
        <el-radio-group v-model="form.status" class="full-width-radio">
          <el-radio-button 
            v-for="option in statusOptions" 
            :key="option.value" 
            :value="option.value"
          >
            <el-tag :type="option.type" effect="plain" size="small">{{ option.label }}</el-tag>
          </el-radio-button>
        </el-radio-group>
      </el-form-item>
    </el-form>
    
    <template #footer>
      <el-button @click="handleClose">取消</el-button>
      <el-button 
        type="primary" 
        @click="handleSubmit" 
        :loading="store.loading"
      >
        {{ store.loading ? '提交中...' : '确定' }}
      </el-button>
    </template>
  </el-dialog>
</template>

<style scoped>
.full-width-radio {
  width: 100%;
}

.full-width-radio :deep(.el-radio-button) {
  flex: 1;
}

.full-width-radio :deep(.el-radio-button__inner) {
  width: 100%;
  display: flex;
  justify-content: center;
}
</style>

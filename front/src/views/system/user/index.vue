<script setup lang="ts">
import type { ColumnsType } from 'ant-design-vue/es/table'
import { Modal } from 'ant-design-vue'
import { onMounted, ref } from 'vue'
import { useRequest } from 'vue-request'
import { z } from 'zod'
import { addUser, deleteUsers, getAllUsers } from '@/apis/modules/user'
import { useMessage, useZodForm } from '@/hooks'
import { transformItemTyped } from '@/utils/modules/table-data-trans-line'

// Zod 表单验证Schema
const userFormSchema = z.object({
  name: z.string().min(1, '用户名不能为空').max(100, '用户名最多100个字符'),
  role: z.string().min(1, '角色不能为空'),
  project: z.string().min(1, '项目不能为空').max(100, '项目名称最多100个字符'),
  age: z
    .number()
    .int('年龄必须是整数')
    .min(1, '年龄必须大于0')
    .max(150, '年龄必须小于150')
    .optional()
    .or(z.literal(undefined)),
  email: z.string().email('请输入有效的邮箱地址').optional().or(z.literal('')),
})

const columns: ColumnsType = [
  { title: 'ID', dataIndex: 'id', key: 'id' },
  { title: '用户名', dataIndex: 'name', key: 'name' },
  { title: '角色', dataIndex: 'role', key: 'role' },
  { title: '项目', dataIndex: 'project', key: 'project' },
  { title: '年龄', dataIndex: 'age', key: 'age' },
  { title: '邮箱', dataIndex: 'email', key: 'email' },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime' },
  { title: '操作', key: 'action' },
]

const dataSource = ref<User[]>([])
const pagination = ref({
  current: 1,
  pageSize: 10,
  total: 0,
  showSizeChanger: true,
  showTotal: (total: number) => `共 ${total} 条`,
})

const selectedRowKeys = ref<number[]>([])

const { msgSuccess, msgError } = useMessage()

const modalVisible = ref(false)
const modalTitle = ref('')

// 使用 useZodForm
const {
  formData,
  validate,
  validateField,
  reset,
  getFieldError,
  hasFieldError,
  allRules,
} = useZodForm(userFormSchema, {
  name: '',
  role: '',
  project: '',
  age: undefined,
  email: '',
})

// 获取用户列表
const { run: fetchUsers, loading } = useRequest(getAllUsers, {
  manual: true,
  onSuccess: (res) => {
    if (res.data.code === 200) {
      dataSource.value = res.data.data.list.map((item) => {
        return transformItemTyped<User>(item, {
          timeKeys: {
            createTime: 'YYYY-MM-DD HH:mm:ss',
          },
        })
      })
      pagination.value.total = res.data.data.total
      pagination.value.current = res.data.data.page
      pagination.value.pageSize = res.data.data.size
    }
    else {
      msgError({ content: res.data.msg || '获取用户列表失败' })
    }
  },
  onError: (error) => {
    msgError({ content: `获取用户列表失败: ${error.message}` })
  },
})

// 添加用户
const { run: submitAddUser, loading: addLoading } = useRequest(addUser, {
  manual: true,
  onSuccess: (res) => {
    if (res.data.code === 200) {
      msgSuccess({ content: '添加用户成功' })
      modalVisible.value = false
      // 刷新列表
      fetchUsers({
        pageNum: pagination.value.current,
        pageSize: pagination.value.pageSize,
      })
    }
    else {
      msgError({ content: res.data.msg || '添加用户失败' })
    }
  },
  onError: (error) => {
    msgError({ content: `添加用户失败: ${error.message}` })
  },
})

// 删除用户
const { run: submitDeleteUsers, loading: deleteLoading } = useRequest(
  deleteUsers,
  {
    manual: true,
    onSuccess: (res) => {
      if (res.data.code === 200) {
        msgSuccess({ content: '删除成功' })
        selectedRowKeys.value = []
        // 刷新列表
        fetchUsers({
          pageNum: pagination.value.current,
          pageSize: pagination.value.pageSize,
        })
      }
      else {
        msgError({ content: res.data.msg || '删除失败' })
      }
    },
    onError: (error) => {
      msgError({ content: `删除失败: ${error.message}` })
    },
  },
)

// 打开新增弹窗
function handleAdd() {
  modalTitle.value = '新增用户'
  reset()
  modalVisible.value = true
}

// 提交表单
function handleSubmit() {
  const result = validate()

  if (!result.success) {
    msgError({ content: result.errors.issues[0].message })
    return
  }

  // 准备提交数据，过滤掉空值
  const submitData: AddUserParams = {
    name: result.data.name,
    role: result.data.role,
    project: result.data.project,
  }

  if (result.data.age !== undefined && result.data.age !== null) {
    submitData.age = result.data.age
  }

  if (result.data.email && result.data.email.trim() !== '') {
    submitData.email = result.data.email
  }

  submitAddUser(submitData)
}

// 单个删除
function handleDelete(record: User) {
  Modal.confirm({
    title: '确认删除',
    content: `确定要删除用户 "${record.name}" 吗？`,
    okText: '确定',
    cancelText: '取消',
    onOk() {
      submitDeleteUsers([record.id])
    },
  })
}

// 批量删除
function handleBatchDelete() {
  if (selectedRowKeys.value.length === 0) {
    msgError({ content: '请至少选择一条数据' })
    return
  }

  Modal.confirm({
    title: '确认删除',
    content: `确定要删除选中的 ${selectedRowKeys.value.length} 条数据吗？`,
    okText: '确定',
    cancelText: '取消',
    onOk() {
      submitDeleteUsers(selectedRowKeys.value)
    },
  })
}

// 选择改变
function onSelectChange(keys: number[]) {
  selectedRowKeys.value = keys
}

// 分页改变
function handleTableChange(pag: any) {
  pagination.value.current = pag.current
  pagination.value.pageSize = pag.pageSize
  fetchUsers({
    pageNum: pag.current,
    pageSize: pag.pageSize,
  })
}

onMounted(() => {
  fetchUsers({
    pageNum: pagination.value.current,
    pageSize: pagination.value.pageSize,
  })
})
</script>

<template>
  <div>
    <div class="mb-4 flex items-center justify-between">
      <div class="flex gap-2 items-center">
        <h2 class="text-xl text-gray-800 font-semibold">
          用户管理
        </h2>
        <AButton
          v-if="selectedRowKeys.length > 0"
          danger
          :loading="deleteLoading"
          @click="handleBatchDelete"
        >
          <template #icon>
            <CustomIcon
              icon="material-symbols:delete-outline"
              class="m-[0_5px_3px_0] inline-block"
            />
          </template>
          批量删除 ({{ selectedRowKeys.length }})
        </AButton>
      </div>
      <AButton type="primary" @click="handleAdd">
        <template #icon>
          <CustomIcon
            icon="material-symbols:add"
            class="color-white m-[0_5px_3px_0] inline-block"
          />
        </template>
        添加用户
      </AButton>
    </div>

    <ATable
      :columns="columns"
      :data-source="dataSource"
      :loading="loading"
      row-key="id"
      :pagination="pagination"
      :scroll="{ x: 1000 }"
      :row-selection="{
        selectedRowKeys,
        onChange: (selectedRowKeys) =>
          onSelectChange(selectedRowKeys.map(Number)),
      }"
      @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'age'">
          {{ record.age || '-' }}
        </template>

        <template v-else-if="column.key === 'email'">
          {{ record.email || '-' }}
        </template>

        <template v-else-if="column.key === 'action'">
          <ASpace>
            <AButton
              type="link"
              size="small"
              danger
              @click="handleDelete(record as any)"
            >
              删除
            </AButton>
          </ASpace>
        </template>
      </template>
    </ATable>

    <!-- 用户表单弹窗 -->
    <AModal
      v-model:open="modalVisible"
      :title="modalTitle"
      width="600px"
      :confirm-loading="addLoading"
      @ok="handleSubmit"
    >
      <AForm
        :model="formData"
        :rules="allRules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 18 }"
        class="mt-4"
      >
        <AFormItem
          label="用户名"
          name="name"
          :validate-status="hasFieldError('name') ? 'error' : ''"
          :help="getFieldError('name')"
        >
          <AInput
            v-model:value="formData.name"
            placeholder="请输入用户名 (必填)"
            @blur="() => validateField('name')"
          />
        </AFormItem>

        <AFormItem
          label="角色"
          name="role"
          :validate-status="hasFieldError('role') ? 'error' : ''"
          :help="getFieldError('role')"
        >
          <ASelect
            v-model:value="formData.role"
            placeholder="请选择角色 (必填)"
            @blur="() => validateField('role')"
          >
            <ASelectOption value="admin">
              管理员 (admin)
            </ASelectOption>
            <ASelectOption value="editor">
              编辑者 (editor)
            </ASelectOption>
            <ASelectOption value="guest">
              访客 (guest)
            </ASelectOption>
          </ASelect>
        </AFormItem>

        <AFormItem
          label="项目"
          name="project"
          :validate-status="hasFieldError('project') ? 'error' : ''"
          :help="getFieldError('project')"
        >
          <AInput
            v-model:value="formData.project"
            placeholder="请输入项目名称 (必填)"
            @blur="() => validateField('project')"
          />
        </AFormItem>

        <AFormItem
          label="年龄"
          name="age"
          :validate-status="hasFieldError('age') ? 'error' : ''"
          :help="getFieldError('age')"
        >
          <AInputNumber
            v-model:value="formData.age"
            placeholder="请输入年龄 (选填)"
            :min="1"
            :max="150"
            class="w-full"
            @blur="() => validateField('age')"
          />
        </AFormItem>

        <AFormItem
          label="邮箱"
          name="email"
          :validate-status="hasFieldError('email') ? 'error' : ''"
          :help="getFieldError('email')"
        >
          <AInput
            v-model:value="formData.email"
            placeholder="请输入邮箱 (选填)"
            @blur="() => validateField('email')"
          />
        </AFormItem>
      </AForm>
    </AModal>
  </div>
</template>

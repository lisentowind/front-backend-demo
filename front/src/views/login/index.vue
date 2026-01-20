<script setup lang="ts">
import type { LoginParams } from '@/types/user'
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMessage } from '@/hooks'

import { useUserStore } from '@/store/modules/user'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()
const { msgWarning } = useMessage()

const loading = ref(false)
const formState = ref<LoginParams>({
  username: 'admin',
  password: '123456',
})

async function handleLogin() {
  try {
    loading.value = true

    if (!formState.value.username || !formState.value.password) {
      msgWarning({ content: '请输入用户名和密码' })
      return
    }

    await userStore.login(formState.value)

    // 获取重定向路径
    const redirect = route.query.redirect as string
    const path = redirect || userStore.getHomePath

    // 延迟跳转，让用户看到成功提示和过渡动画更自然
    setTimeout(() => {
      router.push(path)
    }, 600)
  }
  catch (error) {
    console.error('Login error:', error)
  }
  finally {
    loading.value = false
  }
}

async function handleRegister() {
  try {
    loading.value = true

    if (!formState.value.username || !formState.value.password) {
      msgWarning({ content: '请输入用户名和密码' })
      return
    }

    // 执行初始化，后端会检查账户是否存在且未设置密码
    await userStore.register(formState.value)
  }
  catch (error) {
    console.error('Register error:', error)
  }
  finally {
    loading.value = false
  }
}
</script>

<template>
  <div
    class="flex min-h-screen items-center justify-center from-primary-2 to-primary-6 bg-gradient-to-br"
  >
    <div class="flex items-center justify-center">
      <!-- 左侧：登录表单 -->
      <div class="p-8 rounded-lg bg-white shadow-2xl">
        <div class="mb-8 text-center">
          <div class="mb-4 flex justify-center">
            <CustomIcon
              icon="material-symbols:admin-panel-settings"
              :width="64"
              class="text-primary-8"
            />
          </div>
          <h1 class="text-3xl text-gray-800 font-bold">
            后台管理系统
          </h1>
          <p class="text-gray-500 mt-2">
            基于 Vue3 + TypeScript + Ant Design Vue
          </p>
        </div>

        <AForm :model="formState" @finish="handleLogin">
          <AFormItem>
            <AInput
              v-model:value="formState.username"
              size="large"
              placeholder="请输入用户名"
            >
              <template #prefix>
                <CustomIcon
                  icon="material-symbols:person-outline"
                  :width="20"
                />
              </template>
            </AInput>
          </AFormItem>

          <AFormItem>
            <AInputPassword
              v-model:value="formState.password"
              size="large"
              placeholder="请输入密码"
            >
              <template #prefix>
                <CustomIcon icon="material-symbols:lock-outline" :width="20" />
              </template>
            </AInputPassword>
          </AFormItem>

          <AFormItem>
            <AButton
              type="primary"
              html-type="submit"
              size="large"
              :loading="loading"
              block
            >
              登录
            </AButton>
          </AFormItem>

          <AFormItem>
            <AButton
              size="large"
              :loading="loading"
              block
              @click="handleRegister"
            >
              初始化账户
            </AButton>
          </AFormItem>
        </AForm>
      </div>
    </div>
  </div>
</template>

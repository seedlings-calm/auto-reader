<template>
    <div class="p-6">
        <h1 class="text-2xl mb-4">配置参数</h1>

        <input v-model="username" class="border p-2 w-full mb-2" placeholder="账号" />
        <input v-model="password" type="password" class="border p-2 w-full mb-2" placeholder="密码" />
        <input v-model="startUrl" class="border p-2 w-full mb-2" placeholder="起始文章链接" />
        <input v-model="readSeconds" type="number" class="border p-2 w-full mb-4" placeholder="每篇文章阅读秒数" />

        <button @click="save" class="bg-green-600 text-white px-4 py-2 mr-4">保存配置</button>
        <button @click="start" class="bg-blue-600 text-white px-4 py-2">开始阅读</button>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { ConfigService } from '../../bindings/changeme/'

const username = ref('')
const password = ref('')
const startUrl = ref('')
const readSeconds = ref(15)

const save = async () => {
    await ConfigService.SaveConfig({
        username: username.value,
        password: password.value,
        start_url: startUrl.value,
        read_seconds: readSeconds.value,
    })
    alert('配置保存成功')
}

const start = () => {
    window.location.href = '/run'
}
</script>

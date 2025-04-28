<template>
    <div class="p-6">
        <h1 class="text-2xl mb-4">阅读中...</h1>
        <p v-if="log">{{ log }}</p>
        <button @click="startReading" class="bg-purple-600 text-white px-4 py-2 mt-4">启动流程</button>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { ReaderService } from '../../bindings/changeme'

const log = ref('')

const startReading = async () => {
    log.value = '正在启动自动阅读流程...'
    const result = await ReaderService.StartReading()
    log.value = result
}
</script>

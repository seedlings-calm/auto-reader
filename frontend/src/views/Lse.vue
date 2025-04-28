<template>
    <div class="p-6">
        <h1 class="text-2xl mb-4">输入授权密钥</h1>
        <input v-model="key" class="border p-2 w-full mb-4" placeholder="请输入密钥，如 VIP-XXXX" />
        <button @click="verify" class="bg-blue-600 text-white px-4 py-2">验证</button>
        <p v-if="msg" class="mt-4 text-green-600">{{ msg }}</p>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { Auth } from '../../bindings/changeme'

const key = ref('')
const msg = ref('')

const verify = async () => {
    const valid = await Auth.VerifyKey(key.value)
    if (valid) {
        await Auth.SaveLicense(key.value)
        msg.value = '验证成功，授权已生效'
        // 跳转配置页
        window.location.href = '/config'
    } else {
        msg.value = '密钥无效，请重新输入'
    }
}
</script>

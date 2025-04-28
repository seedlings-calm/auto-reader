import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHashHistory } from 'vue-router'
import License from './views/Lse.vue'
import Config from './views/Config.vue'
import Runner from './views/Runner.vue'
const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        { path: '/', component: License },
        { path: '/run', component: Runner },
        { path: '/config', component: Config },
    ],
})
createApp(App).use(router).mount('#app')

<script setup lang="ts">
import Lse from './views/Lse.vue'
import Config from './views/Config.vue'
import Runner from './views/Runner.vue'
</script>

<template>
  <div class="container">
    <div>
      <a wml-openURL="https://wails.io">
        <img src="/wails.png" class="logo" alt="Wails logo" />
      </a>
      <a wml-openURL="https://vuejs.org/">
        <img src="/vue.svg" class="logo vue" alt="Vue logo" />
      </a>
    </div>
    <Lse msg="Auth" />
    <Config msg="" />
    <Runner msg="Wails + Vue" />
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #e80000aa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>

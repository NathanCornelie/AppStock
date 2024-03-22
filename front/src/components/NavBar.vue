<template>
  <div>
    <v-app-bar :elevation="2">
      <template v-slot:prepend>
        <v-app-bar-nav-icon>
         <router-link to="Home"><img src="/vite.svg" class="logo" alt="Vite logo" /></router-link>
        </v-app-bar-nav-icon>
      </template> <v-app-bar-title>
      App Bar
    </v-app-bar-title>
    </v-app-bar>
  </div>
</template>

<script setup lang="ts">

</script>

<style lang="scss" scoped>

</style>
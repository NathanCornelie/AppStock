<script setup lang="ts">
import NavBar from "./components/NavBar.vue";
</script>

<template>
  <v-app class="bg-grey-lighten-4">
    <NavBar />
    <div class="content">

    </div>
    <router-view/>

  </v-app>

</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
.content{
  margin-top: 64px;
}
</style>

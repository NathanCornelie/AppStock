<template>
    <v-app-bar :elevation="2">
      <template v-slot:prepend >
        <div class="d-flex justify-space-between w-100">
          <div class="d-flex align-center">
            <v-app-bar-nav-icon>
              <router-link to="/"><img src="/vite.svg" class="logo" alt="Vite logo" /></router-link>
            </v-app-bar-nav-icon>
            <h2> Application</h2>
          </div>

        </div>
      </template>
      <template v-slot:append>
        <div>
          <button class="border-b bg-blue-accent-3 mr-3 " >
            Products
          </button>
          <button class="border-b bg-purple-accent-4 " @click="showClientModale()">
            New Document
          </button>
        </div>
      </template>
    </v-app-bar>
</template>

<script setup lang="ts">

import {ClientStore} from "../stores/clientStore.ts";

const clientStore = ClientStore()
const showClientModale = function (){
  clientStore.setIsClientModale(true)
}

</script>

<style lang="scss" scoped>

</style>
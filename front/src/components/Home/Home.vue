
<script setup lang="ts">

import Stats from "./Stats.vue";
import Documents from "./Documents.vue";
import Client from "../Clients/Client.vue";

</script>

<template>
  <div class="d-flex align-content-space-evenly " style="min-height: 580px">
    <Stats class="flex-1-0 ma-4 "/>
    <Documents  class="flex-1-0 ma-4 "/>
  </div>
  <div>
    <Client />
  </div>

</template>

<style></style>
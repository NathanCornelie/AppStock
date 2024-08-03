<script setup lang="ts">
import {ClientStore} from "../../stores/clientStore.ts";
import ListClients from "../Clients/ListClients.vue";
import {Ref, ref} from "vue";
import CreateClient from "../Clients/CreateClient.vue";

const clientStore = ClientStore(); 
const isCreate: Ref<boolean> = ref(false)
const snackbar = ref(false)
const snackbarText = ref("")
const timeout = 3000

const changeIsCreate = () => {
  isCreate.value = !isCreate.value
  console.log(isCreate.value)
}

const snackClientCreated = async (text: string) => {
  snackbarText.value = text
  snackbar.value = true
  isCreate.value = false
  await delay(3000)
  snackbar.value = false
}

function delay(ms: number) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

</script>
<template>
  <v-dialog style="max-width: 1000px" v-model="clientStore.isClientModale" persistent template>
    <v-card>
      <div class="d-flex align-center justify-lg-space-between pa-3"><h2>Clients</h2>
        <v-btn @click="changeIsCreate()"><p v-if="isCreate">All </p>
          <p v-else> Create</p></v-btn>
      </div>
      <div v-if="isCreate">
        <CreateClient v-on:changeIsCreated="changeIsCreate" v-on:snackClientCreated="snackClientCreated"/>
      </div>
      <div v-else>
        <ListClients/>
      </div>
    </v-card>
  </v-dialog>
  <v-snackbar
      v-model="snackbar"
      :timeout="timeout">

    {{ snackbarText }}

    <template v-slot:actions>
      <v-btn
          color="blue"
          variant="text"
          @click="snackbar = false">

        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<style scoped>

</style>
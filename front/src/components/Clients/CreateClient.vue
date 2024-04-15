<script setup lang="ts">
import {CreateClient} from "../../common/models/Client.ts";
import {ClientStore} from "../../stores/clientStore.ts";
import ClientsAPI from "../../common/API/ClientsAPI.ts";
import {Ref, ref, UnwrapRef} from "vue";

const createdClient : Ref<UnwrapRef<CreateClient>> = ref(new CreateClient("","","",""))
const clientStore = ClientStore();

const emit = defineEmits({
  snackClientCreated: (payload: string)=>payload
})
const createClient = async ()=>{
  try{
    const res = await ClientsAPI.CreateClient(createdClient.value)
    if (res) {
      emit("snackClientCreated","Client successfully created")
    }else{
      emit("snackClientCreated","Can't create client")

    }
  }catch (e) {

  }
}

</script>

<template>
  <v-form >
    <v-container>
      <v-col>
        <v-row
            cols="12"
            md="4"
        >
          <v-text-field
              class="py-2 rounded-lg "
              v-model="createdClient.firstname"
              :counter="10"
              label="First name"
              hide-details
              required
          ></v-text-field>
        </v-row>

        <v-row
            cols="12"
            md="4"
        >
          <v-text-field
              class="py-2 rounded-lg "
              v-model="createdClient.lastname"
              :counter="10"
              label="Last name"
              hide-details
              required
          ></v-text-field>
        </v-row>

        <v-row
            cols="12"
            md="4"
        >
          <v-text-field
              v-model="createdClient.email"
              class="py-2 rounded-lg "

              label="E-mail"
              hide-details
              required
          ></v-text-field>
        </v-row>
      </v-col>
    </v-container>
  </v-form>
  <v-card-actions class="justify-end">
    <v-btn color="error"  @click="clientStore.setIsClientModale(false)">Close</v-btn>
    <v-btn  color="primary"  @click="createClient()">Create</v-btn>
  </v-card-actions>

</template>

<style scoped>

</style>
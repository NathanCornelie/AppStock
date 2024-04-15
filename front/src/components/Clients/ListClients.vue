<script setup lang="ts">
import "ag-grid-community/styles/ag-theme-alpine.css"
import "ag-grid-community/styles/ag-grid.css"
import "ag-grid-community/styles/ag-theme-quartz.css"
import {ClientSideRowModelModule, GridReadyEvent} from "ag-grid-community";
import {AgGridVue} from "@ag-grid-community/vue3";
import {onMounted, ref, Ref, UnwrapRef} from "vue";
import {Client} from "../../common/models/Client.ts";
import ClientsAPI from "../../common/API/ClientsAPI.ts";
import {ClientStore} from "../../stores/clientStore.ts";

const clientStore = ClientStore();
const gridOptions = ref({
  enableCellTextSelection:true
})

const cols = [
  {field: "id",width: 230},
  {field: "lastname"},
  {field: "name"},
  {field: "email",flex:1},

];
let clients : Ref<UnwrapRef<Client[]>> = ref([])
const modules = [ClientSideRowModelModule]

onMounted(async ()=>{
  try{
    clients.value = await ClientsAPI.GetClients()

  }catch (e) {
    console.log(e)
  }
})

const onGridReady = async (params:GridReadyEvent) => {
  const updateData =  (data:Client[]) => {
    params.api.setGridOption('rowData', data)
  };
  try{
    clients.value = await ClientsAPI.GetClients()
    updateData(clients.value);
  }catch (e) {
    console.log(e)
  }

}

</script>

<template>
    <ag-grid-vue
        :grid-options = "gridOptions"
        style="height: 500px"
        class="ag-theme-quartz"
        @grid-ready="onGridReady"
        :columnDefs="cols"
        :modules="modules"
    >
    </ag-grid-vue>
  <v-card-actions class="justify-end">
    <v-btn color="error" @click="clientStore.setIsClientModale(false)">Close</v-btn>
  </v-card-actions>

</template>

<style scoped>

</style>
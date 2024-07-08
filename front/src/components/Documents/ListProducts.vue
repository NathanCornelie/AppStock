<script setup lang="ts">
import {AgGridVue} from "@ag-grid-community/vue3";
import {
  CellRendererSelectorResult,
  ClientSideRowModelModule,
  ColDef,
  GridApi,
  GridReadyEvent,
  ICellRendererParams, RowHeightParams,
}
  from "ag-grid-community";
import {onMounted, ref, watch} from "vue";
import { DocumentStore } from "../../stores/documentStore";
import {Product} from "../../common/models/Product.ts";
import ListClientsGridCells from "../../components/Clients/ListClientsGridCells.vue";

const documentStore = DocumentStore()
const onRowClicked = ()=>{}
const cols = ref<ColDef[]>([{field:"id",cellRendererSelector: (params:ICellRendererParams)=>customCell(params,"id")},
  {field:"product",cellRendererSelector: (params:ICellRendererParams)=>customCell(params,"product")},
  {field:"description"},
  {field:"quantity"},
  {field:"price"}]);
const modules = [ClientSideRowModelModule]
const gridApi = ref<GridApi|null>(null)
const pinnedTopRow = ref<Product>(new Product())
const gridOptions = ref({
  enableCellTextSelection: true
})

const customCell = (params: ICellRendererParams,field:string):CellRendererSelectorResult|undefined=>{
  if(params.node.rowPinned)
    return {
      component: ListClientsGridCells,
      params:{
        field: field
      }
    }
  else return undefined

}

onMounted(()=>{

})

watch(documentStore.selectedProducts, ()=>{
   gridApi?.value?.setGridOption('rowData',documentStore.selectedProducts)
})

const onGridReady = (params: GridReadyEvent)=>{
  const updateData = () => {
    params.api.setGridOption('rowData', documentStore.selectedProducts)
  };
  try{
    gridApi.value = params.api
    gridApi.value.setGridOption('getRowHeight',getRowHeight)
    gridApi.value.sizeColumnsToFit()
    gridApi.value.setGridOption('pinnedTopRowData',[pinnedTopRow.value])




    window.addEventListener("resize",handleWindowResize)
    updateData()
  }catch (e){
    console.log(e)
  }
}

const getRowHeight = (params:RowHeightParams)=>{
  return params.node.isRowPinned() ? 80 : 60
}
const handleWindowResize = ()=>{
  if (gridApi.value){
    gridApi.value.sizeColumnsToFit()
  }
}

</script>

<template>
  <v-card style="height: 100%" class="ma-3 rounded-lg">
    <ag-grid-vue
        :grid-options="gridOptions"
        style="height: 100%"
        class="ag-theme-quartz"
        :onRowClicked="onRowClicked"
        @grid-ready="onGridReady"
        :columnDefs="cols"
        :modules="modules"
    >
    </ag-grid-vue>
  </v-card>
</template>

<style scoped>
.ag-row-pinned .ag-floating-top .ag-selectable{
  height: 60px !important;
}

</style>
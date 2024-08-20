<script setup lang="ts">
import {onMounted, PropType, ref, watch} from "vue";
import {ICellRendererParams, } from "ag-grid-community";
import {Product} from "../../common/models/Product.ts";
import ProductsAPI from "../../common/API/ProductsAPI.ts";
import {DocumentStore} from "../../stores/documentStore.ts";

interface ListClientsGridCellsParams extends ICellRendererParams{field:String}
const params = defineProps({params: Object as PropType<ListClientsGridCellsParams>})
const selectedProduct = ref<Product|null>(null);
const products = ref<Product[]>([])
const filteredProducts = ref<Product[]>([])
const field = params.params?.field
onMounted(async ()=>{

  getProducts().then()
})
const documentStore = DocumentStore()
const getProducts = async ()=>{try{
  const resp = await ProductsAPI.GetProducts()
  products.value = resp
  filteredProducts.value = resp
}catch (e){
  console.log("Error getting data :\n"+e)
}}


const filterOnIdFunction = (value: string, query: string, item?: { value: string, raw : Product }):boolean=>{
 if(value){
   if (item?.raw.name.includes(query)) return true
 }
   return false
}
const filterOnProductFunction = (value: string, query: string, item?: { value: string, raw : Product }):boolean=>{
  if(value){

    if ( item?.raw.id.includes(query)) return true }
    return false

}
const handleSelectedProduct = (item:Product)=>{
  let old = documentStore.selectedProducts
  const product = old.find(p=>p.id ===item.id)
  if(product){
    old.push(item)
    documentStore.setSelectedProducts(old)
  }

}

watch(selectedProduct,()=>{
  if(selectedProduct.value) {
    handleSelectedProduct(selectedProduct.value)
    selectedProduct.value = null
  }
})

</script>

<template>
  <div style="height: 80px;display:flex;justify-content: center;align-items: center;">
    <v-autocomplete v-if="field=='id'"
                    label="Autocomplete"
                    :items="filteredProducts"
                    variant="outlined"
                    v-model="selectedProduct"
                    item-text="name"
                    :custom-filter="filterOnIdFunction"
                    hide-details
                    clear-on-select


    >
      <template v-slot:item="{ props, item }">
        <v-list-item
            v-bind="props"
            :subtitle="item.raw.name"
            :title="item.raw.id"
            :on-click="()=>handleSelectedProduct(item.raw)"
        ></v-list-item>
      </template>

    </v-autocomplete>
    <v-autocomplete v-if="field=='product'"
                    label="Autocomplete"
                    v-model="selectedProduct"
                    :items="filteredProducts"
                    variant="outlined"
                    item-text="name"
                    :custom-filter="filterOnProductFunction"
                    hide-details
  >
    <template v-slot:item="{ props, item }">
      <v-list-item
          v-bind="props"
          :subtitle="item.raw.name"
          :title="item.raw.id"
      ></v-list-item>
    </template>

  </v-autocomplete>

  </div>

</template>

<style scoped>

</style>
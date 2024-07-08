<script setup lang="ts">
import {ICellRendererParams} from "ag-grid-community";
import {onMounted, PropType, ref} from "vue";
import {Product} from "../../common/models/Product.ts";
import { VNumberInput } from 'vuetify/labs/VNumberInput'

interface ListClientsGridCellsParams extends ICellRendererParams{field:String,product:Product}
const params = defineProps({params: Object as PropType<ListClientsGridCellsParams>})
const field = params.params?.field
const product = params.params?.product
const quantity= ref<number>(0)
onMounted(()=>{
  console.log(product)
})
</script>
<template>
  <div class="container" v-if="field == 'id'">
    {{ product?.id }}
  </div>
  <div  class="container" v-if="field == 'product'">
    {{ product?.name }}
  </div>
  <div class="container" v-if="field == 'quantity'" style="display: flex;justify-content: space-between;">
    <div style="width: 100%">
      {{ quantity }}
    </div>


    <v-number-input  control-variant="split" label=""
                    v-model="quantity"
                    variant="outlined"
                    :min="0"
                    :max="product?.stock"
                    hide-details
                     hide-input
    >
<!--      <template v-slot:increment><div class="icon_container"><v-icon  icon="mdi-plus"  /></div></template>-->
<!--      <template v-slot:decrement><div  class="icon_container"><v-icon  icon="mdi-minus"  /></div></template>-->

    </v-number-input>
  </div>
</template>
<style scoped lang="scss" >

.container{

  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}

.ag-cell-wrapper{
  display: flex;
  height: 100%;
}


</style>
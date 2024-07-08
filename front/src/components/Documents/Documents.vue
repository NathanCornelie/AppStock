<script setup lang="ts">
import {onMounted, ref, Ref, UnwrapRef} from "vue";
import {ClientStore} from "../../stores/clientStore.ts"
import ClientsAPI from "../../common/API/ClientsAPI.ts";
import {Client} from "../../common/models/Client.ts";
import {Product} from "../../common/models/Product.ts"
import ProductsAPI from "../../common/API/ProductsAPI.ts";
import ListProducts from "../../components/Documents/ListProducts.vue";
import {DocumentStore} from "../../stores/documentStore.ts";

const clientStore = ClientStore();
const documentStore = DocumentStore()
const client : Ref<UnwrapRef<Client>> = ref(new Client("","","",""))
const products: Ref<UnwrapRef<Product[]>> = ref([])

onMounted(async () => {
  getClients().then()
  getProducts().then()
})

const getClients = async ()=>{
  if (clientStore.selectedClientId) {
    try {const resp =await ClientsAPI.GetClientById(clientStore.selectedClientId)
      if(resp) client.value = resp

    } catch (e) {
      console.log("eer"+e)
    }
  }
}
const getProducts = async ()=>{

    try {const resp =await ProductsAPI.GetProducts()
      if(resp) products.value = resp

    } catch (e) {
      console.log(e)
    }
}
const handleProductClick = (item: Product)=>{
  let products = documentStore.selectedProducts
  products.push(item)
  documentStore.setSelectedProducts(products)
}

</script>

<template>
  <div class="d-flex">
    <v-card style="min-height: 300px" class="ma-3 rounded-lg d-flex align-center pa-3">
      <div class="image h_container"/>
      <div class="informations h_container pa-3 d-flex flex-column align-self-center py-8">
        <v-text-field readonly variant="outlined" v-model="client.lastname" label="Last Name" hide-details/>
        <v-text-field readonly variant="outlined" v-model="client.name" label="First Name" hide-details />
        <v-text-field  readonly variant="outlined" v-model="client.email" label="Phone number" hide-details/>
      </div>
    </v-card>
    <v-card style="min-height: 300px; width: 60%" class="ma-3 rounded-lg  pa-2">
      <v-text-field clearable label="Products" variant="outlined" hide-details></v-text-field>
      <v-virtual-scroll :items="products" class="products_container mt-2">
        <template v-slot:default="{ item }">
          <v-card variant="outlined" class="ma-2">
            <div class="d-flex justify-space-between pa-3 border-b h" @click="handleProductClick(item)">
              <p>{{item.name ? item.name : "Product"}}</p>
              <p>{{item.category ? item.category : "no category"}}</p>
              <p>{{item.stock }}</p>
            </div>
          </v-card>

        </template>
      </v-virtual-scroll>
    </v-card>
  </div>

  <ListProducts/>
</template>

<style scoped>
.h_container {
  border: 1px solid #d2d1d1;
  height: 100%;
  border-radius: 15px;
}

.image {
  width: 200px;
}

.informations {
  width: 500px;
  margin-left: 20px;
}
.products_container{
  max-height: 200px;
}
</style>
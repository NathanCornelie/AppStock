<script setup lang="ts">


import {onMounted} from "vue";
import DocumentsAPI from "../../common/API/DocumentsAPI.ts";
import DocumentModel from "@models/Document.ts"
const title = "Documents"
const data = [
  1,2,3
]
let documents : DocumentModel[] = []
onMounted(async ()=>{
  documents = await DocumentsAPI.GetDocumentsByUserId(import.meta.env.VITE_TEST_USER_ID)
  console.log(documents)

})
</script>

<template>
  <v-card :title>
    <v-container class="bg-grey-lighten-4" >
        <v-card v-for="document in documents" :key="document.id" class="my-3 " height="100px"  hover>
          <div class="d-flex justify-space-between px-3">
            <p>Ref : <span> 011548941</span></p>
            <p>{{new Date().toLocaleDateString()}}</p>
          </div>

              <div class="d-flex text-h6 justify-space-between pa-3 align-center ">
                <p class="font-weight-black">Excellium</p>
                <p class="mr-3">3 Products</p>
                <p>280,00</p>


              </div>

        </v-card>

    </v-container>

  </v-card>
</template>

<style scoped>
.doc_container{

}
</style>
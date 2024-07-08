<script setup lang="ts">


import { onMounted, Ref, ref, UnwrapRef} from "vue";
import DocumentsAPI from "../../common/API/DocumentsAPI.ts";
import DocumentModel from "@models/Document.ts";
import Document from "@models/Document.ts";

const title = "Documents"

let documents : Ref<UnwrapRef<Document[]>> = ref<DocumentModel[]>([])
onMounted(async ()=>{
  try{
    documents.value = await DocumentsAPI.GetDocumentsByUserId(import.meta.env.VITE_TEST_USER_ID)

  }catch (e) {
    console.log(e)
  }

})
function getNumberofProducts(doc : Document){
  return  doc.commands? doc.commands.map(c=>c.quantity).reduce((a,b)=>a+b,0):0
}
</script>


<template>
  <v-card :title class="">
    <v-container class="bg-grey-lighten-4 doc_container" >
        <v-card v-for="document in documents.values()" :key="document.id" class="my-3 " height="100px"  hover>
          <div class="d-flex justify-space-between px-3">
            <p>Ref : <span> 011548941</span></p>
            <p>{{new Date().toLocaleDateString()}}</p>
          </div>

              <div class="d-flex text-h6 justify-space-between pa-3 align-center ">
                <p class="font-weight-black">Excellium</p>
                <p class="mr-3">{{getNumberofProducts(document)}} Products</p>
                <p>280,00</p>


              </div>

        </v-card>

    </v-container>

  </v-card>
</template>

<style scoped>
.doc_container{
  min-height: 580px;
}
</style>
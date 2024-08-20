import {defineStore} from 'pinia';
import {Product} from "@models/Product.ts";
import {ref} from "vue";

export const DocumentStore = defineStore('documentStore', {
    state: (): { selectedProducts: Product[], selectedClientId: string | null } => ({

        selectedProducts: [],
        selectedClientId: null
    }),

    getters: {
        getSelectedProducts: (state) => state.selectedProducts,
        getSelectedClient: (state) => state.selectedClientId,
    },

    actions: {
        setSelectedProducts(value: Product[]) {
            this.selectedProducts = value;
        },
        setSelectedClient(value: string) {
            this.selectedClientId = value;
        },

        updateProductQuantity(id:number, quantity:number) {

            const product = this.selectedProducts.find((product) => product.id === id);
            if (product) {
                product.quantity = quantity;
            }

            console.log(this.selectedProducts)
        }
    },

});
import {defineStore} from 'pinia';
import {Product} from "@models/Product.ts";

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
        }
    }
});
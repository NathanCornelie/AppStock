import {defineStore} from 'pinia';

export const ClientStore = defineStore('clientStore', {
    state: (): { isClientModale: boolean, selectedClientId: string | null } => ({
        isClientModale: false,
        selectedClientId: null
    }),

    getters: {
        getIsClientModale: (state) => state.isClientModale,
        getSelectedClient: (state) => state.selectedClientId,
    },

    actions: {
        setIsClientModale(value: boolean) {
            this.isClientModale = value;
        },
        setSelectedClient(value: string) {
            this.selectedClientId = value;
        }
    }
});
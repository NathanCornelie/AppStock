import {defineStore} from 'pinia';

export const ClientStore = defineStore('clientStore', {
    state: () => ({
        isClientModale: false,
        name: "nathan"
    }),
    getters: {
        inClientModale: (state) => state.isClientModale,
        getName: state => state.name
    },
    actions: {
        setIsClientModale(value: boolean) {
            this.isClientModale = value;
        },
        changeName(name: string) {
            this.name = name;
        }

    }
});
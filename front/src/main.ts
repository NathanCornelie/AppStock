import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {createRouter, createWebHistory, Router} from "vue-router";
import Home from "./components/Home/Home.vue";
const vuetify = createVuetify({
    components,
    directives,
})

const routes = [
    {
        path: '/',
        component : Home
    }
]
const router : Router = createRouter({
    history:createWebHistory(),routes})

createApp(App).use(vuetify).use(router).mount('#app')

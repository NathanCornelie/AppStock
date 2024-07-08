import {createApp} from 'vue'
import './style.css'
import App from './App.vue'
import 'vuetify/styles'
import {createVuetify} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {createRouter, createWebHistory, Router} from "vue-router";
import Home from "./components/Home/Home.vue";
import {createPinia} from "pinia";
import Documents from "./components/Documents/Documents.vue";
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import '@mdi/font/css/materialdesignicons.css'
const pinia = createPinia()

const vuetify = createVuetify({
    components,
    directives,
})

const routes = [
    {
        path: '/',
        component: Home
    },
    {
        path: '/documents',
        component: Documents
    }
]
const router: Router = createRouter({
    history: createWebHistory(), routes
})

const app = createApp(App)

app.use(vuetify)
app.use(pinia)
app.use(router)

app.mount('#app')

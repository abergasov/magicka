import { createApp } from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import store from './store'
import router from "@/router";
import axios from './axios'

createApp(App)
    .use(router)
    .use(store)
    .use(axios, {
        baseUrl: '/',
    })
    .mount('#app')

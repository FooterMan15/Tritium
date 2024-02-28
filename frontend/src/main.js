import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import {createRouter, createWebHistory} from 'vue-router'

import Instances from './components/pages/Instances.vue'
import Settings from './components/pages/Settings.vue'
import Account from './components/pages/Account.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: Instances },
        { path: '/settings', component: Settings },
        { path: '/account', component: Account }
    ]
})

const app = createApp(App).use(router).mount('#app')
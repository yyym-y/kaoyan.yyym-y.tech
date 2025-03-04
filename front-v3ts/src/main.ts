import { createApp } from 'vue'
import App from './App.vue'
import router from './router'


const app = createApp(App)

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElIcons from '@element-plus/icons-vue'
app.use(ElementPlus)

import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
app.use(ArcoVue);
import * as ArcoIcons from '@arco-design/web-vue/es/icon';
for (const name in ArcoIcons){
	app.component(name,(ArcoIcons as any)[name])
}

import { createPinia } from 'pinia'
const pinia = createPinia()
app.use(pinia)

for (const name in ElIcons){
	app.component(name,(ElIcons as any)[name])
}

import Vue3TouchEvents, {
	type Vue3TouchEventsOptions,	
} from "vue3-touch-events";
  
app.use<Vue3TouchEventsOptions>(Vue3TouchEvents, {
disableClick: false,
touchHoldTolerance: 700
})

app.use(router).mount('#app')

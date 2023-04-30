import { createApp } from 'vue'
import App from './App.vue'
import { createVuestic } from "vuestic-ui";
import "vuestic-ui/css";
import router from './router'

import './assets/main.css'

const app = createApp(App)

app.use(router)
app.use(createVuestic())

app.mount('#app')

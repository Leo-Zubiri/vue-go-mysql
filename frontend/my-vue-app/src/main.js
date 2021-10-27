import { createApp } from 'vue'
import App from './App.vue'

createApp(App).mount('#app')

//Lo primero al montar la app es cargar la tabla
import TableComp from "./components/TableComp.vue";
createApp(TableComp).mount("#contentShowed");
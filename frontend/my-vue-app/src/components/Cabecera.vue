<template>
  <div>
    <nav class="navbar navbar-expand navbar-light bg-light">
      <div class="nav navbar-nav">
        <a class="nav-item nav-link active" href="#" v-on:click="miMetodo()"
          >Home <span class="sr-only">(current)</span></a
        >
        <a class="nav-item nav-link" href="#" v-on:click="miMetodo2()"
          >Crear Nuevo Usuario</a
        >
        <a class="nav-item nav-link" href="#" v-on:click="miMetodo3()"
          >Editar Usuario</a
        >
      </div>
    </nav>
    <div id="contentShowed"></div>
  </div>
</template>

<script>
import { createApp } from "vue";
import Formulario from "./Formulario.vue";

import TableComp from "./TableComp.vue";
import Search from "./Search.vue";

function miMetodo() {
  createApp(TableComp).mount("#contentShowed");
}
function miMetodo2() {
  createApp(Formulario).mount("#contentShowed");
}

function miMetodo3() {
  createApp(Search).mount("#contentShowed");
}

export default {
  name: "Cabecera",
  methods: {
    miMetodo,
    miMetodo2,
    miMetodo3,
  },
};
</script>

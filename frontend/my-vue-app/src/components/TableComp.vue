<template>
  <div class="table-responsive">
    <button class="btn btn-primary" v-on:click="GET()">Actualizar</button>
    <table class="table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nombre</th>
          <th>Apellido</th>
          <th>Correo</th>
          <th>Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in usuarios" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.nombre }}</td>
          <td>{{ user.apellido }}</td>
          <td>{{ user.email }}</td>
          <td>
            <a
              name=""
              id=""
              class="btn btn-danger"
              v-bind:href="'/delete?id=' + user.id"
              role="button"
              >Eliminar</a
            >
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "TableComp",
  data() {
    return {
      usuarios: [],
    };
  },
  methods: {
    GET() {
      axios({
        method: "GET",
        url: "http://127.0.0.1:3000/get",
        data: "",
        headers: { "content-type": "text/plain" },
      }).then((result) => {
        console.log(result.data);
        this.usuarios = result.data;
        if (this.usuarios.length == 0) alert("Sin Usuarios que mostrar");
      });
    },
  },
};
</script>
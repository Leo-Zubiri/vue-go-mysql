<template>
  <div :key="componentKey">
    <div class="form-group">
      <a
        name=""
        id=""
        class="btn btn-primary"
        role="button"
        v-on:click="setSearch()"
        >Buscar</a
      >
      <input
        type="number"
        class="search"
        name="idSearch"
        id="idSearch"
        aria-describedby="helpId"
        placeholder="ID Usuario"
      />
    </div>

    <div class="card" v-for="user in usuarios" :key="user.id">
      <div class="card-header">Agregar</div>
      <div class="card-body">
        <form>
          <div class="form-group">
            <input
              v-bind:id="'idInput' + user.ID"
              v-model="idDinamic"
              name="idInput"
              type="text"
              placeholder="ID usuario"
              readonly="true"
              class="form-control"
            />
          </div>
          <div class="form-group">
            <input
              v-bind:id="'nombreInput' + user.ID"
              v-model="nombreDinamic"
              name="nombreInput"
              type="text"
              placeholder="First Name"
              class="form-control"
            />
          </div>

          <div class="form-group">
            <input
              v-bind:id="'apeInput' + user.ID"
              v-model="apeDinamic"
              name="apeInput"
              type="text"
              placeholder="Last Name"
              class="form-control"
            />
          </div>

          <div class="form-group">
            <input
              type="email"
              v-model="emailDinamic"
              class="form-control"
              name=""
              v-bind:id="'emailInput' + user.ID"
              aria-describedby="emailHelpId"
              placeholder="Email"
            />
          </div>

          <div>
            <p></p>
            <button class="btn btn-primary" v-on:click="request()">
              SEND!!
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
let idDinamic = "0";
let nombreDinamic = "";
let apeDinamic = "";
let emailDinamic = "";

export default {
  name: "Search",
  data() {
    return {
      usuarios: [],
      idDato: "",
      idDinamic,
      nombreDinamic,
      apeDinamic,
      emailDinamic,
      componentKey: 0,
    };
  },
  methods: {
    setSearch() {
      this.idDato = parseInt(document.getElementById("idSearch").value);
      console.log(this.idDato);
      var data1 = { id: this.idDato, nombre: "a", apellido: "a", email: "a" };
      axios({
        method: "POST",
        url: "http://127.0.0.1:3000/editar",
        data: data1,
        headers: { "content-type": "text/plain" },
      }).then((result) => {
        console.log(result.data);
        this.usuarios = result.data;
        if (this.usuarios.length == 0) {
          alert("Usuario No encontrado");
        } else {
          for (let i = 0; i < this.usuarios.length; i++) {
            idDinamic = this.usuarios[i].id;
            nombreDinamic = this.usuarios[i].nombre;
            apeDinamic = this.usuarios[i].apellido;
            emailDinamic = this.usuarios[i].email;
          }
          this.forceRerender();
        }
      });
    },
    forceRerender() {
      this.componentKey += 1;
    },
  },
};
</script>

<style scoped>
input.search {
  margin-left: 12px;
  margin-top: 30px;
  width: 200px;
}
</style>
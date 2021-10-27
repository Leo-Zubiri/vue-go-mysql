<template>
  <div>
    <div class="card">
      <div class="card-header">Actualizar Info Usuario</div>
      <div class="card-body">
        <form>
          <div class="form-group">
            <input
              id="idInput"
              name="idInput"
              type="text"
              placeholder="ID usuario"
              readonly="true"
              class="form-control"
            />
          </div>
          <div class="form-group">
            <input
              id="nombreInput"
              name="nombreInput"
              type="text"
              placeholder="First Name"
              class="form-control"
            />
          </div>

          <div class="form-group">
            <input
              id="apeInput"
              name="apeInput"
              type="text"
              placeholder="Last Name"
              class="form-control"
            />
          </div>

          <div class="form-group">
            <input
              type="email"
              class="form-control"
              name=""
              id="emailInput"
              aria-describedby="emailHelpId"
              placeholder="Email"
            />
          </div>

          <div>
            <p></p>
            <button class="btn btn-primary" v-on:click="request()">
              Actualizar Informacion
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "FormUpdate",
  methods: {
    request() {
      //var nombreDato = { nombre: nombreInput };
      var nombreDato = document.getElementById("nombreInput").value;
      var apellidoDato = document.getElementById("apeInput").value;
      var email = document.getElementById("emailInput").value;

      var data2 = { nombre: nombreDato, apellido: apellidoDato, email: email };

      console.log(nombreDato);

      console.log(apellidoDato);

      axios({
        method: "UPDATE",
        url: "http://127.0.0.1:3000/editar",
        data: data2,
        headers: { "content-type": "text/plain" },
      }).then((result) => {
        console.log(result.data);
        if (result.statusText == "OK") alert("Usuario Actualizado");
      });
      /*
        .catch((error) => {
          console.error(error);
        });
        */
    },
  },
};
</script>

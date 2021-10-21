<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Send Data to golang backend (server.go)</h1>

        <form v-on:submit.prevent="">
          <div class="form-group">
            <input
              id="nombreInput"
              name="nombreInput"
              type="text"
              placeholder="First Name"
              class="form-control"
            />
          </div>

          <div>
            <input
              id="apeInput"
              name="apeInput"
              type="text"
              placeholder="Last Name"
              class="form-control"
            />
          </div>

          <div>
            <p></p>
            <button class="btn btn-primary" v-on:click="request()">
              SEND!
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
  name: "App",

  data() {
    return {
      websiteUrl: "",
    };
  },

  methods: {
    request() {
      //var nombreDato = { nombre: nombreInput };
      var nombreDato = document.getElementById("nombreInput").value;
      var apellidoDato = document.getElementById("apeInput").value;

      var data2 = { nombre: nombreDato, apellido: apellidoDato };

      console.log(nombreDato);

      console.log(apellidoDato);

      axios({
        method: "POST",
        url: "http://127.0.0.1:3000/data",
        data: data2,
        headers: { "content-type": "text/plain" },
      })
        .then((result) => {
          console.log(result);
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>

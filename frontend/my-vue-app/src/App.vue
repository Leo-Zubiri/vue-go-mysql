

<template>
  <center>
    <div id="app" class="container">
      <VueLogo />
      <Cabecera />
    </div>
  </center>
</template>

<script>
import VueLogo from "./components/VueLogo.vue";
import Cabecera from "./components/Cabecera.vue";

export default {
  components: {
    VueLogo,
    Cabecera,
  },

  name: "App",

  data() {
    return {
      websiteUrl: "",
    };
  },
};
</script>

<template>
  <v-app class="app">
    <v-main>
      <NavBar />
      <v-container>
        <v-row justify="center" class="search-bar"
          ><v-col cols="10">
            <transition name="nav"><router-view /></transition></v-col
        ></v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import NavBar from "./components/NavBar.vue";

export default {
  name: "App",

  components: {
    NavBar,
  },

  data: () => ({
    //
  }),
};
</script>
<style>
.app {
  background: rgb(252, 250, 69);
  background: linear-gradient(
    90deg,
    rgba(252, 250, 69, 0.5858718487394958) 0%,
    rgba(0, 0, 0, 1) 20%,
    rgba(0, 0, 0, 1) 80%,
    rgba(252, 250, 69, 0.4542191876750701) 100%
  );
}
.nav-enter-from,
.nav-leave-to {
  opacity: 0;
}
.nav-enter-to,
.nav-leave-from {
  opacity: 1;
}
.nav-enter-active,
.nav-leave-active {
  transition: opacity 0.6s ease;
}
</style>

<template>
  <v-app class="app">
    <v-main>
      <NavBar />
      <v-container>
        <v-row justify="center" class="search-bar"
          ><v-col cols="10">
            <all-events></all-events>
            <by-validators></by-validators></v-col
        ></v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import AllEvents from "./components/AllEvents/AllEvents.vue";
import ByValidators from "./components/ByValidators/ByValidators.vue";
import NavBar from "./components/NavBar.vue";

export default {
  name: "App",

  components: {
    NavBar,
    AllEvents,
    ByValidators,
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
</style>

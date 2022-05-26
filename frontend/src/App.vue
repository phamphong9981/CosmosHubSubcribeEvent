<template>
  <v-app style="background: rgb(0,0,0);
background: radial-gradient(circle, rgba(0,0,0,1) 50%, rgba(252,250,69,1) 82%, rgba(253,59,29,1) 100%);">
    <v-main>
      <NavBar />
      <v-container
        ><history-table></history-table>
        <search-bar></search-bar>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import HistoryTable from "./components/AllEvents/HistoryTable.vue";
import SearchBar from "./components/ByValidators/SearchBar.vue";
import NavBar from "./components/NavBar.vue";

export default {
  name: "App",

  components: {
    NavBar,
    HistoryTable,
    SearchBar,
  },

  data: () => ({
    //
  }),
};
</script>
<style></style>

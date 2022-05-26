<template>
  <v-app style="background: rgb(252,250,69);
background: linear-gradient(90deg, rgba(252,250,69,0.5858718487394958) 0%, rgba(0,0,0,1) 20%, rgba(0,0,0,1) 80%, rgba(252,250,69,0.4542191876750701) 100%);">
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

<template>
  <div>
    <search-bar style="margin-bottom: 30px"></search-bar>
    <history-table></history-table>
  </div>
</template>

<script>
import { useStore } from 'vuex';
import HistoryTable from "./HistoryTable.vue";
import SearchBar from "./SearchBar.vue";
export default {
  components: { HistoryTable, SearchBar },
  setup() {
    const store=useStore()
    store.dispatch("table/getDataDelegator",[""])
  },
};
</script>

<style></style>

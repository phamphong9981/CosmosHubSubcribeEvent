<template>
  <div>
    <search-bar style="margin-bottom:30px"></search-bar>
    <history-table></history-table>
  </div>
</template>

<script>
import HistoryTable from "./HistoryTable.vue";
import SearchBar from './SearchBar.vue';
export default {
  components: { HistoryTable, SearchBar },
};
</script>

<style>
</style>
<template>
  <div style="text-align:center">
    <search-bar style="margin-bottom: 30px"></search-bar>
    <history-table></history-table>
    <button style="margin-top:13px; font-size:20px; font-family: 'Uchen', serif;" @click="viewMore()">View more</button>
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
    let count=0
    store.dispatch("table/getDataDelegator",[""])
    store.commit("table/updateDelegator","")
    function viewMore(){
      count++
      store.dispatch("table/viewMoreByDelegator",[count])
    }
    return {
      viewMore
    }
  },
};
</script>

<style></style>

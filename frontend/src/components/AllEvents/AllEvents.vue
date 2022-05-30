<template>
  <div>
    <h1 style="text-align: center; margin-bottom: 20px">All events</h1>
    <history-table></history-table>
  </div>
</template>

<script>
import { useStore } from "vuex";
import HistoryTable from "./HistoryTable.vue";
// import { getData } from "@/api/get.js";
import { ref } from "@vue/reactivity";
import { onMounted } from "vue";
export default {
  components: { HistoryTable },
  setup() {
    const store = useStore();
    const data = ref(null);
    onMounted(async () => {
      const response = await fetch("http://localhost:8088/unbond/all", {
        mode: "no-cors",
      });
      console.log(response);
      store.commit("table/updateData", response);
    });
    return {
      store,
      data,
    };
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=PT+Sans&display=swap");
h1 {
  font-family: "PT Sans", sans-serif;
  font-size: 43px;
}
</style>

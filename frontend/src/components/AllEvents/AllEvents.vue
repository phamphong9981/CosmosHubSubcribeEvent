<template>
  <div style="text-align: center">
    <h1 style="text-align: center; margin-bottom: 20px">All events</h1>
    <history-table></history-table>
    <button
      style="margin-top: 13px; font-size: 20px; font-family: 'Uchen', serif"
      @click="viewMore()"
    >
      View more
    </button>
  </div>
</template>

<script>
import { onBeforeUnmount } from "@vue/runtime-core";
import { useStore } from "vuex";
import HistoryTable from "./HistoryTable.vue";

export default {
  components: { HistoryTable },
  setup() {
    const store = useStore();
    const socket = new WebSocket("ws://localhost:8088/websocket");
    store.dispatch("table/getData");
    socket.onmessage = function (message) {
      console.log(message);
      store.commit("table/newData", message.data);
    };
    socket.onopen = function () {
      console.log("Successfully connected to the echo websocket server...");
    };
    onBeforeUnmount(() => {
      socket.close();
    });
    async function viewMore() {
      await store.dispatch("table/viewMoreAll");
    }
    return {
      store,
      viewMore,
    };
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=PT+Sans&display=swap");
@import url("https://fonts.googleapis.com/css2?family=Uchen&display=swap");
h1 {
  font-family: "PT Sans", sans-serif;
  font-size: 43px;
}
</style>

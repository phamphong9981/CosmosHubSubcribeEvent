<template>
  <v-row justify="center" class="search-bar"
    ><v-col cols="8">
      <form action="" method="get">
        <input
          type="text"
          placeholder="Search by validator address..."
          v-model="validator"
        />
        <button type="submit" @click.prevent="search()">
          <i class="fa-solid fa-magnifying-glass"></i>
        </button>
      </form> </v-col
  ></v-row>
</template>

<script>
import { ref } from "@vue/reactivity";
import { useStore } from "vuex";
import { onBeforeUnmount } from "@vue/runtime-core";
export default {
  setup() {
    const store = useStore();
    const validator = ref("");
    const socket = ref(null);
    const search = function () {
      if (socket.value) {
        socket.value.close();
      }
      store.dispatch("table/getDataValidator", [validator.value]);
      if (validator.value != "") {
        socket.value = new WebSocket(
          "ws://localhost:8088/websocket/validator/" + validator.value
        );
        socket.value.onmessage = function (message) {
          store.commit("table/newData", message.data);
        };
        socket.value.onopen = function () {
          console.log(
            "Successfully connected to the " +
              validator +
              " websocket server..."
          );
        };
      }
    };
    onBeforeUnmount(() => {
      if (socket.value) {
        socket.value.close();
      }
    });
    return {
      validator,
      search,
      store,
    };
  },
};
</script>

<style>
form {
  background: white;
  border-radius: 40px;
}
input[type="text"] {
  height: 57px;
  border: none;
  outline: none;
  color: black;
  width: 95%;
  padding: 15px;
}
button i {
  color: gray;
}
::placeholder {
  color: gray;
}
.search-bar {
  margin-top: 20px;
}
</style>

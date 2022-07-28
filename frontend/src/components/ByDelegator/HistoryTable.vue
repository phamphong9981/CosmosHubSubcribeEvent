<template>
  <div>
    <h1 style="text-align: center; margin-bottom: 20px">{{ validator }}</h1>
    <div class="background">
      <table>
        <tr>
          <th>Validator address</th>
          <th>Amount</th>
          <th>Time</th>
          <th>Transaction hash</th>
        </tr>
        <transition-group name="table-row">
          <tr v-for="item in data" :key="JSON.stringify(item)">
            <td>{{ item.validator }}</td>
            <td>{{ item.amount }}</td>
            <td>{{ item.time }}</td>
            <td>{{ item.tx_hash }}</td>
          </tr>
        </transition-group>
      </table>
    </div>
  </div>
</template>

<script>
import { useStore } from "vuex";
export default {
  setup() {
    // pull data ve
    const store = useStore();
    const data = store.state.table.items;
    return {
      data,
      store,
    };
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Anton&family=Josefin+Slab:wght@400&display=swap");
* {
  color: white;
}
table {
  width: 100%;
  border-radius: 20px;
}
table,
th,
td {
  border: 1px solid white;
  border-collapse: collapse;
  text-align: center;
  font-family: "Josefin Slab", serif;
}

th {
  height: 80px;
  font-family: "Anton", sans-serif;
}
td {
  height: 40px;
}
.background {
  background: gray;
}
</style>

<template>
  <div class="background">
    <table>
      <tr>
        <th>Delegator address</th>
        <th>Validator address</th>
        <th>Amount</th>
        <th>Time</th>
      </tr>
      <transition-group name="table-row">
        <tr v-for="item in data" :key="JSON.stringify(item)">
          <td>{{ item.delegator }}</td>
          <td>{{ item.validator }}</td>
          <td>{{ item.amount }}</td>
          <td>{{ item.time }}</td>
        </tr>
      </transition-group>
    </table>
  </div>
</template>

<script>
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();
    const data = store.state.table.items;
    return {
      data,
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
  /* max-height: 500px; */
  overflow: hidden;
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
  font-size: 20px;
}

.background {
  background: gray;
  overflow: hidden;
}
tr:nth-child(0) {
  height: 80px;
}
.table-row-move {
  transition: transform 1s;
}
</style>

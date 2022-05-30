<template>
  <div>
    <h1 style="text-align: center; margin-bottom:20px">{{validator}}</h1>
    <div class="background">
      <table>
        <tr>
          <th>Amount</th>
          <th>Time</th>
          <th>Completion time</th>
        </tr>
        <tr v-for="item in desserts" :key="item.name">
          <td>{{ item.name }}</td>
          <td>{{ item.calories }}</td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    validator: {
      type: String,
      required: false,
      default: "12313d4sfdsgsfgfsg64fs65gfds5g46",
    },
  },
  setup(props) {
    // pull data ve
    console.log(props);
    return {
      desserts: [
        {
          name: "Frozen Yogurt",
          calories: 159,
          fat: 6.0,
          carbs: 24,
          protein: 4.0,
          iron: "1%",
        },
        {
          name: "Ice cream sandwich",
          calories: 237,
          fat: 9.0,
          carbs: 37,
          protein: 4.3,
          iron: "1%",
        },
        {
          name: "Eclair",
          calories: 262,
          fat: 16.0,
          carbs: 23,
          protein: 6.0,
          iron: "7%",
        },
        {
          name: "Cupcake",
          calories: 305,
          fat: 3.7,
          carbs: 67,
          protein: 4.3,
          iron: "8%",
        },
        {
          name: "Gingerbread",
          calories: 356,
          fat: 16.0,
          carbs: 49,
          protein: 3.9,
          iron: "16%",
        },
        {
          name: "Jelly bean",
          calories: 375,
          fat: 0.0,
          carbs: 94,
          protein: 0.0,
          iron: "0%",
        },
        {
          name: "Lollipop",
          calories: 392,
          fat: 0.2,
          carbs: 98,
          protein: 0,
          iron: "2%",
        },
        {
          name: "Honeycomb",
          calories: 408,
          fat: 3.2,
          carbs: 87,
          protein: 6.5,
          iron: "45%",
        },
        {
          name: "Donut",
          calories: 452,
          fat: 25.0,
          carbs: 51,
          protein: 4.9,
          iron: "22%",
        },
        {
          name: "KitKat",
          calories: 518,
          fat: 26.0,
          carbs: 65,
          protein: 7,
          iron: "6%",
        },
      ],
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

